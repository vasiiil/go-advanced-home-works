import cn from "classnames";
import styles from "./Header.module.css";
import { motion, useReducedMotion } from 'framer-motion';
import { HeaderProps } from "./Header.props";
import Logo from "../logo.svg";
import { ButtonIcon } from '@/components/ButtonIcon/ButtonIcon';
import { useEffect, useState } from 'react';
import { useRouter } from 'next/router';

export const Header = ({ className, ...props }: HeaderProps): JSX.Element => {
	const [isOpened, setIsOpened] = useState<boolean>(false);
	const shouldReduceMotion = useReducedMotion();
	const variants = {
		opened: {
			opacity: 1,
			x: 0,
			transition: {
				stiffness: 20,
			},
		},
		closed: {
			opacity: shouldReduceMotion ? 1 : 0,
			x: '100%',
		},
	};

	const router = useRouter();
	useEffect(() => {
		setIsOpened(false);
	}, [router]);

	return (
		<header className={cn(className, styles['header'])} {...props}>
			<Logo />
			<ButtonIcon appearance='white' icon='menu' onClick={() => setIsOpened(true)} />
			<motion.div
				className={styles['mobile-menu']}
				variants={variants}
				initial="closed"
				animate={isOpened ? 'opened' : 'closed'}
			>
				<ButtonIcon className={styles['menu-close']} appearance='white' icon='close' onClick={() => setIsOpened(false)} />
			</motion.div>
		</header>
	);
};