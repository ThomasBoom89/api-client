<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';
	import { LogDebug } from '$lib/wailsjs/runtime/runtime';
	import { Test } from '$lib/wailsjs/go/main/App';
	import { configuration } from '$lib/wailsjs/go/models.ts';
	import { getConfigurationStore } from './createConfigurationState.svelte.ts';

	const html = document.documentElement;
	const configurationStore = getConfigurationStore();
	let config: configuration.Configuration = configurationStore.configuration;
	$effect(() => {
		if (config.theme === '') {
			if (window.matchMedia('(prefers-color-scheme: dark)').matches) {
				html.setAttribute('data-theme', 'dark');
				configurationStore.setProperty('theme', 'dark');
			} else {
				html.setAttribute('data-theme', 'light');
				configurationStore.setProperty('theme', 'light');
			}
		} else {
			html.setAttribute('data-theme', config.theme);
		}
	});

	function navigate(page: string = ''): null {
		console.warn(page);
		if (page == undefined) {
			return null;
		}
		goto(page, {});

		return null;
	}

	function switchTheme(): void {
		LogDebug('switchTheme');
		const currentTheme = html.getAttribute('data-theme');
		if (currentTheme == 'dark') {
			html.setAttribute('data-theme', 'light');
			configurationStore.setProperty('theme', 'light');
		} else {
			html.setAttribute('data-theme', 'dark');
			configurationStore.setProperty('theme', 'dark');
		}
	}

	function test(): void {
		Test();
	}
</script>

<header class="px-2 pb-2">
	<nav class="flex flex-row justify-between items-center">
		<ul class="flex flex-row gap-2">
			<li aria-current={$page.url.pathname === '/' ? 'page' : undefined}>
				<button
					onclick={() => {
						navigate('/');
					}}>Home
				</button
				>
			</li>
			<li aria-current={$page.url.pathname === '/overview' ? 'page' : undefined}>
				<a href="/overview">Project Overview</a>
			</li>
			<li aria-current={$page.url.pathname === '/request' ? 'page' : undefined}>
				<button
					onclick={() => {
						navigate('/request');
					}}>Request
				</button
				>
			</li>
		</ul>
		<ul class="flex flex-row gap-2 items-center">
			<li>
				<button onclick={test}>Test</button>
			</li>
			<li class="flex items-center">
				<button onclick={switchTheme} data-testid="dark-light-toggle">
					{#if config.theme === 'dark'}
						<svg
							class="w-6 h-6"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								fill-rule="evenodd"
								d="M13 3a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0V3ZM6.343 4.929A1 1 0 0 0 4.93 6.343l1.414 1.414a1 1 0 0 0 1.414-1.414L6.343 4.929Zm12.728 1.414a1 1 0 0 0-1.414-1.414l-1.414 1.414a1 1 0 0 0 1.414 1.414l1.414-1.414ZM12 7a5 5 0 1 0 0 10 5 5 0 0 0 0-10Zm-9 4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2H3Zm16 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2h-2ZM7.757 17.657a1 1 0 1 0-1.414-1.414l-1.414 1.414a1 1 0 1 0 1.414 1.414l1.414-1.414Zm9.9-1.414a1 1 0 0 0-1.414 1.414l1.414 1.414a1 1 0 0 0 1.414-1.414l-1.414-1.414ZM13 19a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0v-2Z"
								clip-rule="evenodd"
							/>
						</svg>
					{:else}
						<svg
							class="w-6 h-6"
							aria-hidden="true"
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							fill="currentColor"
							viewBox="0 0 24 24"
						>
							<path
								fill-rule="evenodd"
								d="M11.675 2.015a.998.998 0 0 0-.403.011C6.09 2.4 2 6.722 2 12c0 5.523 4.477 10 10 10 4.356 0 8.058-2.784 9.43-6.667a1 1 0 0 0-1.02-1.33c-.08.006-.105.005-.127.005h-.001l-.028-.002A5.227 5.227 0 0 0 20 14a8 8 0 0 1-8-8c0-.952.121-1.752.404-2.558a.996.996 0 0 0 .096-.428V3a1 1 0 0 0-.825-.985Z"
								clip-rule="evenodd"
							/>
						</svg>
					{/if}
				</button>
			</li>
		</ul>
	</nav>
</header>
