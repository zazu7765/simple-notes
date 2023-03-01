<script>
	import { afterUpdate } from 'svelte';
	import { onMount } from 'svelte';
	  onMount(() => {
		window.sessionStorage.setItem('store', $storage);
	window.sessionStorage.setItem('id', $state);

	  });

	//    });
	import { invalidate } from '$app/navigation';
	import NavBar from '../lib/NavBar.svelte';
	import { storage, state } from '$lib/notes';
	


	/** @type {import('./$types').PageData} */
	export let data;
	console.log(data.user.bool + 'wwwww');
	let login = data.user.bool;

	// 	import { onMount } from "svelte";
</script>

{#key login}
	<NavBar {login} />
{/key}
<div class="px-4">
	<div
		class="
          flex
          justify-center
          items-center
          bg-white
          mx-auto
          max-w-2xl
          rounded-lg
          my-16
          p-16
        "
	>
		<h1 class="text-2xl py-5 my-4 block px-5 bg-gray-200 rounded-md text-slate-700 font-medium">I love taking notes🗒️...</h1>
	</div>
</div>

<style>
</style>
