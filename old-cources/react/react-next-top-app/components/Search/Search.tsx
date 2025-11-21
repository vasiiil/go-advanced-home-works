import cn from 'classnames';
import { KeyboardEventHandler, useState } from "react";
import { Button } from "../Button/Button";
import { Input } from "../Input/Input";
import styles from "./Search.module.css";
import { SearchProps } from "./Search.props";
import GlassIcon from "./glass.svg";
import { useRouter } from "next/router";

export const Search = ({ className, ...props }: SearchProps): JSX.Element => {
	const [search, setSearch] = useState<string>('');
	const router = useRouter();

	const goToSearch = ():void => {
		router.push({
			pathname: '/search',
			query: {
				q: search
			}
		});
	};

	const handleKeyDown: KeyboardEventHandler<HTMLInputElement> = (e): void => {
		if (e.key === 'Enter') {
			goToSearch();
		}
	};

	return (
		<form className={cn(className, styles['search'])} {...props} role='search'>
			<Input
				className={styles['input']}
				placeholder="Поиск..."
				value={search}
				onChange={(e):void => setSearch(e.target.value)}
				onKeyDown={handleKeyDown}
			/>
			<Button
				appearance="primary"
				className={styles['button']}
				onClick={goToSearch}
				aria-label="Искать по сайту"
			>
				<GlassIcon />
			</Button>
		</form>
	);
};