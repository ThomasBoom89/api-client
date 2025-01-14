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
<div class="flex flex-col gap-2">
	<div class="relative mt-2 mx-auto">
		<input
			bind:value={newProjectName}
			id="new-project"
			type="text"
			placeholder="New Project"
			class="relative w-full h-10 px-4 text-sm placeholder-transparent transition-all border rounded outline-none focus-visible:outline-none peer
				 border-background-accent focus:border-text-accent focus:outline-none"
		/>
		<label
			for="new-project"
			class="cursor-text peer-focus:cursor-default peer-autofill:-top-2 absolute left-2 -top-2 z-[1] px-2 text-xs
				  transition-all before:bg-background before:absolute before:top-0 before:left-0 before:z-[-1] before:block before:h-full
				  before:w-full before:transition-all peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm
				  peer-focus:-top-2 peer-focus:text-xs"
			>New Project
		</label>
		<svg
			onclick={() => {
				projectStore.create(newProjectName);
				newProjectName = '';
			}}
			xmlns="http://www.w3.org/2000/svg"
			class="absolute top-2.5 right-4 h-5 w-5 cursor-pointer"
			fill="none"
			id="create-new-project"
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
	<div data-testid="projects" class="grid grid-cols-projectOverview gap-2">
		{#each projectStore.projects as project}
			<div
				data-testid="project"
				class="flex flex-row gap-2 overflow-hidden rounded shadow-sm shadow-[--color-background-accent] p-2"
			>
				{#if project.id !== projectInEdit}
					<button class=" text-left w-full truncate" onclick={() => navigationSystem.navigateToProject(project.id)}>
						{project.name}
					</button>
					<button
						aria-label="edit"
						onclick={() => {
							projectInEdit = project.id;
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
					<input type="text" class="border-none w-full" bind:value={project.name} />
					<button
						aria-label="save"
						onclick={() => {
							projectStore.update(project);
							projectInEdit = 0;
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
				<button aria-label="delete" onclick={() => projectStore.delete(project)}>
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
