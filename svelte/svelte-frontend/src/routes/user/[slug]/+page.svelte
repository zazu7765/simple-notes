<script module lang="ts">
	import { redirect } from '@sveltejs/kit';
	import Sidebar from '../../../lib/Sidebar.svelte';
	import { write } from '../../../lib/notes';
	import { browser } from '$app/env';
	export let data: any;
	let login = data.bool;
	$write = data['responseNote'];
	let pp;
	write.subscribe((value) => {
		pp = value;
	});
	console.log(pp['Data'] + ' note');
	// console.log(data['responseNote'])

	let num = 5;
	console.log(data);
	let content = data['response']['Data'];
	let arr: Array<string>[] = [];
	for (const item in content) {
		arr.push([content[item]['ID'], content[item]['Title']]);
	}
	console.log(arr);
	let go = () => {
		throw redirect(302, '/login');
	};
	let ww = "w-0";
	let blur = "";
	function openNav() {
		ww = "w-64"
		blur = "blur-sm";
	}
	function closeNav() {
	  ww = "w-0";
	  blur = '';
	
	}
	let screenWidth;
	if (browser){
        innerWidth = window.innerWidth;
    }

</script>


<!-- <svelte:window bind:innerWidth={screenWidth} />
{#if screenWidth < 600}
	
{:else}
 
{/if} -->
<button on:click={()=>{
	openNav();
}} type="button" class="Class

z-50 inline-flex  items-center p-6 mt-2 ml-3 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
	<span class="sr-only">Open sidebar</span>
	<svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
	   <path clip-rule="evenodd" fill-rule="evenodd" d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"></path>
	</svg>
 </button>
 
 <aside class="ease-in-out duration-300 fixed top-0 left-0 z-40 {ww} h-screen " aria-label="Sidebar">
	<div class="h-full py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
	   <ul class="space-y-2">
		
			<button  class="flex items-center pl-2.5 mb-5" on:click={()=>{closeNav()}}>
				<img
					src="https://flowbite.com/docs/images/logo.svg"
					class="h-6 mr-3 sm:h-7"
					alt="Flowbite Logo"
				/>
				<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
					>Simple Notes
					</span
				>
			</button>
			<ul class="space-y-2">
				<li>
					<a
						href="#"
						class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
					>
					 <svg class="h-10 w-10" fill="#000000" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"><g id="SVGRepo_bgCarrier" stroke-width="0"></g><g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g><g id="SVGRepo_iconCarrier"><path d="M698.8 563.1l69.2-.1V460.8l-69.1-.2-7.7.2c-14.4-.1-18.4-12-18.4-26.5 0-7.2 2.9-13.7 7.6-18.4l48.8-48.8-72.4-72.4-48.9 48.9c-4.7 4.7-11.2 7.6-18.4 7.6-14.5 0-26.2-11.7-26.3-26.1v-69.2H460.8v68.9h-.1c-.1 14.5-11.8 26.2-26.3 26.2-7.1 0-13.5-2.9-18.2-7.4l-49.1-49.1-72.3 72.4 48.8 48.8s0 .1-.1.1c4.7 4.8 7.6 11.3 7.6 18.5 0 14.4-4 26.2-18.4 26.5H256v102.4l69-.2v.2c14.4.1 26.1 11.8 26.1 26.3 0 7.1-2.9 13.5-7.4 18.2l-49 49 72.4 72.4 48.8-48.8s.1 0 .1.1c4.7-4.6 11.2-7.6 18.4-7.6 14.4 0 26.2 4.1 26.4 18.5 0 0-.1 7.5 0 7.5v69.3l102.4-.2v-69.1h.1c.2-14.4 11.8-26 26.2-26 7.2 0 13.6 2.9 18.4 7.5h.1l48.8 48.8 72.4-72.4-48.8-48.8c-4.6-4.7-7.5-11.2-7.5-18.4-.1-14.6 11.5-26.3 25.9-26.4zM512 614c-56.5 0-102.3-45.8-102.3-102.3S455.5 409.3 512 409.3s102.3 45.8 102.3 102.4C614.4 568.2 568.5 614 512 614z"></path></g></svg>
						<span class="flex-1 whitespace-nowrap">Settings</span>
					</a>
				</li>
				<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />

				{#each arr as name, index}
					<li>
						<Sidebar {name} />
					</li>
				{/each}

				<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />
				<li>
					<!-- weird shit target="_self" -->
					<a
						href="/logout"
						target="_self"
						class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
					>
						<svg
							aria-hidden="true"
							class="flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white"
							fill="currentColor"
							viewBox="0 0 20 20"
							xmlns="http://www.w3.org/2000/svg"
							><path
								fill-rule="evenodd"
								d="M3 3a1 1 0 00-1 1v12a1 1 0 102 0V4a1 1 0 00-1-1zm10.293 9.293a1 1 0 001.414 1.414l3-3a1 1 0 000-1.414l-3-3a1 1 0 10-1.414 1.414L14.586 9H7a1 1 0 100 2h7.586l-1.293 1.293z"
								clip-rule="evenodd"
							/></svg
						>
						<span class="flex-1 ml-3 whitespace-nowrap">Logout</span>
					</a>
				</li>
			</ul>
		</div>
	</aside>

	<div id="main" class="p-4 {blur} inset-0">
		<div
			class="p-4 border-2 bg-white border-gray-500 border-dashed rounded-lg dark:border-gray-700"
		>
			<!-- <div class="editor min-h-screen" use:quill={options} on:text-change={e => content = e.detail} /> -->
		</div>
	</div>


<style>
</style>
