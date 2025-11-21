import { FunctionComponent, KeyboardEvent, useRef, useState } from "react";
import { Footer } from "./Footer/Footer";
import { Header } from "./Header/Header";
import styles from "./Layout.module.css";
import { LayoutProps } from "./Layout.props";
import cn from "classnames";

const Layout = ({ children }: LayoutProps): JSX.Element => {
	const [isSkipLinkDisplayed, setIsSkipLinkDisplayed] = useState<boolean>(false);

	const bodyRef = useRef<HTMLDivElement>(null);

	const skipContentAction = (key: KeyboardEvent) => {
		if (key.code === 'Space' || key.code === 'Enter') {
			key.preventDefault();
			bodyRef.current?.focus();
		}
		setIsSkipLinkDisplayed(false);
	};

	return (
		<div className={styles.wrapper}>
			<a
				onFocus={() => setIsSkipLinkDisplayed(true)}
				onKeyDown={skipContentAction}
				tabIndex={0}
				className={cn(styles['skip-link'], {
					[styles['displayed']]: isSkipLinkDisplayed,
				})}
			>Сразу к содержанию</a>
			<Header className={styles.header} />
			<main ref={bodyRef} tabIndex={0} className={styles.body} role='main'>
				{children}
			</main>
			<Footer className={styles.footer} />
		</div>
	);
};

export const withLayout = <T extends Record<string, unknown>>(Component: FunctionComponent<T>) => {
	return function withLayoutComponent(props: T): JSX.Element {
		return (
				<Layout>
					<Component {...props} />
				</Layout>
		);
	};
};