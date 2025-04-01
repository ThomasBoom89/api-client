<script lang="ts">
	import { PageTabIndex } from './enums/PageTabIndex';

	let {
		tabs,
		changeTab,
		currentTab,
	}: {
		tabs: PageTabIndex[];
		changeTab: (tabId: PageTabIndex) => void;
		currentTab: PageTabIndex;
	} = $props();
</script>

<ul data-testid="request-tabs" class="flex items-center" role="tablist">
	{#each tabs as tab}
		<li class="flex-1" role="presentation">
			<button
				onclick={() => changeTab(tab)}
				class:bg-[var(--color-background-accent)]={currentTab === tab}
				class="text-text hover:bg-background-accent -mb-px inline-flex h-12 w-full items-center justify-center
			 		gap-2 rounded-t border-b-[1px] px-6 tracking-wide transition duration-300 focus-visible:outline-hidden"
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
