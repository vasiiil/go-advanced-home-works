import cn from 'classnames';
import styles from "./Card.module.css";
import { CardProps } from "./Card.props";
import { ForwardedRef, forwardRef } from 'react';

export const Card = forwardRef(({ color = 'white', className, children, ...props }: CardProps, ref: ForwardedRef<HTMLDivElement>): JSX.Element => {
	return (
		<div className={cn(styles.card, className, {
			[styles.blue]: color === 'blue',
			[styles.white]: color === 'white'
		})}
		ref={ref}
		{...props}
		>
			{children}
		</div>
	);
});
Card.displayName = 'Card';