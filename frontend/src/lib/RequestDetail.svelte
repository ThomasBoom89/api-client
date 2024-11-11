<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import { getRequestStore } from './requestStore.svelte';

	const requestStore = getRequestStore();

	let { request, submit }: { request: frontend.HttpRequestDto; submit: () => void } = $props();
	let update = (): void => {
		requestStore.update(request);
	};
</script>

<p>{request.name}</p>
<div class="flex flex-row gap-2">
	<select
		bind:value={request.method}
		onchange={() => {
			update();
		}}
	>
		<option value="GET">GET</option>
		<option value="POST">POST</option>
		<option value="PUT">PUT</option>
		<option value="DELETE">DELETE</option>
		<option value="PATCH">PATCH</option>
		<option value="HEAD">HEAD</option>
		<option value="OPTION">OPTION</option>
	</select>
	<input
		class="w-full"
		type="text"
		bind:value={request.url}
		onchange={() => {
			update();
		}}
	/>
	<input type="button" value="submit" onclick={submit} />
</div>
