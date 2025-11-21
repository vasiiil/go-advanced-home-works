import cn from 'classnames';
import { Button } from "../Button/Button";
import { Input } from "../Input/Input";
import { Rating } from "../Rating/Rating";
import { Textarea } from "../Textarea/Textarea";
import styles from "./ReviewForm.module.css";
import { ReviewFormProps } from "./ReviewForm.props";
import CloseIcon from "./close.svg";
import { useForm, Controller } from 'react-hook-form';
import { IReviewForm, IReviewSentResponse } from './ReviewForm.interface';
import axios from 'axios';
import { API } from '@/helpers/api';
import { useState } from 'react';

export const ReviewForm = ({ productId, isOpened, className, ...props }: ReviewFormProps): JSX.Element => {
	const { register, control, handleSubmit, formState: { errors }, reset, clearErrors } = useForm<IReviewForm>();
	const [isSuccess, setIsSuccess] = useState<boolean>(false);
	const [error, setError] = useState<string>();
	const onSubmit = async (formData: IReviewForm) => {
		try {
			const { data } = await axios.post<IReviewSentResponse>(API.review.createDemo, { ...formData });
			if (data.message) {
				setIsSuccess(true);
				reset();
			}
			else {
				setError('Что-то пошло не так');
			}
		} catch (e: unknown) {
			if (e instanceof Error) {
				setError(e.message);
			} else {
				setError('Неизвестная ошибка: ' + e);
			}
		}
	};

	return (
		<form onSubmit={handleSubmit(onSubmit )}>
			<div className={cn(styles['review-form'], className)}
				{...props}
			>
				<Input
					{...register('name', { required: { value: true, message: 'Заполните имя!' } })}
					error={errors.name}
					placeholder="Имя"
					tabIndex={isOpened ? 0 : -1}
					aria-invalid={!!errors.name}
				/>
				<Input
					{...register('title', { required: { value: true, message: 'Заполните заголовок!' } })}
					error={errors.title}
					placeholder="Заголовок отзыва"
					className={styles['title']}
					tabIndex={isOpened ? 0 : -1}
					aria-invalid={!!errors.title}
				/>
				<div className={styles['rating']}>
					<span>Оценка:</span>
					<Controller
						control={control}
						name="rating"
						rules={{required: {value: true, message: 'Укажите рейтинг!' } }}
						render={({ field }) => (
							<Rating
								rating={field.value}
								ref={field.ref}
								isEditable
								setRating={field.onChange}
								error={errors.rating}
								tabIndex={isOpened ? 0 : -1}
							/>
						)}
					/>
				</div>
				<Textarea
					{...register('description', { required: { value: true, message: 'Заполните текст!' } })}
					error={errors.description}
					placeholder="Текст отзыва"
					className={styles['description']}
					tabIndex={isOpened ? 0 : -1}
					aria-label='Текст отзыва'
					aria-invalid={!!errors.description}
				/>
				<div className={styles['submit']}>
					<Button appearance="primary" tabIndex={isOpened ? 0 : -1} onClick={() => clearErrors()}>Отправить</Button>
					<span className={styles['info']}>* Перед публикацией отзыв пройдет предварительную модерацию и проверку</span>
				</div>
			</div>
			{isSuccess && <div role='alert' className={cn(styles['panel'], styles['success'])}>
				<div className={styles['success-title']}>Ваш отзыв отправлен</div>
				<div>Спасибо, ваш отзыв будет опубликован после порверки.</div>
				<button
					className={styles['close']}
					onClick={() => setIsSuccess(false)}
					aria-label='Закрыть оповещение'
				>
					<CloseIcon />
				</button>
			</div>}
			{error && <div className={cn(styles['panel'], styles['error'])}>
				Что-то пошло не так, попробуйте обновить страницу
				<button
					className={styles['close']}
					onClick={() => setError(undefined)}
					aria-label='Закрыть оповещение'
				>
					<CloseIcon />
				</button>
			</div>}
		</form>
	);
};