<script lang="ts">
	import { PageTabIndex } from './enums/PageTabIndex';

	let {
		tabs,
		changeTab,
		currentTab,
	}: { tabs: PageTabIndex[]; changeTab: (tabId: PageTabIndex) => void; currentTab: PageTabIndex } = $props();
</script>

<ul data-testid="request-tabs" class="flex items-center" role="tablist">
	{#each tabs as tab}
		<li class="flex-1" role="presentation">
			<button
				onclick={() => changeTab(tab)}
				class:bg-[var(--color-background-accent)]={currentTab === tab}
				class="inline-flex items-center justify-center w-full h-12 gap-2 px-6 -mb-px
			 tracking-wide transition duration-300 rounded-t focus-visible:outline-hidden text-text border-b-[1px]
			    hover:bg-background-accent"
				id="tab-label-1a"
				role="tab"
				aria-setsize="3"
				aria-posinset="1"
				tabindex="0"
				aria-controls="tab-panel-1a"
				aria-selected="true"
			>
				<span>{tab}</span>
			</button>
		</li>
	{/each}
</ul>
