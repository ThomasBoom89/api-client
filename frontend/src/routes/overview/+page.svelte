<script lang="ts">
	import { getProjectStore } from '../../lib/projectStore.svelte';
	import { getNavigationSystem } from '$lib/navigationSystem.svelte.ts';

	const projectStore = getProjectStore();
	const navigationSystem = getNavigationSystem();
	let newProjectName: string = $state('');
	let projectInEdit: number = $state(0);
</script>

<svelte:head>
	<title>Api-Client :: Project Overview</title>
	<meta name="description" content="project overview" />
</svelte:head>

<div class="flex flex-row gap-y-2">
	<input type="text" placeholder="insert new project name" bind:value={newProjectName} />
	<button
		type="button"
		onclick={() => {
			projectStore.create(newProjectName);
			newProjectName = '';
		}}
		>create
	</button>
</div>
<ol data-testid="projects">
	{#each projectStore.projects as project}
		<li class="flex flex-row gap-2">
			{#if project.id !== projectInEdit}
				<button onclick={() => navigationSystem.navigateToProject(project.id)}>
					{project.name}
				</button>
				<button onclick={() => (projectInEdit = project.id)}>edit</button>
			{:else}
				<input type="text" bind:value={project.name} />
				<button
					onclick={() => {
						projectStore.update(project);
						projectInEdit = 0;
					}}
					>save
				</button>
			{/if}
			<button onclick={() => projectStore.delete(project)}>delete</button>
		</li>
	{/each}
</ol>
