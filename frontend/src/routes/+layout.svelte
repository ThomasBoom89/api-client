<script>
	import './styles.css';
	import Header from './Header.svelte';
	import { getConfigurationStore, initializeConfigurationStore } from '../lib/configurationStore.svelte.ts';
	import { initializeThemeStore } from '../lib/themeStore.svelte.ts';
	import Footer from './Footer.svelte';
	import { initializeProjectStore } from '../lib/projectStore.svelte.ts';
	import { initializeCollectionStore } from '../lib/collectionStore.svelte.ts';
	import { initializeRequestStore } from '../lib/requestStore.svelte.ts';

	let { children, data } = $props();
	initializeConfigurationStore(data.configuration);
	initializeProjectStore(data.projects);
	initializeCollectionStore(data.collections);
	initializeRequestStore(data.requests);
	initializeThemeStore(getConfigurationStore());
	document.getElementById('loader')?.remove();
</script>

<div class="w-[100vw] h-[100vh] bg-background text-text p-2 grid grid-rows-[min-content_min-content_1fr_min-content]">
	<Header />
	<hr class="border-background-accent" />
	<main class="w-full overflow-y-scroll mt-2">
		{@render children()}
	</main>
	<Footer />
</div>

<style>
</style>
