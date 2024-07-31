<script lang="ts">
	import Loader from '$lib/Loader.svelte';

	let { data } = $props();
</script>

<svelte:head>
	<title>Api-Client :: Project Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

{#await data.promise}
	<Loader />
{:then data}
	<h2>Project Overview</h2>
	{#each data.projects as project}
		<section>
			<a href="/project/{project.id}">
				<h3>{project.name}</h3>
			</a>
		</section>
	{/each}
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}
