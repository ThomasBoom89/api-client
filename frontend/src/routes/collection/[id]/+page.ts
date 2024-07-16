export async function load({ params }) {
	const id: number = Number(params.id);
	const requests: InternalRequest[] = [
		{
			id: 1,
			type: 'http',
			name: 'nice endpoint',
		},
		{
			id: 2,
			type: 'websocket',
			name: 'nice realtime',
		},
	];

	return {
		promise: new Promise<{ requests: InternalRequest[]; delay: any }>(async (resolve) => {
			const del = await delay(1000);
			resolve({ requests: requests, delay: del });
		}),
	};
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));
