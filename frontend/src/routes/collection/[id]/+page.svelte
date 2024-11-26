<script lang="ts">
	import { frontend } from '../../../lib/wailsjs/go/models';
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';
	import RequestDetail from '../../../lib/RequestDetail.svelte';
	import ResponseDetail from '../../../lib/ResponseDetail.svelte';
	import { Submit } from '../../../lib/wailsjs/go/frontend/Request';
	import Tabs from '../../../lib/Tabs.svelte';
	import { PageTabIndex } from '../../../lib/enums/PageTabIndex';
	import Body from '../../../lib/Body.svelte';

	const id: number = Number($page.params.id);

	const requestStore = getRequestStore();
	const tabs: PageTabIndex[] = [PageTabIndex.Body, PageTabIndex.Params, PageTabIndex.Header, PageTabIndex.Response];
	let selectedRequest = $state(requestStore.getByCollectionId(id)[0]);
	let currentResponse: frontend.RequestResponseDTO = $state(new frontend.RequestResponseDTO());
	let loading = $state(false);
	let currentTab = $state(PageTabIndex.Body);

	function changeSelectedRequest(request: frontend.HttpRequestDto) {
		currentResponse = new frontend.RequestResponseDTO();
		selectedRequest = request;
	}

	let submit = () => {
		currentTab = PageTabIndex.Response;
		loading = true;
		Submit(selectedRequest.id)
			.then((res: frontend.RequestResponseDTO) => {
				currentResponse = res;
			})
			.finally(() => {
				loading = false;
			});
	};

	function changeTab(tabId: PageTabIndex) {
		currentTab = tabId;
	}
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
	<section class="h-full overflow-y-hidden">
		<RequestDetail request={selectedRequest} {submit} />
		<Tabs {tabs} {changeTab} {currentTab} />
		{#if currentTab === PageTabIndex.Body}
			<Body request={selectedRequest} />
		{:else if currentTab === PageTabIndex.Header}
			<p>hallo aus tab 2</p>
		{:else if currentTab === PageTabIndex.Response}
			<ResponseDetail response={currentResponse} {loading}></ResponseDetail>
		{/if}
	</section>
</div>
