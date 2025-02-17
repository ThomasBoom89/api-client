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

<div class="flex flex-col h-full">
	<div class="flex flex-row gap-1 pr-4 mt-4">
		<div class="relative mx-auto w-full">
			<input
				bind:value={newHeaderKey}
				id="new-header-key"
				type="text"
				placeholder="New Header Key"
				class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded-sm outline-hidden focus-visible:outline-hidden peer
				 border-background-accent focus:border-text-accent focus:outline-hidden"
			/>
			<label
				for="new-header-key"
				class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-1 px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:bg-background before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
				>New Header Key
			</label>
		</div>
		<div class="relative mx-auto w-full">
			<input
				bind:value={newHeaderValue}
				id="new-header-value"
				type="text"
				placeholder="New Header Value"
				class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded-sm outline-hidden focus-visible:outline-hidden peer
				 border-background-accent focus:border-text-accent focus:outline-hidden"
			/>
			<label
				for="new-header-value"
				class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-1 px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:bg-background before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
				>New Header Value
			</label>
		</div>
		<button
			aria-label="save"
			onclick={() => {
				appendHeader();
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
	</div>
	<div data-testid="request-headers" class="pr-4 overflow-y-auto flex flex-col max-h-full mt-4">
		{#each request.header as header, iter}
			<div class="flex flex-row gap-1">
				<input
					bind:value={header.key}
					oninput={() => {
						update();
					}}
					type="text"
					class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded-sm outline-hidden focus-visible:outline-hidden peer
				 border-background-accent focus:border-text-accent focus:outline-hidden"
				/>
				<input
					bind:value={header.value}
					oninput={() => {
						update();
					}}
					type="text"
					class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded-sm outline-hidden focus-visible:outline-hidden peer
				 border-background-accent focus:border-text-accent focus:outline-hidden"
				/>
				<button aria-label="delete" onclick={() => deleteHeader(iter)} class="h-10">
					<svg
						class="h-5"
						xmlns="http://www.w3.org/2000/svg"
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
		{/each}
	</div>
</div>
