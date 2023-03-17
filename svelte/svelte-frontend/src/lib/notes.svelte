<script lang="ts">
	import Del from '../routes/user/[slug]/+page.svelte'
	export let content: string;
	import { storage, write, state } from './notes';
	export let id;
	export let pointer;
	let notes;

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

	class="stack group mx-auto relative min-w-full h-full z-0 {pointer} max-w-sm py-6  border  rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 "
>
	<button
		on:animationend={() => {
			hovering();
		}}
		on:click={() => {
			
			hover = 'aaaaaa';
			console.log('asssss2');
		}}
		class="cursor-pointer absolute z-10 flex top-0 h-full duration-150 pointer-events-none hover:pointer-events-auto  m-auto right-0 bg-red-400  {hover} lg:w-0 rounded-md text-white group-hover:h-[100%] group-hover:w-[33%] "
	>
		<div class="m-auto">
			<button class=" group-hover:flex ">{del}</button>
		</div>
	</button>
	<button
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

		<h5 class="mx-auto text-2xl font-bold tracking-tight text-gray-900 {pointer} dark:text-white">
			{content}
		</h5>
		<p class="mx-auto font-normal text-gray-700 dark:text-gray-400 {pointer}">asdfasdf</p>
	</button>
</div>
