<script>
	import './styles.css';
	import Header from './Header.svelte';
	import createConfigurationState from './createConfigurationState.svelte.ts';
	import { setContext } from 'svelte';

	let { children, data } = $props();
	const store = createConfigurationState(data.configuration);
	setContext('configurationStore', store);
</script>
<Header />
<main>
	{@render children()}
</main>
<footer>
</footer>
<style>
</style>
