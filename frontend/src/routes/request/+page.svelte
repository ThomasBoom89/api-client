<script lang="ts">
	import { Submit } from '$lib/wailsjs/go/frontend/Request';

	let response = $state('noch nix da');
	let url: HTMLInputElement;
	let submit = () => {
		Submit(url.value).then((res) => {
			response = res;
		});
	};
</script>

<svelte:head>
	<title>Request</title>
	<meta name="description" content="Request page" />
</svelte:head>

<input class="bg-background border-2" type="text" bind:this={url} />
<button onclick={submit}>abschicken</button>
<p>{response}</p>
