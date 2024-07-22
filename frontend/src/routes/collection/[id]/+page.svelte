<script lang="ts">
	import { page } from '$app/stores';
	import Loader from '$lib/Loader.svelte';

	let { data } = $props();
	const id: number = Number($page.params.id);

</script>

<svelte:head>
	<title>Api-Client :: Collection Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

{#await data.promise}
	<Loader />
{:then data}
	<h2>Collection {id}</h2>
	<section class="flex flex-col gap-y-2">
		{#each data.requests as request}
			<section class="border-[1px]">
				<p>{request.type}</p>
				<p>{request.name}</p>
			</section>
		{/each}
	</section>
{/await}
