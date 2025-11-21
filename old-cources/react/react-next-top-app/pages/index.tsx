import { useEffect, useState } from "react";
import { Htag, Button } from "../components";
import { withLayout } from "../layout/Layout";

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
		</>
	);
}

export default withLayout(Home);