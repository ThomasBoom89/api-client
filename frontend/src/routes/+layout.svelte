<script>
	import './styles.css';
	import Header from './Header.svelte';
	import { initializeConfigurationStore } from './createConfigurationState.svelte.ts';

	let { children, data } = $props();
	document.getElementById('loader')?.remove();
	initializeConfigurationStore(data.configuration);
</script>
<div class="w-[100vw] h-[100vh] bg-background text-text p-2">
	<Header />
	<hr class="border-background-accent" />
	<main class="p-2">
		{@render children()}
	</main>
	<footer>
	</footer>
</div>
<style>
</style>
