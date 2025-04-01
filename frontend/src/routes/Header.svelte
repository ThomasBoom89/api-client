<script lang="ts">
	import { getThemeStore } from '../lib/themeStore.svelte.ts';
	import { NavigationState } from '$lib/enums/NavigationState.ts';
	import { getNavigationSystem } from '$lib/navigationSystem.svelte.ts';
	import { getCollectionStore } from '$lib/collectionStore.svelte.ts';
	import { getProjectStore } from '$lib/projectStore.svelte.ts';

	const themeStore = getThemeStore();
	const navigationSystem = getNavigationSystem();
	const collectionStore = getCollectionStore();
	const projectStore = getProjectStore();
</script>

<header class="px-2 pb-2">
	<div class="flex flex-row items-center justify-between">
		<nav aria-label="Breadcrumb">
			<ol class="flex list-none items-stretch gap-2">
				<li class="flex items-center gap-2">
					<button
						onclick={() => navigationSystem.navigateToHome()}
						class="text-text hover:text-text-accent flex max-w-[20ch] items-center gap-1 truncate whitespace-nowrap
						transition-colors"
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-5 w-5"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="1.5"
							aria-hidden="true"
							aria-labelledby="title-01 description-01"
							role="link"
						>
							<title id="title-01">Home</title>
							<desc id="description-01"> Home button indicating the homepage of the website.</desc>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
							/>
						</svg>
						Home
					</button>
					{#if navigationSystem.currentState > NavigationState.Home}
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="text-text h-4 w-4 flex-none rotate-180"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="1.5"
							aria-hidden="true"
							aria-labelledby="title-02 description-02"
							role="graphics-symbol"
						>
							<title id="title-02">Arrow</title>
							<desc id="description-02">
								Arrow icon that points to the next page in big screen resolution sizes and previous page in small screen
								resolution sizes.
							</desc>
							<path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
						</svg>
					{/if}
				</li>
				{#if navigationSystem.currentState > NavigationState.Home}
					<li class="flex items-center gap-2">
						<button
							onclick={() => navigationSystem.navigateToOverview()}
							class="text-text hover:text-text-accent flex max-w-[20ch] truncate whitespace-nowrap transition-colors"
							>projects
						</button>
						{#if navigationSystem.currentState > NavigationState.Overview}
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="text-text h-4 w-4 flex-none rotate-180"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								stroke-width="1.5"
								aria-hidden="true"
								aria-labelledby="title-03 description-03"
								role="graphics-symbol"
							>
								<title id="title-03">Arrow</title>
								<desc id="description-03">
									Arrow icon that points to the next page in big screen resolution sizes and previous page in small
									screen resolution sizes.
								</desc>
								<path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
							</svg>
						{/if}
					</li>
				{/if}
				{#if navigationSystem.currentState > NavigationState.Overview}
					<li class="flex items-center gap-2">
						<button
							onclick={() => navigationSystem.navigateToProject(navigationSystem.currentProjectId)}
							class="text-text hover:text-text-accent flex max-w-[20ch] truncate whitespace-nowrap transition-colors"
							>{projectStore.getById(navigationSystem.currentProjectId).name}
						</button>
						{#if navigationSystem.currentState > NavigationState.Project}
							<svg
								xmlns="http://www.w3.org/2000/svg"
								class="text-text h-4 w-4 flex-none rotate-180"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
								stroke-width="1.5"
								aria-hidden="true"
								aria-labelledby="title-04 description-04"
								role="graphics-symbol"
							>
								<title id="title-04">Arrow</title>
								<desc id="description-04">
									Arrow icon that points to the next page in big screen resolution sizes and previous page in small
									screen resolution sizes.
								</desc>
								<path stroke-linecap="round" stroke-linejoin="round" d="M15 19l-7-7 7-7" />
							</svg>
						{/if}
					</li>
				{/if}
				{#if navigationSystem.currentState > NavigationState.Project}
					<li class="flex flex-1 items-center">
						<button
							onclick={() => navigationSystem.navigateToCollection(navigationSystem.currentCollectionId)}
							aria-current="page"
							class="text-text-disabled pointer-events-none max-w-[20ch] truncate whitespace-nowrap"
							>{collectionStore.getById(navigationSystem.currentCollectionId).name}
						</button>
					</li>
				{/if}
			</ol>
		</nav>
		<ul class="ml-auto flex flex-row items-center gap-2">
			<li class="flex items-center">
				<button
					onclick={() => {
						themeStore.switchTheme();
					}}
					data-testid="dark-light-toggle"
				>
					{#if themeStore.currentTheme === 'dark'}
						<svg
							class="h-6 w-6"
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
							class="h-6 w-6"
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
	</div>
	<hr class="border-background-accent" />
</header>
