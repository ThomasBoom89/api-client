<svelte:head>
    <title>About</title>
    <meta name="description" content="About this app"/>
</svelte:head>

<script lang="ts">
    import {Greet} from "$lib/wailsjs/go/main/App";

    let name = $state("peter")

    function onclick() {
        Greet("peter").then((greet: string) => {
            name = greet;
        })
    }
</script>

<div class="text-column">
    <h1>About this app</h1>
    <button {onclick}>Test</button>
    <p>{name}</p>
</div>
