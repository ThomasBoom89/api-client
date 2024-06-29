export const prerender = true;
export const ssr = false;

import { Read } from '$lib/wailsjs/go/configuration/ReadWriter';

export async function load() {
	return {
		configuration: await Read(),
		delay: await delay(3000)
	};
}

const delay = (ms: number) => new Promise(res => setTimeout(res, ms));
