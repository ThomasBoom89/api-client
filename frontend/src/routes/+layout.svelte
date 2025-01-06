<script>
	import './styles.css';
	import Header from './Header.svelte';
	import { getConfigurationStore, initializeConfigurationStore } from '../lib/configurationStore.svelte.ts';
	import { initializeThemeStore } from '../lib/themeStore.svelte.ts';
	import { initializeProjectStore } from '../lib/projectStore.svelte.ts';
	import { initializeCollectionStore } from '../lib/collectionStore.svelte.ts';
	import { initializeRequestStore } from '../lib/requestStore.svelte.ts';
	import { initializeNavigationSystem } from '$lib/navigationSystem.svelte.ts';

	let { children, data } = $props();
	initializeConfigurationStore(data.configuration);
	initializeProjectStore(data.projects);
	initializeCollectionStore(data.collections);
	initializeRequestStore(data.requests);
	initializeThemeStore(getConfigurationStore());
	initializeNavigationSystem();
	document.getElementById('loader')?.remove();
</script>

<div class="w-[100vw] h-[100vh] bg-background text-text flex flex-col overflow-hidden">
	<Header />
	<main class="flex flex-col w-full h-full overflow-hidden flex-1">
		{@render children()}
	</main>
</div>

<style>
</style>
