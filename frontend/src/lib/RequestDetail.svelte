<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import { getRequestStore } from './requestStore.svelte';

	const requestStore = getRequestStore();

	let { request, submit }: { request: frontend.HttpRequestDto; submit: () => void } = $props();
	let update = (): void => {
		requestStore.update(request);
	};
</script>

<div class="flex flex-row gap-2 my-2">
	<div class="relative min-w-28">
		<select
			name="request-method"
			id="request-method"
			bind:value={request.method}
			onchange={() => {
				update();
			}}
			class="relative w-full h-10 px-4 text-sm transition-all bg-background border rounded-sm outline-hidden appearance-none
		 focus-visible:outline-hidden peer border-background-accent text-text focus:border-text-accent
		  focus:focus-visible:outline-hidden"
		>
			<option value="GET">GET</option>
			<option value="POST">POST</option>
			<option value="PUT">PUT</option>
			<option value="DELETE">DELETE</option>
			<option value="PATCH">PATCH</option>
			<option value="HEAD">HEAD</option>
			<option value="OPTION">OPTION</option>
		</select>
		<label
			for="request-method"
			class="pointer-events-none absolute top-2.5 left-2 z-1 px-2 text-sm text-text transition-all
		before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full before:w-full
		before:bg-background before:transition-all peer-valid:-top-2 peer-valid:text-xs peer-focus:-top-2 peer-focus:text-xs
		 peer-focus:text-text peer-disabled:cursor-not-allowed peer-disabled:text-text peer-disabled:before:bg-transparent"
		>
			Method type
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
	<div class="relative mx-auto w-full">
		<input
			bind:value={request.url}
			oninput={() => {
				update();
			}}
			id="request-url"
			type="url"
			placeholder="Request Url"
			class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded-sm outline-hidden focus-visible:outline-hidden peer
				 border-background-accent focus:border-text-accent focus:outline-hidden"
		/>
		<label
			for="request-url"
			class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-1 px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:bg-background before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
			>Request URL
		</label>
		<svg
			onclick={submit}
			xmlns="http://www.w3.org/2000/svg"
			class="absolute top-1 right-3 h-7 w-7 cursor-pointer"
			fill="none"
			id="run-request-icon"
			viewBox="0 0 16 20"
			stroke="currentColor"
			aria-hidden="true"
			aria-label="create icon"
			role="graphics-symbol"
		>
			<path stroke="none" d="M0 0h24v24H0z" fill="none" />
			<path d="M13 4m-1 0a1 1 0 1 0 2 0a1 1 0 1 0 -2 0" />
			<path d="M4 17l5 1l.75 -1.5" />
			<path d="M15 21l0 -4l-4 -3l1 -6" />
			<path d="M7 12l0 -3l5 -1l3 3l3 1" />
		</svg>
	</div>
</div>
