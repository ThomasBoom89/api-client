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
		<div class="relative mt-2 mx-auto">
			<input
				bind:value={newRequestName}
				id="new-request"
				type="text"
				placeholder="New Request"
				class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
			/>
			<label
				for="new-request"
				class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-[1] px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
				>New Request
			</label>
			<svg
				onclick={() => {
					requestStore.create(collectionId, newRequestName);
					newRequestName = '';
				}}
				xmlns="http://www.w3.org/2000/svg"
				class="absolute top-2.5 right-4 h-5 w-5 cursor-pointer"
				fill="none"
				id="create-new-request"
				viewBox="0 0 24 24"
				stroke="currentColor"
				aria-hidden="true"
				aria-label="create icon"
				role="graphics-symbol"
			>
				<path stroke="none" d="M0 0h24v24H0z" fill="none" />
				<path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0 -18 0" />
				<path d="M9 12h6" />
				<path d="M12 9v6" />
			</svg>
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
