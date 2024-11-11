<script lang="ts">
	import { frontend } from '../../../lib/wailsjs/go/models';
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';
	import RequestDetail from '../../../lib/RequestDetail.svelte';
	import ResponseDetail from '../../../lib/ResponseDetail.svelte';
	import { Submit } from '../../../lib/wailsjs/go/frontend/Request';

	const id: number = Number($page.params.id);

	const requestStore = getRequestStore();

	let selectedRequest = $state(requestStore.getByCollectionId(id)[0]);
	let currentResponse: frontend.RequestResponseDTO = $state(new frontend.RequestResponseDTO());
	let loading = $state(false);

	function changeSelectedRequest(request: frontend.HttpRequestDto) {
		currentResponse = new frontend.RequestResponseDTO();
		selectedRequest = request;
	}

	let submit = () => {
		loading = true;
		Submit(selectedRequest.id).then((res: frontend.RequestResponseDTO) => {
			loading = false;
			currentResponse = res;
		});
	};
</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<h2>Collection {id}</h2>
<div class="grid grid-cols-[30%_70%] h-full">
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
	<section class="h-full overflow-hidden">
		<RequestDetail request={selectedRequest} {submit} />
		<ResponseDetail response={currentResponse} {loading}></ResponseDetail>
	</section>
</div>
