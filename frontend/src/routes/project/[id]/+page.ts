export async function load({ params }) {
	const id: number = Number(params.id);
	const collections: Collection[] = [
		{
			id: 1,
			name: 'collection one',
		},
		{
			id: 2,
			name: 'collection two',
		},
	];

	return {
		promise: new Promise<{ collections: Collection[]; delay: any }>(async (resolve) => {
			const del = await delay(1000);
			resolve({ collections: collections, delay: del });
		}),
	};
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));
