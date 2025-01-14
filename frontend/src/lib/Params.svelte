<script lang="ts">
	import { getRequestStore } from './requestStore.svelte';
	import { frontend } from './wailsjs/go/models';

	let { request }: { request: frontend.HttpRequestDto } = $props();
	const requestStore = getRequestStore();
	let newParameterKey = $state('');
	let newParameterValue = $state('');
	let buildPreview = (request: frontend.HttpRequestDto): string => {
		if (request.parameter === null || request.parameter === undefined) {
			return request.url;
		}
		let param: { [key: string]: string } = {};
		for (const parameter of request.parameter) {
			param[parameter.key] = parameter.value;
		}
		const searchParams = new URLSearchParams(param);
		if (searchParams.toString() === '') {
			return request.url;
		}

		return request.url + '?' + searchParams.toString();
	};
	let preview = $derived(buildPreview(request));

	let update = (): void => {
		requestStore.update(request);
	};

	let deleteParameter = (index: number): void => {
		request.parameter.splice(index, 1);
		requestStore.update(request);
	};

	let appendParameter = (): void => {
		if (newParameterKey === '' || newParameterValue === '') {
			return;
		}
		let parameter = new frontend.HttpRequestParameterDto();
		parameter.key = newParameterKey;
		parameter.value = newParameterValue;
		parameter.httpRequestID = request.id;
		if (request.parameter === undefined || request.parameter === null) {
			request.parameter = [];
		}
		request.parameter.push(parameter);
		requestStore.update(request);
		newParameterKey = '';
		newParameterValue = '';
	};
</script>

<div class="flex flex-col h-full">
	<div data-testid="request-parameter-preview" class="pr-4">{preview}</div>
	<div class="flex flex-row gap-1 pr-4 mt-4">
		<div class="relative mx-auto w-full">
			<input
				bind:value={newParameterKey}
				id="new-parameter-name"
				type="text"
				placeholder="Parameter name"
				class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
			/>
			<label
				for="new-parameter-name"
				class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-[1] px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:bg-background before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
				>Parameter Name
			</label>
		</div>
		<div class="relative mx-auto w-full">
			<input
				bind:value={newParameterValue}
				id="new-parameter-value"
				type="text"
				placeholder="Request Url"
				class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
			/>
			<label
				for="new-parameter-value"
				class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-[1] px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:bg-background before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
				>Parameter Value
			</label>
		</div>
		<button
			aria-label="save"
			onclick={() => {
				appendParameter();
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
	<div data-testid="request-parameters" class="flex flex-col overflow-y-auto max-h-full pr-4 mt-4">
		{#each request.parameter as parameter, iter}
			<div class="flex flex-row gap-1">
				<input
					bind:value={parameter.key}
					oninput={() => {
						update();
					}}
					type="text"
					class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
				/>
				<input
					bind:value={parameter.value}
					oninput={() => {
						update();
					}}
					type="text"
					class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
				/>
				<button aria-label="delete" onclick={() => deleteParameter(iter)} class="h-10">
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
