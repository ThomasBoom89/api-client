<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import Loader from './Loader.svelte';
	import ClipboardButton from './ClipboardButton.svelte';

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
		<div class="flex justify-between mr-2">
			<h4 class="text-text-highlight">Status-Code:</h4>
			<ClipboardButton data={response.statusCode.toFixed(0)} />
		</div>
		<p>{response.statusCode}</p>
		<hr class="my-2 mr-2" />
		<div class="flex justify-between mr-2">
			<h4 class="text-text-highlight">Elapsed Time:</h4>
			<ClipboardButton data={response.elapsedTime} />
		</div>
		<p>{response.elapsedTime}</p>
		<hr class="my-2 mr-2" />
		<div class="flex justify-between mr-2">
			<h4 class="text-text-highlight">Send Header:</h4>
			<ClipboardButton data={JSON.stringify(response.sendHeader)} />
		</div>
		{#each Object.entries(response.sendHeader) as [key, value]}
			<p><span class="text-text-response-headers">{key}</span> {value}</p>
		{/each}
		<hr class="my-2 mr-2" />
		<div class="flex justify-between mr-2">
			<h4 class="text-text-highlight">Received Header:</h4>
			<ClipboardButton data={JSON.stringify(response.receivedHeader)} />
		</div>
		{#each Object.entries(response.receivedHeader) as [key, value]}
			<p class="mr-2 break-words">
				<span class="text-text-response-headers">{key}</span>
				{value}
			</p>
		{/each}
		{#if response.error !== ''}
			<hr class="my-2 mr-2" />
			<h4 class="text-text-highlight">Error:</h4>
			<p>{response.error}</p>
		{/if}
		<hr class="my-2 mr-2" />
		<div class="flex justify-between mr-2">
			<h4 class="text-text-highlight">Payload:</h4>
			<ClipboardButton data={response.responseBody} />
		</div>
		{#if isJsonResponse}
			<div class="mr-2">
				<pre class="overflow-x-auto whitespace-pre-wrap">{response.responseBody}</pre>
			</div>
		{:else}
			<div class="mr-2 break-words">{response.responseBody}</div>
		{/if}
	</div>
{/if}
