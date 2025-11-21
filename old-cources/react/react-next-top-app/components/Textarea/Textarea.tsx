import cn from 'classnames';
import styles from "./Textarea.module.css";
import { TextareaProps } from "./Textarea.props";
import { ForwardedRef, forwardRef } from 'react';

export const Textarea = forwardRef((
	{ error, className, ...props }: TextareaProps,
	ref: ForwardedRef<HTMLTextAreaElement>
): JSX.Element => {
	return (
		<div className={cn(className, styles['wrapper'])}>
			<textarea className={cn(styles['textarea'], {
				[styles['error']]: error
			})} ref={ref} {...props} />
			{error && <span role='alert' className={styles['error-message']}>{error.message}</span>}
		</div>
	);
});
Textarea.displayName = 'Textarea';