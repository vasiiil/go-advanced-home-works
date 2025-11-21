import cn from 'classnames';
import styles from "./Ptag.module.css";
import { PtagProps } from "./Ptag.props";

export const Ptag = ({ size = 'm', className, children, ...props }: PtagProps): JSX.Element => {
	return (
		<p
			className={cn(styles.p, className, {
				[styles.m]: size === 'm',
				[styles.l]: size === 'l',
				[styles.s]: size === 's',
			})}
			{ ...props }
		>
			{children}
		</p>
	);
};