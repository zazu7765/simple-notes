<script lang="ts">
	import { toggleBold, toggleItalic, toggleStrike } from '../TipTap/buttonActions';
	import Button from './Button.svelte';
	
  export let editor:any;
</script>

{#if editor}
<div class="fixed-menu">
	<Button {editor} key="bold" on:click={() => toggleBold(editor)} />
	<Button {editor} key="italic" on:click={() => toggleItalic(editor)} />
	<Button {editor} key="strike" on:click={() => toggleStrike(editor)} />
</div>
{/if}

<style>
	.fixed-menu {
    background: #ccc;
		padding: 2px;
		min-height: 36px;
	}
</style>