<script>
	import './styles.css';
	import Header from './Header.svelte';

	let { children } = $props();
</script>
<Header />
<main>
	{@render children()}
</main>
<footer>
</footer>
<style>
</style>
