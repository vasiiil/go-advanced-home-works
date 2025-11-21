import cn from "classnames";
import { useEffect, useState, KeyboardEvent, forwardRef, ForwardedRef, useRef } from "react";
import styles from "./Rating.module.css";
import { RatingProps } from "./Rating.props";
import StarIcon from "./star.svg";

export const Rating = forwardRef((
	{ isEditable = false, rating, setRating, error, tabIndex, ...props }: RatingProps,
	ref: ForwardedRef<HTMLDivElement>
): JSX.Element => {
	const [ratingArray, setRatingArray] = useState<JSX.Element[]>(new Array(5).fill(<></>));
	const ratingArrayRef = useRef<(HTMLSpanElement | null)[]>([]);

	useEffect(() => {
		constructRating(rating);
	// eslint-disable-next-line react-hooks/exhaustive-deps
	}, [rating, tabIndex]);

	const computeFocus = (r: number, index: number): number => {
		if (!isEditable) {
			return -1;
		}
		if (!rating && index === 0) {
			return tabIndex ?? 0;
		}
		if (r && r === index + 1) {
			return tabIndex ?? 0;
		}

		return -1;
	};

	const constructRating = (currentRating: number): void => {
		const updatedArray = ratingArray.map((r: JSX.Element, i: number) => {
			return (
				<span
					key={i}
					className={cn(styles.star, {
						[styles.filled]: i < currentRating,
						[styles.editable]: isEditable
					})}
					onMouseEnter={(): void => changeDisplay(i + 1)}
					onMouseLeave={(): void => changeDisplay(rating)}
					onClick={(): void => setRatingDispatcher(i + 1)}
					tabIndex={computeFocus(rating, i)}
					onKeyDown={handleKey}
					ref={r => ratingArrayRef.current?.push(r)}
					role={isEditable ? 'slider' : ''}
					aria-valuenow={rating}
					aria-valuemin={1}
					aria-valuemax={5}
					aria-label={isEditable ? 'Укажите рейтинг' : ('рейтинг' + rating)}
					aria-invalid={!!error}
				>
					<StarIcon />
				</span>
			);
		});

		setRatingArray(updatedArray);
	};

	const changeDisplay = (i: number): void => {
		if (!isEditable) {
			return;
		}

		constructRating(i);
	};

	const setRatingDispatcher = (rating: number): void => {
		if (!isEditable || !setRating) {
			return;
		}

		setRating(rating);
	};

	const handleKey = (e: KeyboardEvent): void => {
		if (!isEditable || !setRating) {
			return;
		}
		if (e.code === 'ArrowRight' || e.code === 'ArrowUp') {
			if (!rating) {
				setRating(1);
			}
			else {
				e.preventDefault();
				setRating(rating < 5 ? rating + 1 : 5);
			}
			ratingArrayRef.current[rating]?.focus();
		}
		if (e.code === 'ArrowLeft' || e.code === 'ArrowDown') {
			e.preventDefault();
			setRating(rating > 1 ? rating - 1 : 1);
			ratingArrayRef.current[rating - 2]?.focus();
		}

	};

	return (
		<div { ...props } ref={ref} className={cn(styles['wrapper'], {
			[styles['error']]: error
		})}>
			{ratingArray.map((r, i) => (<span key={i}>{r}</span>))}
			{error && <span role='alert' className={styles['error-message']}>{error.message}</span>}
		</div>
	);
});
Rating.displayName = 'Rating';