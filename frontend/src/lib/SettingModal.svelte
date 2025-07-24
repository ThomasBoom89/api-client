<script>
	import { getSettingStore } from '$lib/settingStore.svelte.js';

	let { showModal = $bindable() } = $props();
	const settingStore = getSettingStore();
	let dialog = $state(); // HTMLDialogElement
	$effect(() => {
		if (showModal) dialog.showModal();
	});
</script>

<dialog
	bind:this={dialog}
	onclose={() => (showModal = false)}
	onclick={(e) => {
		if (e.target === dialog) dialog.close();
	}}
	class="top-[10%] left-[10%] h-[80%] w-[80%] rounded-2xl border p-2"
>
	<div class="flex flex-col">
		<h1>Request configuration:</h1>
		<hr />
		<br />
		<div class="flex w-fit items-center gap-2">
			<label> Skip Verify TLS: </label>
			<input
				data-testid="skip-tls-verify"
				class="border-background-accent focus:border-text-accent relative m-auto h-6 w-6 rounded-sm border px-4
				text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
				type="checkbox"
				bind:checked={settingStore.currentTlsVerify}
			/>
		</div>
	</div>
</dialog>
