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
		httpRequestStore.update(request);
	}
</script>

<div class="relative mt-2">
	<select
		name="body-type"
		id="body-type"
		bind:value={request.body.type}
		onchange={() => {
			update();
		}}
		class="relative w-full h-10 px-4 text-sm transition-all bg-background border rounded-sm outline-hidden appearance-none
		 focus-visible:outline-hidden peer border-background-accent text-text focus:border-text-accent
		  focus:focus-visible:outline-hidden"
	>
		<option value="none">none</option>
		<option value="plaintext">text/plain</option>
		<option value="json">application/json</option>
	</select>
	<label
		for="body-type"
		class="pointer-events-none absolute top-2.5 left-2 z-1 px-2 text-sm text-text transition-all
		before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full before:w-full
		before:bg-background before:transition-all peer-valid:-top-2 peer-valid:text-xs peer-focus:-top-2 peer-focus:text-xs
		 peer-focus:text-text peer-disabled:cursor-not-allowed peer-disabled:text-text peer-disabled:before:bg-transparent"
		>Body Type
	</label>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		class="pointer-events-none absolute top-2.5 right-2 h-5 w-5 fill-text transition-all"
		viewBox="0 0 20 20"
		fill="currentColor"
		aria-labelledby="title-04 description-04"
		role="graphics-symbol"
	>
		<title id="title-04">Arrow Icon</title>
		<desc id="description-04">Arrow icon of the select list.</desc>
		<path
			fill-rule="evenodd"
			d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
			clip-rule="evenodd"
		/>
	</svg>
</div>
<textarea
	oninput={() => {
		update();
	}}
	class:hidden={request.body.type === 'none'}
	class="w-full h-[70%]"
	placeholder="> your body here <"
	bind:value={request.body.payload}
></textarea>

{#if request.body.type === 'json' && jsonError}
	<p data-testid="json-body-error" class="text-red-500 h-36">{jsonError}</p>
{/if}
