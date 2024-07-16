export async function load() {
	const projects: Project[] = [
		{
			id: 1,
			name: 'first project',
		},
		{
			id: 2,
			name: 'second project',
		},
	];

	return {
		promise: new Promise<{ projects: Project[]; delay: any }>(async (resolve) => {
			const del = await delay(1000);
			resolve({ projects: projects, delay: del });
		}),
	};
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));
