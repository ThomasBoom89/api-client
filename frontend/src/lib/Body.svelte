<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import { getRequestStore } from './requestStore.svelte';

	let { request }: { request: frontend.HttpRequestDto } = $props();
	const httpRequestStore = getRequestStore();
	let jsonError: string = $state('');

	let validate_json = () => {
		if (request.body.type === 'json') {
			try {
				JSON.parse(request.body.payload);
			} catch (error) {
				jsonError = (error as SyntaxError).message;
			}
		}
	};

	$effect(() => {
		jsonError = '';
		validate_json();
	});

	function update(): void {
		jsonError = '';
		if (request.body.type === 'json') {
			validate_json();
			if (jsonError.length > 0) {
				return;
			}
		}
		httpRequestStore.update(request).then((result) => {
			request = result;
		});
	}
</script>

<div class="flex gap-2">
	<label for="type-selector" class="">select body type:</label>
	<select
		id="type-selector"
		class="w-full flex-1"
		bind:value={request.body.type}
		onchange={() => {
			update();
		}}
	>
		<option value="none">kein Body</option>
		<option value="plaintext">text/plain</option>
		<option value="json">application/json</option>
	</select>
</div>
<textarea
	oninput={() => {
		update();
	}}
	class="w-full h-[70%]"
	placeholder="write your body"
	bind:value={request.body.payload}
></textarea>

{#if request.body.type === 'json' && jsonError}
	<p class="text-red-500 h-36">{jsonError}</p>
{/if}
