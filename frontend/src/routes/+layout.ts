export const prerender = true;
export const ssr = false;

import { Read } from '$lib/wailsjs/go/configuration/ReadWriter';

export async function load({ params }) {
	return {
		configuration: await Read()
	};
}
