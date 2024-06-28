import { Read } from '$lib/wailsjs/go/configuration/ReadWriter';
import { configurationStore } from './createConfigurationState.svelte.ts';

class InitialState {
	error = $state();
	isLoading = $state(false);
}

export default function useFetchInitial() {
	const initialState = new InitialState();

	async function fetchData() {
		initialState.isLoading = true;
		try {
			const [config] = await Promise.all([Read(), delay(3000)]);
			configurationStore.configuration = config;
			initialState.error = undefined;
		} catch (err) {
			initialState.error = err;
		}
		initialState.isLoading = false;
	}

	fetchData().then();

	return initialState;
}

const delay = (ms: number) => new Promise(res => setTimeout(res, ms));
