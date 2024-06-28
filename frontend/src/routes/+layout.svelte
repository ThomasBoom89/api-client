<script>
	import './styles.css';
	import Header from './Header.svelte';
	import useFetchInitial from './useFetchInitial.svelte.ts';

	const initialState = useFetchInitial();
	let { children } = $props();
</script>

{#if initialState.isLoading}
	<section class="w-[100vw] h-[100vh] flex flex-col items-center justify-center overflow-hidden bg-neutral-800">
		<span class="animate-ping w-full h-full rounded-full opacity-50 bg-blue-200"></span>
		<p class="absolute top-[50%] text-white">loading application...</p>
	</section>
{:else if initialState.error}
	<p>Application couldn't start! :(</p>
{:else}
	<div class="w-[100vw] h-[100vh] bg-background text-text p-2">
		<Header />
		<main>
			{@render children()}
		</main>
		<footer>
		</footer>
	</div>
{/if}
<style>
</style>
