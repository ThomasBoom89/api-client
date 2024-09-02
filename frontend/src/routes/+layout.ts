import { Get } from '$lib/wailsjs/go/frontend/Configuration';
import { GetAll } from '../lib/wailsjs/go/frontend/Projects';
import { GetAll as GetAllCollections } from '../lib/wailsjs/go/frontend/Collections';
import { GetAll as GetAllRequests } from '../lib/wailsjs/go/frontend/HttpRequests';

export const prerender = false;
export const ssr = false;

export async function load() {
	return {
		configuration: await Get(),
		projects: await GetAll(),
		collections: await GetAllCollections(),
		requests: await GetAllRequests(),
		// delay: await delay(1000),
	};
}

const delay = (ms: number) => new Promise((res) => setTimeout(res, ms));
