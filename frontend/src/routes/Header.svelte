<script lang="ts">
	import { page } from '$app/stores';
	import { goto } from '$app/navigation';

	function navigate (page: string = ''): null {
		console.warn(page);
		if (page == undefined) {
			return null;
		}
		goto(page, {});

		return null;
	}
</script>

<header>
	<nav>
		<ul>
			<li aria-current={$page.url.pathname === '/' ? 'page' : undefined}>
				<button onclick="{() => {navigate('/')}}">Home</button>
			</li>
			<li aria-current={$page.url.pathname === '/about' ? 'page' : undefined}>
				<button onclick={() => {navigate('/about')}}>About</button>
			</li>
		</ul>
	</nav>
</header>

<style>
</style>
