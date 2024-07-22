<script lang="ts">

	import { page } from '$app/stores';
	import Loader from '$lib/Loader.svelte';

	let { data } = $props();
	const id: number = Number($page.params.id);
</script>

<svelte:head>
	<title>Api-Client :: Project Details</title>
	<meta name="description" content="project details" />
</svelte:head>

{#await data.promise}
	<div class="w-full h-full">
		<Loader />
	</div>
{:then data}
	<h2>Project {id}</h2>
	{#each data.collections as collection}
		<section>
			<a href="/collection/{collection.id}">
				{collection.name}
			</a>
		</section>
	{/each}
{/await}
