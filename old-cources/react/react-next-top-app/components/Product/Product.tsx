import cn from 'classnames';
import { priceRu, wordEndings } from "../../helpers/helpers";
import { Button } from "../Button/Button";
import { Card } from "../Card/Card";
import { Divider } from "../Divider/Divider";
import { Rating } from "../Rating/Rating";
import { Tag } from "../Tag/Tag";
import styles from "./Product.module.css";
import { ProductProps } from "./Product.props";
import Image from "next/image";
import { forwardRef, useRef, useState } from "react";
import { Review } from "../Review/Review";
import { ReviewForm } from "../ReviewForm/ReviewForm";
import { motion } from "framer-motion";
import { ForwardedRef } from 'react';

export const Product = motion(forwardRef(function Product({ product, className, ...props }: ProductProps, ref: ForwardedRef<HTMLDivElement>): JSX.Element {
	const getImagePath = (src: string): string => /^https?:\/\//i.test(src) ? src : process.env.NEXT_PUBLIC_DOMAIN + src;
	const [isReviewOpened, setIsReviewOpened] = useState<boolean>(false);
	const reviewRef = useRef<HTMLDivElement>(null);

	const scrollToReview = () => {
		setIsReviewOpened(true);
		reviewRef.current?.scrollIntoView({
			behavior: 'smooth',
			block: 'start',
		});
		reviewRef.current?.focus();
	};

	const variants = {
		visible: { opacity: 1, height: 'auto' },
		hidden: { opacity: 0, height: 0 },
	};

	return (
		<div className={className} {...props} ref={ref}>
			<Card className={styles['product']}>
				<div className={styles['logo']}>
					<Image
						src={getImagePath(product.image)}
						alt={product.title}
						width={70}
						height={70}
					/>
				</div>
				<div className={styles['title']}>{product.title}</div>
				<div className={styles['price']}>
					<span><span className="visualy-hidden">цена</span>{priceRu(product.price)}</span>
					{product.oldPrice && <Tag className={styles['old-price']} color="green">
						<span className="visualy-hidden">скидка</span>
						{priceRu(product.price - product.oldPrice)}
					</Tag>}
				</div>
				<div className={styles['credit']}>
					<span className="visualy-hidden">кредит</span>
					{priceRu(product.credit)}<span className={styles['per-month']}>/мес</span>
				</div>
				<div className={styles['rating']}>
					<span className="visualy-hidden">{'рейтинг' + (product.reviewAvg ?? product.initialRating)}</span>
					<Rating rating={product.reviewAvg ?? product.initialRating} />
				</div>
				<div className={styles['tags']}>{product.categories.map(c => <Tag className={styles['category']} color="ghost" key={c}>{c}</Tag>)}</div>
				<div className={styles['price-title']} aria-hidden={true}>цена</div>
				<div className={styles['credit-title']} aria-hidden={true}>кредит</div>
				<div className={styles['rating-title']}>
					<a href="#ref" onClick={scrollToReview}>{product.reviewCount} {wordEndings(product.reviewCount, ['отзыв', 'отзыва', 'отзывов'])}</a>
				</div>
				<Divider className={styles['hr']} />
				<div className={styles['description']}>{product.description}</div>
				<div className={styles['features']}>
					{product.characteristics.map((c): JSX.Element => (
						<div className={styles['characteristic']} key={c.name}>
							<span className={styles['characteristic-name']}>{c.name}</span>
							<span className={styles['characteristic-dots']}></span>
							<span className={styles['characteristic-value']}>{c.value}</span>
						</div>
					))}
				</div>
				<div className={styles['advantages-block']}>
					{product.advantages && <div className={styles['advantages']}>
						<div className={styles['advantages-item-title']}>Перимущества</div>
						<div>{product.advantages}</div>
					</div>}
					{product.disadvantages && <div className={styles['disadvantages']}>
						<div className={styles['advantages-item-title']}>Недостатки</div>
						<div>{product.disadvantages}</div>
					</div>}
				</div>
				<Divider className={cn(styles['hr'], styles['hr2'])} />
				<div className={styles['actions']}>
					<Button appearance="primary">Узнать подробнее</Button>
					<Button
						className={styles['review-button']}
						appearance="ghost"
						arrow={isReviewOpened ? "down" : "right"}
						onClick={(): void => setIsReviewOpened(!isReviewOpened)}
						aria-expanded={isReviewOpened}
					>Читать отзывы</Button>
				</div>
			</Card>
			<motion.div
				animate={isReviewOpened ? 'visible' : 'hidden'}
				variants={variants}
				initial="hidden"
			>
				<Card color="blue" className={styles['reviews']} ref={reviewRef} tabIndex={isReviewOpened ? 0 : -1}>
					{product.reviews.map(review => (
						<div key={review._id}>
							<Review review={review} />
							<Divider />
						</div>
					))}
					<ReviewForm productId={product._id} isOpened={isReviewOpened} />
				</Card>
			</motion.div>
		</div>
	);
}));