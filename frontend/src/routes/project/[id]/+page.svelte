<script lang="ts">
	import { page } from '$app/stores';
	import { getCollectionStore } from '../../../lib/collectionStore.svelte';

	const collectionStore = getCollectionStore();
	const id: number = Number($page.params.id);
</script>

<svelte:head>
	<title>Api-Client :: Project Details</title>
	<meta name="description" content="project details" />
</svelte:head>

<h2>Project {id}</h2>
{#each collectionStore.getByProjectId(id) as collection}
	<section>
		<a href="/collection/{collection.id}">
			{collection.name}
		</a>
	</section>
{/each}
