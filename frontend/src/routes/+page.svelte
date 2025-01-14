<script lang="ts">
	import { getNavigationSystem } from '$lib/navigationSystem.svelte.ts';

	const navigationSystem = getNavigationSystem();
</script>

<svelte:head>
	<title>Api-Client :: Startseite</title>
	<meta name="description" content="Startseite" />
</svelte:head>

<button onclick={() => navigationSystem.navigateToOverview()}>to project overview</button>
