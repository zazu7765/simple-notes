<script>
	import Footer from '$lib/Footer.svelte';
	import { afterUpdate } from 'svelte';
	import { onMount } from 'svelte';
	
	
	let loading = false;
	onMount(() => {

		window.sessionStorage.setItem('store', $storage);
		window.sessionStorage.setItem('id', $state);
		loading = true;
	});


	//    });
	import { invalidate } from '$app/navigation';
	import NavBar from '../lib/NavBar.svelte';
	import { storage, state } from '$lib/notes';

	/** @type {import('./$types').PageData} */
	export let data;

	let login = data.user.bool;

	// 	import { onMount } from "svelte";
</script>

{#if loading == true}
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
			<h1
				class="lg:text-2xl sm:text-xl py-2 my-4 block px-5 text-center rounded-md text-slate-1000 font-medium"
			>
				<b>Take notes with Simple Notes</b><br />
				<p class="text-base mt-3">Organize, manage and always be ahead of your schdule.</p>
			</h1>

			<button
				on:click={() => {
					window.location.assign('/sign');
				}}
				class="md:p-2 py-1 my-2 block  px-1 text-base  rounded-md  text-slate-1000"
				><p class=" w-[50%] m-auto border-2 rounded-lg border-black hover:bg-slate-300">
					Start taking notes now! {import.meta.env.VITE_PUBLIC_ASS} {import.meta.env.VITE_PUBLIC_URL}
				</p></button
			>
		</div>

		<footer class="bottom-0" />
		<div class="sticky top-[95vh]"><Footer /></div>
	</div>
{:else}
	<div class=" absolute min-h-screen top-10 py-10 items-center justify-center">
		<div
			class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
			role="status"
		>
			<span
				class="!absolute py-10 !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
				>Loading...</span
			>
		</div>
	</div>
{/if}

<style>
	:global(body) {
		background-color: #ecf2ff;
	}
</style>
