<script module lang="ts">
	import { redirect } from '@sveltejs/kit';
	import Sidebar from '../../../lib/Sidebar.svelte';
	import { write } from '../../../lib/notes';
	
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
	let ww = "w-64";
	function openNav() {
		ww = "w-64"
	}
	function closeNav() {
	  ww = "w-0";
	}
</script>


	
<button on:click={()=>{
	openNav();
}} type="button" class=" z-50 inline-flex items-center p-2 mt-2 ml-3 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
	<span class="sr-only">Open sidebar</span>
	<svg class="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
	   <path clip-rule="evenodd" fill-rule="evenodd" d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"></path>
	</svg>
 </button>
 
 <aside class=" fixed top-0 left-0 z-40 {ww} h-screen sm:translate-x-0" aria-label="Sidebar">
	<div class="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
	   <ul class="space-y-2">
		
			<button  class="flex items-center pl-2.5 mb-5" on:click={()=>{closeNav()}}>
				<img
					src="https://flowbite.com/docs/images/logo.svg"
					class="h-6 mr-3 sm:h-7"
					alt="Flowbite Logo"
				/>
				<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
					>Flowbite</span
				>
			</button>
			<ul class="space-y-2">
				<li>
					<a
						href="#"
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
								d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
								clip-rule="evenodd"
							/></svg
						>
						<span class="flex-1 ml-3 whitespace-nowrap">Users</span>
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

	<div id="main" class="p-4 sm:ml-64 inset-0">
		<div
			class="p-4 border-2 bg-white border-gray-500 border-dashed rounded-lg dark:border-gray-700"
		>
			<!-- <div class="editor min-h-screen" use:quill={options} on:text-change={e => content = e.detail} /> -->
		</div>
	</div>


<style>
</style>
