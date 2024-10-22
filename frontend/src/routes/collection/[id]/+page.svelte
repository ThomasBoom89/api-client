<script lang="ts">
	import { frontend } from '../../../lib/wailsjs/go/models';
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';
	import RequestDetail from '../../../lib/RequestDetail.svelte';

	const requestStore = getRequestStore();
	const id: number = Number($page.params.id);
	let selectedRequest = $state(requestStore.getByCollectionId(id)[0]);

	function changeSelectedRequest(request: frontend.HttpRequestDto) {
		selectedRequest = request;
	}
</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<h2>Collection {id}</h2>
<div class="grid grid-cols-2 h-full">
	<div class="flex flex-col gap-y-2 pr-2 overflow-y-auto">
		{#each requestStore.getByCollectionId(id) as request}
			<section
				class="border-[1px]"
				onclick={() => {
					changeSelectedRequest(request);
				}}
			>
				{#if request.type === 'http'}
					<p>{request.method}>>{request.name}</p>
					<p>{request.url}</p>
				{/if}
			</section>
		{/each}
	</div>
	<section class="">
		<RequestDetail request={selectedRequest} />
	</section>
</div>
