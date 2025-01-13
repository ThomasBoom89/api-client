<script lang="ts">
	import { page } from '$app/stores';
	import { getCollectionStore } from '../../../lib/collectionStore.svelte';
	import { getNavigationSystem } from '$lib/navigationSystem.svelte.ts';

	let newCollectionName = $state('');
	let collectionInEdit: number = $state(0);
	const collectionStore = getCollectionStore();
	const navigationSystem = getNavigationSystem();
	const id: number = Number($page.params.id);
</script>

<svelte:head>
	<title>Api-Client :: Project Details</title>
	<meta name="description" content="project details" />
</svelte:head>

<div class="flex flex-col gap-2">
	<div class="relative mt-2 mx-auto">
		<input
			bind:value={newCollectionName}
			id="new-collection"
			type="text"
			placeholder="New Collection"
			class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
		/>
		<label
			for="new-collection"
			class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-[1] px-2 text-xs
				  transition-all before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
			>New Collection
		</label>
		<svg
			onclick={() => {
				collectionStore.create(id, newCollectionName);
				newCollectionName = '';
			}}
			xmlns="http://www.w3.org/2000/svg"
			class="absolute top-2.5 right-4 h-5 w-5 cursor-pointer"
			fill="none"
			id="create-new-collection"
			viewBox="0 0 24 24"
			stroke="currentColor"
			aria-hidden="true"
			aria-label="create icon"
			role="graphics-symbol"
		>
			<path stroke="none" d="M0 0h24v24H0z" fill="none" />
			<path d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0 -18 0" />
			<path d="M9 12h6" />
			<path d="M12 9v6" />
		</svg>
	</div>
	<ol data-testid="collections">
		{#each collectionStore.getByProjectId(id) as collection}
			<li class="flex flex-row gap-2">
				{#if collection.id !== collectionInEdit}
					<button onclick={() => navigationSystem.navigateToCollection(collection.id)}>
						{collection.name}
					</button>
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
</div>
