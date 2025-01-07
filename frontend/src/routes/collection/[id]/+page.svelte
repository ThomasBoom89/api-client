<script lang="ts">
	import { frontend } from '../../../lib/wailsjs/go/models';
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';
	import { getCollectionStore } from '../../../lib/collectionStore.svelte';
	import Request from '../../../lib/Request.svelte';

	const collectionId: number = Number($page.params.id);
	const collectionStore = getCollectionStore();
	const requestStore = getRequestStore();
	let selectedRequest = $state<frontend.HttpRequestDto>();
	let newRequestName = $state('');
	let requestInEdit = $state(0);
	let currentResponse: frontend.RequestResponseDTO = $state(new frontend.RequestResponseDTO());

	function changeSelectedRequest(request: frontend.HttpRequestDto) {
		selectedRequest = request;
		currentResponse = new frontend.RequestResponseDTO();
	}
</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<div class="grid grid-cols-[40%_60%] overflow-hidden h-full">
	<div class="flex flex-col gap-y-2 pr-2 overflow-y-auto">
		<div class="flex flex-row gap-y-2">
			<input type="text" placeholder="insert new request name" bind:value={newRequestName} />
			<button
				onclick={() => {
					requestStore.create(collectionId, newRequestName);
					newRequestName = '';
				}}
				>create
			</button>
		</div>
		<ol data-testid="requests">
			{#each requestStore.getByCollectionId(collectionId) as request}
				<li class="border-[1px]">
					<div class="flex flex-row gap-x-2">
						{#if request.type === 'http'}
							{request.method}>>
						{/if}
						{#if request.id !== requestInEdit}
							<button onclick={() => changeSelectedRequest(request)} class="w-full text-left">
								{request.name}
							</button>
							<button onclick={() => (requestInEdit = request.id)}>edit</button>
						{:else}
							<input bind:value={request.name} />
							<button
								onclick={() => {
									requestStore.update(request);
									requestInEdit = 0;
								}}
								>save
							</button>
						{/if}
						<button onclick={() => requestStore.delete(request)}>delete</button>
					</div>
				</li>
			{/each}
		</ol>
	</div>
	<section class="flex flex-col h-full overflow-y-hidden">
		{#if selectedRequest !== undefined}
			<Request request={selectedRequest} {currentResponse}></Request>
		{/if}
	</section>
</div>
