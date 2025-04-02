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
		class="bg-background peer border-background-accent text-text focus:border-text-accent relative h-10 w-full
			appearance-none rounded-sm border px-4 text-sm outline-hidden transition-all focus-visible:outline-hidden
		  focus:focus-visible:outline-hidden"
	>
		<option value="none">none</option>
		<option value="plaintext">text/plain</option>
		<option value="json">application/json</option>
	</select>
	<label
		for="body-type"
		class="text-text before:bg-background peer-focus:text-text peer-disabled:text-text pointer-events-none absolute
		top-2.5 left-2 z-1 px-2 text-sm transition-all peer-valid:-top-2 peer-valid:text-xs peer-focus:-top-2
		peer-focus:text-xs peer-disabled:cursor-not-allowed before:absolute before:top-0 before:left-0 before:z-[-1]
		before:block before:h-full before:w-full before:transition-all peer-disabled:before:bg-transparent"
		>Body Type
	</label>
	<svg
		xmlns="http://www.w3.org/2000/svg"
		class="fill-text pointer-events-none absolute top-2.5 right-2 h-5 w-5 transition-all"
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
	class="h-[70%] w-full"
	placeholder="> your body here <"
	bind:value={request.body.payload}
></textarea>

{#if request.body.type === 'json' && jsonError}
	<p data-testid="json-body-error" class="h-36 text-red-500">{jsonError}</p>
{/if}
