<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import Loader from './Loader.svelte';

	let { response, loading }: { response: frontend.RequestResponseDTO; loading: boolean } = $props();
	let isJsonResponse: boolean = $derived.by(() => {
		try {
			JSON.parse(response.responseBody);
		} catch (e) {
			return false;
		}

		return true;
	});
</script>

{#if loading}
	<Loader></Loader>
{:else if response.error !== ''}
	<p data-testid="response-error" class="text-red-700">{response.error}</p>
{:else if response.elapsedTime === undefined}{:else}
	<div class="flex flex-col h-full overflow-y-auto overflow-x-hidden p-2 border">
		<h4 class="text-text-highlight">Current Response:</h4>
		<p class="break-words mr-2">{response.method} {response.url}</p>
		<hr class="my-2 mr-2" />
		<h4 class="text-text-highlight">Status-Code:</h4>
		<p>{response.statusCode}</p>
		<hr class="my-2 mr-2" />
		<h4 class="text-text-highlight">Elapsed Time:</h4>
		<p>{response.elapsedTime}</p>
		<hr class="my-2 mr-2" />
		<h4 class="text-text-highlight">Send Header:</h4>
		{#each Object.entries(response.sendHeader) as [key, value]}
			<p><span class="text-text-response-headers">{key}</span> {value}</p>
		{/each}
		<hr class="my-2 mr-2" />
		<h4 class="text-text-highlight">Received Header:</h4>
		{#each Object.entries(response.receivedHeader) as [key, value]}
			<p class="break-words mr-2"><span class="text-text-response-headers">{key}</span> {value}</p>
		{/each}
		{#if response.error !== ''}
			<hr class="my-2 mr-2" />
			<h4 class="text-text-highlight">Error:</h4>
			<p>{response.error}</p>
		{/if}
		<hr class="my-2 mr-2" />
		<h4 class="text-text-highlight">Payload:</h4>
		{#if isJsonResponse}
			<div class="mr-2">
				<pre class="whitespace-pre-wrap overflow-x-auto">{response.responseBody}</pre>
			</div>
		{:else}
			<div class="break-words mr-2">{response.responseBody}</div>
		{/if}
	</div>
{/if}
