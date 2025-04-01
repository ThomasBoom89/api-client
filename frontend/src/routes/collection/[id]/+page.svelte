<script lang="ts">
	import { frontend } from '$lib/wailsjs/go/models.ts';
	import { page } from '$app/stores';
	import { getRequestStore } from '$lib/requestStore.svelte';
	import HttpRequest from '$lib/HttpRequest.svelte';
	import Websocket from '$lib/Websocket.svelte';
	import { RequestTypes } from '$lib/enums/RequestTypes.ts';

	const collectionId: number = Number($page.params.id);
	const requestStore = getRequestStore();
	let selectedRequest: frontend.HttpRequestDto | frontend.WebsocketRequestDto | undefined = $state();
	let newRequestName = $state('');
	let newRequestType: RequestTypes = $state(RequestTypes.HTTP);
	let requestInEdit = $state(0);

	let currentResponse: frontend.RequestResponseDTO = $state(new frontend.RequestResponseDTO());

	function changeSelectedRequest(request: frontend.HttpRequestDto | frontend.WebsocketRequestDto) {
		selectedRequest = request;
		currentResponse = new frontend.RequestResponseDTO();
	}
</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<div class="grid h-full grid-cols-[40%_60%] overflow-hidden">
	<div class="flex flex-col gap-y-2 overflow-y-auto pr-2">
		<div class="mx-auto mt-2 flex w-full flex-row">
			<div class="relative min-w-34">
				<select
					name="request-type"
					id="request-type"
					bind:value={newRequestType}
					class="bg-background peer border-background-accent text-text focus:border-text-accent relative h-10 w-full
					appearance-none rounded-sm border px-4 text-sm outline-hidden transition-all focus-visible:outline-hidden
		  		focus:focus-visible:outline-hidden"
				>
					<option value={RequestTypes.HTTP}>HTTP</option>
					<option value={RequestTypes.WEBSOCKET}>WEBSOCKET</option>
				</select>
				<label
					for="request-type"
					class="text-text before:bg-background peer-focus:text-text peer-disabled:text-text pointer-events-none
					absolute top-2.5 left-2 z-1 px-2 text-sm transition-all peer-valid:-top-2 peer-valid:text-xs
					peer-focus:-top-2 peer-focus:text-xs peer-disabled:cursor-not-allowed before:absolute before:top-0
					before:left-0 before:z-[-1] before:block before:h-full before:w-full before:transition-all
					peer-disabled:before:bg-transparent"
				>
					Request type
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
			<div class="relative w-full">
				<input
					bind:value={newRequestName}
					id="new-request"
					type="text"
					placeholder="New Request"
					class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
					text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden
					focus-visible:outline-hidden"
				/>
				<label
					for="new-request"
					class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
					peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
					peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
					before:block before:h-full before:w-full before:transition-all"
					>New Request
				</label>
				<svg
					onclick={() => {
						requestStore.create(collectionId, newRequestName, newRequestType);
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
		</div>
		<div data-testid="requests" class="flex flex-col gap-2">
			{#each requestStore.getByCollectionId(collectionId) as request}
				<div class="flex w-full flex-row items-center justify-center gap-x-2">
					<div
						data-testid="request"
						class="flex w-full flex-row items-center justify-center gap-2 overflow-hidden rounded-sm p-2 shadow-sm
						{request === selectedRequest ? 'shadow-text-accent' : 'shadow-background-accent'}"
					>
						{#if request.type === RequestTypes.HTTP}
							<span class="border-http text-text h-fit rounded-sm border px-2 text-xs">http </span>
						{:else if request.type === RequestTypes.WEBSOCKET}
							<span class="border-websocket text-text h-fit rounded-sm border px-2 text-xs">websocket </span>
						{/if}
						{#if request.id !== requestInEdit}
							<button class=" w-full truncate text-left" onclick={() => changeSelectedRequest(request)}>
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
							<input type="text" class="w-full border-none" bind:value={request.name} />
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
	<section class="flex h-full flex-col overflow-y-hidden">
		{#if selectedRequest !== undefined && selectedRequest.type === RequestTypes.HTTP}
			<HttpRequest request={selectedRequest} {currentResponse}></HttpRequest>
		{:else if selectedRequest !== undefined && selectedRequest.type === RequestTypes.WEBSOCKET}
			<Websocket request={selectedRequest}></Websocket>
		{/if}
	</section>
</div>
