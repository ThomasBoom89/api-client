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
	<div class="relative mx-auto mt-2">
		<input
			bind:value={newCollectionName}
			id="new-collection"
			type="text"
			placeholder="New Collection"
			class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4 text-sm
			placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
		/>
		<label
			for="new-collection"
			class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
			peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
			peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
			before:block before:h-full before:w-full before:transition-all"
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
	<div data-testid="collections" class="grid-cols-collection-overview grid gap-2">
		{#each collectionStore.getByProjectId(id) as collection}
			<div
				data-testid="collection"
				class="shadow-background-accent flex flex-row gap-2 overflow-hidden rounded-sm p-2 shadow-xs"
			>
				{#if collection.id !== collectionInEdit}
					<button
						class=" w-full truncate text-left"
						onclick={() => navigationSystem.navigateToCollection(collection.id)}
					>
						{collection.name}
					</button>
					<button
						aria-label="edit"
						onclick={() => {
							collectionInEdit = collection.id;
						}}
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
						>
							<path stroke="none" d="M0 0h24v24H0z" fill="none" />
							<path d="M7 7h-1a2 2 0 0 0 -2 2v9a2 2 0 0 0 2 2h9a2 2 0 0 0 2 -2v-1" />
							<path d="M20.385 6.585a2.1 2.1 0 0 0 -2.97 -2.97l-8.415 8.385v3h3l8.385 -8.415z" />
							<path d="M16 5l3 3" />
						</svg>
					</button>
				{:else}
					<input type="text" class="w-full border-none" bind:value={collection.name} />
					<button
						aria-label="save"
						onclick={() => {
							collectionStore.update(collection);
							collectionInEdit = 0;
						}}
					>
						<svg
							xmlns="http://www.w3.org/2000/svg"
							width="24"
							height="24"
							viewBox="0 0 24 24"
							fill="none"
							stroke="currentColor"
							stroke-width="2"
							stroke-linecap="round"
							stroke-linejoin="round"
						>
							<path stroke="none" d="M0 0h24v24H0z" fill="none" />
							<path d="M6 4h10l4 4v10a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12a2 2 0 0 1 2 -2" />
							<path d="M12 14m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" />
							<path d="M14 4l0 4l-6 0l0 -4" />
						</svg>
					</button>
				{/if}
				<button aria-label="delete" onclick={() => collectionStore.delete(collection)}>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
					>
						<path stroke="none" d="M0 0h24v24H0z" fill="none" />
						<path d="M4 7l16 0" />
						<path d="M10 11l0 6" />
						<path d="M14 11l0 6" />
						<path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
						<path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
					</svg>
				</button>
			</div>
		{/each}
		<div style="width: 0px; height: 0px; grid-column: 1 / -1"></div>
	</div>
</div>
