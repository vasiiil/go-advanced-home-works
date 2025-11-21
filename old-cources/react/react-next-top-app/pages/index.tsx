import { GetStaticProps } from "next";
import { useEffect, useState } from "react";
import { Htag, Button, Ptag, Tag, Input, Textarea } from "../components";
import { withLayout } from "../layout/Layout";
import axios from "axios";
import { IMenuItem } from "../interfaces/menu.interface";
import { API } from '@/helpers/api';

function Home(): JSX.Element {
	const [counter, setCounter] = useState(0);

	useEffect(() => {
		console.log('Counter: ' + counter);
		return function cleanup() {
			console.log('Unmount');
		};
	});

	return (
		<>
			<Htag tag='h1'>{ counter }</Htag>
			<Button appearance="primary" arrow="right" onClick={(): void => setCounter(x => x + 1)}>Button</Button>
			<Button appearance="ghost" arrow="down" onClick={(): void => setCounter(x => x - 1)}>Button ghost</Button>
			<Ptag>M paragraph</Ptag>
			<Ptag size="l">L paragraph</Ptag>
			<Ptag size="s">S paragraph</Ptag>
			<Tag size="s">S ghost</Tag>
			<Tag color="red">M Red</Tag>
			<Tag size="s" color="green">S green</Tag>
			<Tag size="m" color="primary">M primary</Tag>
			<Input placeholder="test" />
			<Textarea placeholder="textarea placeholder" />
		</>
	);
}

export default withLayout(Home);

export const getStaticProps: GetStaticProps<IHomeProps> = async () => {
	const firstCategory = 0;
	const { data: menu } = await axios.post<IMenuItem[]>(API.topPage.find, {
		firstCategory
	});
	return {
		props: {
			menu,
			firstCategory,
		},
	};
};

interface IHomeProps extends Record<string, unknown> {
	menu: IMenuItem[];
	firstCategory: number;
}