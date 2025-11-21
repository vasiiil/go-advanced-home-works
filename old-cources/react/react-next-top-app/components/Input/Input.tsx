import cn from 'classnames';
import styles from "./Input.module.css";
import { InputProps } from "./Input.props";
import { ForwardedRef, forwardRef } from 'react';

export const Input = forwardRef((
	{ className, error, ...props }: InputProps,
	ref: ForwardedRef<HTMLInputElement>
): JSX.Element => {
	return (
		<div className={cn(className, styles['input-wrapper'])}>
			<input className={cn(styles['input'], {
				[styles['error']]: error
			})} ref={ref} {...props} />
			{error && <span role='alert' className={styles['error-message']}>{error.message}</span>}
		</div>
	);
});
Input.displayName = 'Input';