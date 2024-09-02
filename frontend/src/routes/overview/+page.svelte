<script lang="ts">
	import { getProjectStore } from '../../lib/projectStore.svelte';

	const projectStore = getProjectStore();
</script>

<svelte:head>
	<title>Api-Client :: Project Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<h2>Project Overview</h2>
{#each projectStore.projects as project}
	<section>
		<a href="/project/{project.id}">
			<h3>{project.name}</h3>
		</a>
	</section>
{/each}
