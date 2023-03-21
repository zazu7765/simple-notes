<script lang="ts">
	import Del from '../routes/user/[slug]/+page.svelte';
	export let content: string;
	import { storage, write, state } from './notes';
	export let id;
	export let pointer;
	let notes;
	import { fade,fly } from "svelte/transition"

	write.subscribe((value) => {
		notes = value;
	});
	let arrayNote: string[] = [];
	for (const note in notes['Data']) {
		if (notes['Data'][note]['NotebookID'] == id) {
			arrayNote.push(notes['Data'][note]);
		}
	}
	let valueText = '';
	let valueText2 = 'New notebook!';
	let hover = '';
	let del = '';
	function hovering() {
		if (hover == '') {
			hover = 'w-[33%]';
			del = 'Delete';
		} else {
			hover = '';
			del = '';
		}
	}
</script>

<div 
on:pointerenter={() => {
	hovering();
	
	console.log('asssss2');
}}

	class=" text-ellipsis  group mx-auto relative min-w-full h-full {pointer} max-w-sm py-6  border  rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 "
>{#if hover=="w-[33%]"}
	<button transition:fly="{{ x: 10, duration: 200 }}"	



	
		class="cursor-pointer absolute z-100 flex top-0 h-full transition-width transition-slowest ease duration-150  m-auto right-0 bg-red-400  {hover} rounded-md text-white "
	>
		<div class="m-auto">
			
			<button  class=" group-hover:flex text-[#ECF2FF] ">Delete</button>
			
		</div>
		
	</button>
	{/if}
	<button
		class="min-w-full inline-block text-ellipsis"
		
		on:click={() => {
			$state = id;

			$storage = 'notes';
			$: if ($storage) {
				window.sessionStorage.setItem('store', $storage);
			}

			$: if ($state) {
				window.sessionStorage.setItem('id', $state);
			}
		}}
	>
		{#if id == '0'}
			<div class="">
				<input bind:value={valueText} class=" block {pointer}" placeholder={valueText2} />
			</div>
		{/if}

		<p class="w-fit flex-none mx-auto text-2xl  overflow-hidden text-clip font-bold tracking-tight text-gray-900 {pointer} dark:text-white">
			{#if content.length<35}
			{content}
			{:else}
			{content.slice(0,34)}...
			{/if}
		</p>
		<p class="mx-auto font-normal max-h-full max-w-sm text-ellipsis text-gray-700 dark:text-gray-400 {pointer}">asdfasdf</p>
	</button>
</div>
