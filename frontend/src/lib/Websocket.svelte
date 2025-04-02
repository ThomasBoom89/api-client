<script lang="ts">
	import { frontend } from '$lib/wailsjs/go/models.ts';
	import { getRequestStore } from '$lib/requestStore.svelte.ts';
	import { getWebsocketStore } from '$lib/websocketStore.svelte.ts';
	import { WebsocketMessageType } from '$lib/enums/WebsocketMessageType.ts';

	let { request }: { request: frontend.WebsocketRequestDto } = $props();
	let message: string = $state('');
	let websocketStore = getWebsocketStore();
	let websocketState = $derived(websocketStore.getStateById(request.id));
	let requestStore = getRequestStore();
	let update = () => {
		requestStore.update(request);
	};
	let connect = () => {
		if (request.url === '') {
			return;
		}
		websocketStore.connect(request);
	};

	let disconnect = () => {
		websocketStore.disconnect(request);
	};

	let sendMessage = () => {
		if (message === '') {
			return;
		}
		websocketStore.sendMessage(request, message);
		message = '';
	};
</script>

<div class="mt-2 flex h-full w-full flex-col gap-2">
	<div class="flex w-full flex-row">
		<div class="relative w-full">
			<input
				id="request-url"
				type="url"
				placeholder="Request Url"
				bind:value={request.url}
				oninput={() => {
					update();
				}}
				class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
				text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
			/>
			<label
				for="request-url"
				class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
				  peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
				  peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
					before:block before:h-full before:w-full before:transition-all"
				>Request URL
			</label>
		</div>
		{#if websocketState.connected === true}
			<button
				id="disconnect"
				onclick={() => disconnect()}
				class="inline-flex h-10 min-w-32 items-center justify-center gap-2 rounded-sm border bg-red-700 px-6 transition
				 duration-300 hover:bg-red-600"
			>
				<span>disconnect</span>
			</button>
		{:else}
			<button
				id="connect"
				onclick={() => connect()}
				class="inline-flex h-10 min-w-32 items-center justify-center gap-2 rounded-sm border bg-green-700 px-6
				 transition duration-300 hover:bg-green-600"
			>
				<span>connect</span>
			</button>
		{/if}
	</div>
	<div class="flex w-full flex-row">
		<div class="relative w-full">
			<input
				id="message"
				type="url"
				placeholder="Request Url"
				bind:value={message}
				class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
				 text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
			/>
			<label
				for="message"
				class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
				  peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
				  peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
				  before:block before:h-full before:w-full before:transition-all"
				>Message
			</label>
		</div>
		<button
			id="send-message"
			onclick={() => sendMessage()}
			class="hover:text-text-accent inline-flex h-10 items-center justify-center gap-2 rounded-sm
	 border px-6 transition duration-300"
		>
			<span>send</span>
		</button>
	</div>
	<div id="messages" class="flex h-full flex-col gap-2 overflow-y-auto border">
		{#each websocketState.messages as message}
			<div class="flex flex-row">
				{#if message.type === WebsocketMessageType.Incoming}
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
						class="text-red-600"
					>
						<path stroke="none" d="M0 0h24v24H0z" fill="none" />
						<path d="M12 5l0 14" />
						<path d="M16 15l-4 4" />
						<path d="M8 15l4 4" />
					</svg>
				{:else}
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
						class="text-green-600"
					>
						<path stroke="none" d="M0 0h24v24H0z" fill="none" />
						<path d="M12 5l0 14" />
						<path d="M16 9l-4 -4" />
						<path d="M8 9l4 -4" />
					</svg>
				{/if}
				<p class="ml-2">{message.value}</p>
			</div>
		{/each}
	</div>
</div>
