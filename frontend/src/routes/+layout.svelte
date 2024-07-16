<script>
	import './styles.css';
	import Header from './Header.svelte';
	import { initializeConfigurationStore } from './createConfigurationState.svelte.ts';
	import Footer from './Footer.svelte';

	let { children, data } = $props();
	initializeConfigurationStore(data.configuration);
	document.getElementById('loader')?.remove();
</script>

<div class="w-[100vw] h-[100vh] bg-background text-text p-2 flex flex-col">
	<Header />
	<hr class="border-background-accent" />
	<main class="w-full h-full">
		{@render children()}
	</main>
	<Footer />
</div>

<style>
</style>
