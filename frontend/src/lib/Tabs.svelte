<script lang="ts">
	import { PageTabIndex } from './enums/PageTabIndex';

	let {
		tabs,
		changeTab,
		currentTab,
	}: { tabs: PageTabIndex[]; changeTab: (tabId: PageTabIndex) => void; currentTab: PageTabIndex } = $props();
</script>

<ul data-testid="request-tabs" class="flex justify-between w-full">
	{#each tabs as tab}
		<li
			class:bg-[var(--color-background-accent)]={currentTab === tab}
			class="cursor-pointer"
			onclick={() => changeTab(tab)}
		>
			{tab}
		</li>
	{/each}
</ul>
