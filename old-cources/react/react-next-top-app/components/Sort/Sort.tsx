import cn from 'classnames';
import styles from "./Sort.module.css";
import { SortEnum, SortProps } from "./Sort.props";
import SortIcon from "./sort.svg";

export const Sort = ({ sort, setSort, className, ...props }: SortProps): JSX.Element => {
	return (
		<div className={cn(styles['sort'], className)} {...props}>
			<div className={styles['sort-name']} id="sort">Сортировка</div>
			<button
				id="rating"
				onClick={(): void => setSort(SortEnum.Rating)}
				className={cn({
					[styles['active']]: sort === SortEnum.Rating
				})}
				aria-selected={sort === SortEnum.Rating}
				aria-labelledby="sort rating"
			>
				<SortIcon className={styles['sort-icon']} />По рейтингу
			</button>
			<button
				id="price"
				onClick={(): void => setSort(SortEnum.Price)}
				className={cn({
					[styles['active']]: sort === SortEnum.Price
				})}
				aria-selected={sort === SortEnum.Price}
				aria-labelledby="sort price"
			>
				<SortIcon className={styles['sort-icon']} />По цене
			</button>
		</div>
	);
};