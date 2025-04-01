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

<div class="bg-background text-text flex h-[100vh] w-[100vw] flex-col overflow-hidden">
	<Header />
	<main class="flex h-full w-full flex-1 flex-col overflow-hidden px-2 pb-2">
		{@render children()}
	</main>
</div>

<style>
</style>
