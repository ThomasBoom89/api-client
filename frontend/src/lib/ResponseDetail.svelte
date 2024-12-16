<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import Loader from './Loader.svelte';

	let { response, loading }: { response: frontend.RequestResponseDTO; loading: boolean } = $props();
</script>

{#if loading}
	<Loader></Loader>
{:else if response.error !== ''}
	<p data-testid="response-error" class="text-red-700">{response.error}</p>
{:else if response.elapsedTime === undefined}{:else}
	<div class="flex flex-col h-full overflow-y-auto overflow-x-hidden p-2 border">
		<h3>Current Response {response.method} {response.url}</h3>
		<h4>Send Header:</h4>
		{#each Object.entries(response.sendHeader) as [key, value]}
			<p>{key} {value}</p>
		{/each}
		<h4>Status-Code: {response.statusCode}</h4>
		<h4>Received Header:</h4>
		{#each Object.entries(response.receivedHeader) as [key, value]}
			<p>{key} {value}</p>
		{/each}
		{#if response.error !== ''}
			<div>Error: {response.error}</div>
		{/if}
		<div>Elapsed Time: {response.elapsedTime}</div>
		<div>{response.responseBody}</div>
	</div>
{/if}
