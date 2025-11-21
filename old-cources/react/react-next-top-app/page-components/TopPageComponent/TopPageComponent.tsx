import { Advantages, HhData, Htag, Product, Sort, Tag } from "../../components";
import { ITopPageComponentProps } from "./TopPageComponent.props";
import styles from "./TopPageComponent.module.css";
import { TopLevelCategory } from "../../interfaces/page.interface";
import { SortEnum } from "../../components/Sort/Sort.props";
import { useEffect, useReducer } from "react";
import { sortReducer } from "./sort.reducer";
import { useReducedMotion } from 'framer-motion';

export const TopPageComponent = ({ page, products, firstCategory }: ITopPageComponentProps): JSX.Element => {
	const initialReduserState = { products, sort: SortEnum.Rating };
	const [{ products: sortedProducts, sort }, dispatchSort] = useReducer(sortReducer, initialReduserState);
	const shouldReduceMotion = useReducedMotion();

	useEffect(() => {
		if (products) {
			dispatchSort({ type: "RESET_STATE", products });
		}
	}, [products]);

	const setSort = (sort: SortEnum): void => {
		dispatchSort({ type: sort });
	};

	return (
		<div className={styles.wrapper}>
			<div className={styles.title}>
				<Htag tag="h1">{page.title}</Htag>
				{products && <Tag color="gray" size="m" aria-label={products.length + 'элементов'}>{products.length}</Tag>}
				<Sort sort={sort} setSort={setSort}/>
			</div>
			<div role='list'>
				{sortedProducts && sortedProducts.map(p => (<Product role='listitem' key={p._id} product={p} layout={shouldReduceMotion ? false : true} />))}
			</div>
			<div className={styles['hh-title']}>
				<Htag tag="h2">Вакансии - {page.category}</Htag>
				<Tag color="red" size="m">hh.ru</Tag>
			</div>
			{firstCategory === TopLevelCategory.Courses && page.hh && <HhData {...page.hh} />}
			{page.advantages && page.advantages.length > 0 && <>
				<Htag tag="h2">Перимущества</Htag>
				<Advantages advantages={page.advantages} />
			</>}
			{page.seoText && <div className={styles['seo']} dangerouslySetInnerHTML={{ __html: page.seoText}} />}
			<Htag tag="h2">Получаемые навыки</Htag>
			{page.tags.map(t => <Tag key={t} color="primary">{t}</Tag>)}
		</div>
	);
};