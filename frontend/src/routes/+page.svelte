<script lang="ts">
	import { getNavigationSystem } from '$lib/navigationSystem.svelte.ts';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime';

	const navigationSystem = getNavigationSystem();
</script>

<svelte:head>
	<title>Api-Client :: Startseite</title>
	<meta name="description" content="Startseite" />
</svelte:head>

<button
	onclick={() => navigationSystem.navigateToOverview()}
	class="hover:text-text-accent inline-flex h-12 items-center justify-center gap-2 rounded-sm
	 border px-6 transition duration-300"
>
	<span>to project overview</span>
</button>
<br />
<p class="text-center text-3xl">
	Welcome to the api-client <strong class="text-3xl">pre-release</strong>!
</p>
<br />
<p class="text-center text-2xl">Thank you for downloading this amazing piece of software!</p>
<br />
<div class="flex flex-row items-center justify-center gap-2 text-center text-2xl">
	<p class="text-2xl">If you have further questions go here:</p>

	<button
		class="hover:text-text-accent inline-flex items-center justify-center"
		onclick={() => {
			BrowserOpenURL('https://github.com/ThomasBoom89/api-client/issues');
		}}
	>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
		>
			<path stroke="none" d="M0 0h24v24H0z" fill="none" />
			<path d="M12 6h-6a2 2 0 0 0 -2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-6" />
			<path d="M11 13l9 -9" />
			<path d="M15 4h5v5" />
		</svg>
		Github Issues
	</button>
</div>
<div class="flex flex-row items-center justify-center gap-2 text-center text-2xl">
	<p class="text-2xl">If you have further questions go here:</p>

	<button
		class="hover:text-text-accent inline-flex items-center justify-center"
		onclick={() => {
			BrowserOpenURL('https://github.com/ThomasBoom89/api-client/discussions');
		}}
	>
		<svg
			xmlns="http://www.w3.org/2000/svg"
			width="24"
			height="24"
			viewBox="0 0 24 24"
			fill="none"
			stroke="currentColor"
			stroke-width="2"
			stroke-linecap="round"
			stroke-linejoin="round"
		>
			<path stroke="none" d="M0 0h24v24H0z" fill="none" />
			<path d="M12 6h-6a2 2 0 0 0 -2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-6" />
			<path d="M11 13l9 -9" />
			<path d="M15 4h5v5" />
		</svg>
		Github Discussions
	</button>
</div>
