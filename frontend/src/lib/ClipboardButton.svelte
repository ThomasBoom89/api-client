<script lang="ts">
	import { writable } from 'svelte/store';
	import { fly } from 'svelte/transition';

	export let data: string = '';
	let copied: boolean = false;
	const showNotification = writable(false);
	const showErrorNotification = writable(false);

	async function handleCopy() {
		try {
			if (copied) return;
			await navigator.clipboard.writeText(data);

			copied = true;
			showNotification.set(true);

			setTimeout(() => {
				copied = false;
				showNotification.set(false);
			}, 2000);
		} catch (e) {
			showErrorNotification.set(true);

			setTimeout(() => {
				copied = false;
				showErrorNotification.set(false);
			}, 2000);
		}
	}
</script>

{#if $showNotification}
	<div class="absolute bottom-2 left-2 text-xs text-text-accent" transition:fly={{ y: -10, duration: 300 }}>
		Copied to clipboard!
	</div>
{/if}
{#if $showErrorNotification}
	<div class="absolute bottom-2 left-2 text-xs text-text-response-headers" transition:fly={{ y: -10, duration: 300 }}>
		Failed copying to clipboard!
	</div>
{/if}
<button
	aria-label="clipboard-btn"
	on:click={() => handleCopy()}
	class={copied ? 'text-text-disabled' : 'text-text-accent cursor-pointer'}
	disabled={copied}
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
		<path d="M9 5h-2a2 2 0 0 0 -2 2v12a2 2 0 0 0 2 2h10a2 2 0 0 0 2 -2v-12a2 2 0 0 0 -2 -2h-2"></path>
		<path d="M9 3m0 2a2 2 0 0 1 2 -2h2a2 2 0 0 1 2 2v0a2 2 0 0 1 -2 2h-2a2 2 0 0 1 -2 -2z"></path>
		<path d="M9 12l.01 0"></path>
		<path d="M13 12l2 0"></path>
		<path d="M9 16l.01 0"></path>
		<path d="M13 16l2 0"></path>
	</svg>
</button>
