<script lang="ts">
	import { page } from '$app/stores';
	import { getCollectionStore } from '../../../lib/collectionStore.svelte';
	import { getProjectStore } from '../../../lib/projectStore.svelte';

	let newCollectionName = $state('');
	let collectionInEdit: number = $state(0);
	const collectionStore = getCollectionStore();
	const projectStore = getProjectStore();
	const id: number = Number($page.params.id);
</script>

<svelte:head>
	<title>Api-Client :: Project Details</title>
	<meta name="description" content="project details" />
</svelte:head>

<h2>Project {projectStore.getById(id).name}</h2>
<div class="flex flex-row gap-y-2">
	<input type="text" placeholder="insert new collection name" bind:value={newCollectionName} />
	<button
		onclick={() => {
			collectionStore.create(id, newCollectionName);
			newCollectionName = '';
		}}
		>create
	</button>
</div>
<ol data-testid="collections">
	{#each collectionStore.getByProjectId(id) as collection}
		<li class="flex flex-row gap-2">
			{#if collection.id !== collectionInEdit}
				<a href="/collection/{collection.id}">
					{collection.name}
				</a>
				<button onclick={() => (collectionInEdit = collection.id)}>edit</button>
			{:else}
				<input type="text" bind:value={collection.name} />
				<button
					onclick={() => {
						collectionStore.update(collection);
						collectionInEdit = 0;
					}}
					>save
				</button>
			{/if}
			<button onclick={() => collectionStore.delete(collection)}>delete</button>
		</li>
	{/each}
</ol>
