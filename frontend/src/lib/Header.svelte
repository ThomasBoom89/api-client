<script lang="ts">
	import { getRequestStore } from './requestStore.svelte';
	import { frontend } from './wailsjs/go/models';

	let { request }: { request: frontend.HttpRequestDto } = $props();
	const requestStore = getRequestStore();
	let newHeaderKey = $state('');
	let newHeaderValue = $state('');

	let update = (): void => {
		requestStore.update(request);
	};

	let deleteHeader = (index: number): void => {
		request.header.splice(index, 1);
		requestStore.update(request);
	};

	let appendHeader = (): void => {
		let header = new frontend.HttpRequestHeaderDto();
		header.key = newHeaderKey;
		header.value = newHeaderValue;
		header.httpRequestID = request.id;
		if (request.header === undefined || request.header === null) {
			request.header = [];
		}
		request.header.push(header);
		requestStore.update(request);
		newHeaderKey = '';
		newHeaderValue = '';
	};
</script>

<div class="flex flex-row">
	<input bind:value={newHeaderKey} type="text" placeholder="Enter header key" />
	<input bind:value={newHeaderValue} type="text" placeholder="Enter header value" />
	<button
		onclick={() => {
			appendHeader();
		}}
		>save
	</button>
</div>
<ul data-testid="request-headers">
	{#each request.header as header, iter}
		<li class="flex flex-row">
			<input
				bind:value={header.key}
				oninput={() => {
					update();
				}}
			/>
			<input
				bind:value={header.value}
				oninput={() => {
					update();
				}}
			/>
			<button
				onclick={() => {
					deleteHeader(iter);
				}}
				>delete
			</button>
		</li>
	{/each}
</ul>
