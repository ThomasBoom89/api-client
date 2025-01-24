<script lang="ts">
	import { frontend } from '../../../lib/wailsjs/go/models';
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';
	import Request from '../../../lib/Request.svelte';

	const collectionId: number = Number($page.params.id);
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
				  transition-all before:bg-background before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
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
		<div data-testid="requests" class="flex flex-col gap-2">
			{#each requestStore.getByCollectionId(collectionId) as request}
				<div class="flex flex-row w-full gap-x-2 justify-center items-center">
					<div
						data-testid="request"
						class="flex flex-row justify-center w-full items-center gap-2 overflow-hidden rounded shadow-sm p-2
						{request === selectedRequest ? 'shadow-[--color-text-accent]' : 'shadow-[--color-background-accent]'}"
					>
						{#if request.type === 'http'}
							<span class="h-fit rounded border-http border px-2 text-xs text-text">http </span>
						{/if}
						{#if request.id !== requestInEdit}
							<button class=" text-left w-full truncate" onclick={() => changeSelectedRequest(request)}>
								{request.name}
							</button>
							<button
								aria-label="edit"
								onclick={() => {
									requestInEdit = request.id;
								}}
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="24"
									height="24"
									viewBox="0 0 24 24"
									fill="none"
									stroke="currentColor"
									stroke-width="2"
									stroke-linecap="round"
									stroke-linejoin="round"
								>
									<path stroke="none" d="M0 0h24v24H0z" fill="none" />
									<path d="M7 7h-1a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-1" />
									<path d="M20.385 6.585a2.1 2.1 0 0 0 -2.97 -2.97l-8.415 8.385v3h3l8.385 -8.415z" />
									<path d="M16 5l3 3" />
								</svg>
							</button>
						{:else}
							<input type="text" class="border-none w-full" bind:value={request.name} />
							<button
								aria-label="save"
								onclick={() => {
									requestStore.update(request);
									requestInEdit = 0;
								}}
							>
								<svg
									xmlns="http://www.w3.org/2000/svg"
									width="24"
									height="24"
									viewBox="0 0 24 24"
									fill="none"
									stroke="currentColor"
									stroke-width="2"
									stroke-linecap="round"
									stroke-linejoin="round"
								>
									<path stroke="none" d="M0 0h24v24H0z" fill="none" />
									<path d="M6 4h10l4 4v10a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12a2 2 0 0 1 2 -2" />
									<path d="M12 14m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" />
									<path d="M14 4l0 4l-6 0l0 -4" />
								</svg>
							</button>
						{/if}
						<button aria-label="delete" onclick={() => requestStore.delete(request)}>
							<svg
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path stroke="none" d="M0 0h24v24H0z" fill="none" />
								<path d="M4 7l16 0" />
								<path d="M10 11l0 6" />
								<path d="M14 11l0 6" />
								<path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
								<path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
							</svg>
						</button>
					</div>
				</div>
			{/each}
		</div>
	</div>
	<section class="flex flex-col h-full overflow-y-hidden">
		{#if selectedRequest !== undefined}
			<Request request={selectedRequest} {currentResponse}></Request>
		{/if}
	</section>
</div>
