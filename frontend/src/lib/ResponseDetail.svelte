<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import Loader from './Loader.svelte';

	let { response, loading }: { response: frontend.RequestResponseDTO; loading: boolean } = $props();
</script>

{#if loading}
	<Loader></Loader>
{:else if response.elapsedTime === undefined}{:else}
	<div class="flex flex-col gap-2 h-full">
		<h3>Current Response {response.method} {response.url}</h3>
		{#if response.error !== ''}
			<div>Error: {response.error}</div>
		{/if}
		<div>Elapsed Time: {response.elapsedTime}</div>
		<div class="overflow-y-auto">{response.responseBody}</div>
	</div>
{/if}
