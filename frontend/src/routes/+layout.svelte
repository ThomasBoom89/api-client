<script>
	import './styles.css';
	import Header from './Header.svelte';
	let { children } = $props();
</script>


<div class="w-[100vw] h-[100vh] bg-gray-600 text-white px-2 pb-2">
	<p>bla</p>
	<Header />
	<main>
		{@render children()}
	</main>
	<footer>
	</footer>
</div>

<style>
</style>
