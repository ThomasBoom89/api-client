<script lang="ts">
	import { ClipboardSetText } from './wailsjs/runtime/runtime';

	const { data }: { data: string } = $props();
	let copied = $state(false);
	let error = $state(false);

	let handleCopy = async (): Promise<void> => {
		try {
			if (copied) {
				return;
			}
			let res = await ClipboardSetText(data);
			copied = !!res;
			error = !res;
		} catch (err) {
			error = true;
			console.error('Error copying to clipboard: ', err);
		} finally {
			setTimeout(() => {
				error = false;
				copied = false;
			}, 2000);
		}
	};

	let cssClass = (): string => {
		if (error) {
			return 'text-error';
		}

		if (copied) {
			return 'text-ok';
		}

		return 'text-text-highlight cursor-pointer';
	};
</script>

<button aria-label="clipboard-btn" onclick={handleCopy} class={cssClass()} disabled={copied || error}>
	{#if !error}
		{#if copied}
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
				<path d="M9 14l2 2l4 -4"></path>
			</svg>
		{:else}
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
		{/if}
	{:else}
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
			<path d="M10 12l4 4m0 -4l-4 4"></path>
		</svg>
	{/if}
</button>
