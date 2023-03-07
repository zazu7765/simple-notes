<script>
	import Footer from "$lib/Footer.svelte"
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
<div class="px-4 min-h-screen">
	<div
		class="
          grid
          justify-center
          items-center
		  content-center 
          mx-auto
          max-w-2xl
          rounded-lg
          mt-16
          
        "
	>
		<h1 class="lg:text-2xl sm:text-xl py-2 my-4 block px-5 text-center rounded-md text-slate-1000 font-medium"><b>Take notes with Simple Notes</b><br>
			<p class="text-base mt-3">Organize, manage and always be ahead of your schdule.</p>
		</h1>
		
	</div>
	
	<footer class="bottom-0" >
		
		</footer>
		<div class="sticky top-[95vh]"><Footer></Footer></div>
</div>


<style>
	:global(body){
		background-color: #ECF2FF;
	}
</style>
