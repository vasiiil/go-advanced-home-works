import { KeyboardEvent, useContext, useState } from "react";
import { AppContext } from "../../context/app.context";
import { IFirstLevelMenuItem, IPageItem } from "../../interfaces/menu.interface";
import styles from "./Menu.module.css";
import cn from "classnames";
import Link from "next/link";
import { useRouter } from "next/router";
import { firstLevelMenu } from "../../helpers/helpers";
import { motion, useReducedMotion } from "framer-motion";

export const Menu = (): JSX.Element => {
	const { menu, setMenu, firstCategory } = useContext(AppContext);
	const [announce, setAnnounce] = useState<'closed' | 'opened' | undefined>();

	const router = useRouter();

	const shouldReduceMotion = useReducedMotion();
	const variants = {
		visible: {
			marginBottom: 20,
			transition: shouldReduceMotion ? {} : {
				when: 'afterChildren',
				staggerChildren: 0.1,
			}
		},
		hidden: {
			marginBottom: 0,
		},
	};

	const variantsChildren = {
		visible: {
			opacity: 1,
			height: 29,
		},
		hidden: {
			opacity: shouldReduceMotion ? 1 : 0,
			height: 0,
		},
	};

	const openSecondLevel = (secondCategory: string): void => {
		setMenu && setMenu(menu.map(m => {
			if (m._id.secondCategory === secondCategory) {
				setAnnounce(m.isOpened ? 'closed' : 'opened');
				m.isOpened = !m.isOpened;
			}

			return m;
		}));
	};

	const openSecondLevelKey = (key: KeyboardEvent, secondCategory: string) => {
		if (key.code === 'Space' || key.code === 'Enter') {
			key.preventDefault();
			openSecondLevel(secondCategory);
		}
	};

	const buildFirstLevel = (): JSX.Element => {
		return (
			<ul className={styles.firstLevelList}>
				{firstLevelMenu.map(flmItem => (
					<li key={flmItem.route} aria-expanded={flmItem.id === firstCategory}>
						<Link href={`/${flmItem.route}`}>
							<div className={cn(styles.firstLevel, {
								[styles.firstLevelActive]: flmItem.id === firstCategory
							})}>
								{flmItem.icon}
								<span>{flmItem.name}</span>
							</div>
						</Link>
						{flmItem.id === firstCategory && buildSecondLevel(flmItem)}
					</li>
				))}
			</ul>
		);
	};

	const buildSecondLevel = (flmItem: IFirstLevelMenuItem): JSX.Element => {
		return (
			<ul className={styles.secondBlock}>
				{menu.map(slmItem => {
					if (slmItem.pages.map(p => p.alias).includes(router.asPath.split('/')[2])) {
						slmItem.isOpened = true;
					}
					return (
						<li key={slmItem._id.secondCategory}>
							<button
								className={styles.secondLevel}
								onKeyDown={(key: KeyboardEvent) => openSecondLevelKey(key, slmItem._id.secondCategory)}
								onClick={(): void => openSecondLevel(slmItem._id.secondCategory)}
								aria-expanded={slmItem.isOpened}
							>{slmItem._id.secondCategory}</button>
							<motion.ul
								layout
								variants={variants}
								initial={slmItem.isOpened ? 'visible' : 'hidden'}
								animate={slmItem.isOpened ? 'visible' : 'hidden'}
								className={styles.secondLevelBlock}
							>
								{buildThirdLevel(slmItem.pages, flmItem.route, slmItem.isOpened ?? false)}
							</motion.ul>
						</li>
					);
				})}
			</ul>
		);
	};

	const buildThirdLevel = (pages: IPageItem[], route: string, isOpened: boolean): JSX.Element[] => {
		return (
			pages.map(page => (
				<motion.li key={page.category}
					variants={variantsChildren}
				>
					<Link
						tabIndex={isOpened ? 0 : -1}
						href={`/${route}/${page.alias}`}
						aria-current={`/${route}/${page.alias}` === router.asPath ? 'page' : false}
					>
						<span className={cn(styles.thirdLevel, {
							[styles.thirdLevelActive]: `/${route}/${page.alias}` === router.asPath
						})}>
							{page.category}
						</span>
					</Link>
				</motion.li>
			))
		);
	};

	return (
		<nav className={styles.menu} role='navigation'>
			{announce && <span role='log' className='visualy-hidden'>{announce === 'opened' ? 'развернуто' : 'свернуто'}</span>}
			{buildFirstLevel()}
		</nav>
	);
};