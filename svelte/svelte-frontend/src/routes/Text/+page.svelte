<script lang="ts">
	import Editor from '@tinymce/tinymce-svelte';

	let value = 'some content';
</script>

<main>
	<Editor
		apiKey="ro357dg4ght7s6z2r3m0f2uslrfl9i9dfogk5mvn3dj2xhr5"
		channel="6"
		id="uuid"
		scriptSrc="undefined"
		modelEvents="input change undo redo"
		value="value"
		text="readonly-text-output"
	/>
	<div>{@html value}</div>
	<textarea bind:value />
</main>
