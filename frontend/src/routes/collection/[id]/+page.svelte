<script lang="ts">
	import { page } from '$app/stores';
	import { getRequestStore } from '../../../lib/requestStore.svelte';

	const requestStore = getRequestStore();
	const id: number = Number($page.params.id);
</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<h2>Collection {id}</h2>
<section class="flex flex-col gap-y-2">
	{#each requestStore.getByCollectionId(id) as request}
		<section class="border-[1px]">
			{#if request.type === 'http'}
				<p>{request.name}</p>
				<p>{request.url}</p>
			{/if}
		</section>
	{/each}
</section>
