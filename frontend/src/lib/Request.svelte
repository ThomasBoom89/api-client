<script lang="ts">
	import { PageTabIndex } from './enums/PageTabIndex';
	import Tabs from './Tabs.svelte';
	import RequestDetail from './RequestDetail.svelte';
	import ResponseDetail from './ResponseDetail.svelte';
	import Body from './Body.svelte';
	import { Submit } from './wailsjs/go/frontend/Request';
	import { frontend } from './wailsjs/go/models';
	import Params from './Params.svelte';
	import Header from './Header.svelte';

	let { request, currentResponse }: { request: frontend.HttpRequestDto; currentResponse: frontend.RequestResponseDTO } =
		$props();

	let loading = $state(false);
	let currentTab = $state(PageTabIndex.Body);
	const tabs: PageTabIndex[] = [PageTabIndex.Body, PageTabIndex.Parameter, PageTabIndex.Header, PageTabIndex.Response];

	let submit = () => {
		currentTab = PageTabIndex.Response;
		loading = true;
		Submit(request.id)
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

<RequestDetail {request} {submit} />
<Tabs {tabs} {changeTab} {currentTab} />
{#if currentTab === PageTabIndex.Body}
	<Body {request} />
{:else if currentTab === PageTabIndex.Parameter}
	<Params {request} />
{:else if currentTab === PageTabIndex.Header}
	<Header {request} />
{:else if currentTab === PageTabIndex.Response}
	<ResponseDetail response={currentResponse} {loading}></ResponseDetail>
{/if}
