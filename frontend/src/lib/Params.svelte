<script lang="ts">
	import { getRequestStore } from './requestStore.svelte';
	import { frontend } from './wailsjs/go/models';

	let { request }: { request: frontend.HttpRequestDto } = $props();
	const requestStore = getRequestStore();
	let newParameterKey = $state('');
	let newParameterValue = $state('');
	let buildPreview = (request: frontend.HttpRequestDto): string => {
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

<div>{preview}</div>
<div class="flex flex-row">
	<input bind:value={newParameterKey} type="text" placeholder="Enter parameter name" />
	<input bind:value={newParameterValue} type="text" placeholder="Enter parameter value" />
	<button
		onclick={() => {
			appendParameter();
		}}
		>save
	</button>
</div>
{#each request.parameter as parameter, iter}
	<div class="flex flex-row">
		<input
			bind:value={parameter.key}
			oninput={() => {
				update();
			}}
		/>
		<input
			bind:value={parameter.value}
			oninput={() => {
				update();
			}}
		/>
		<button
			onclick={() => {
				deleteParameter(iter);
			}}
			>delete
		</button>
	</div>
{/each}
