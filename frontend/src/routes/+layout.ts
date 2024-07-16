import { Get } from '$lib/wailsjs/go/frontend/Configuration';

export const prerender = false;
export const ssr = false;

export async function load() {
	return {
		configuration: await Get(),
		delay: await delay(1000),
	};
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));
