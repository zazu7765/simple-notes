<script lang="ts">
	import Del from '../routes/user/[slug]/+page.svelte';
	export let content: string;
	import { storage, write, state,arrS } from './notes';
	import { toast } from '@zerodevx/svelte-toast';
	export let id;
	export let token;
	export let pointer;
	export let pre;
	let notes;

	import { fade, fly } from 'svelte/transition';
	import { redirect } from '@sveltejs/kit';

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
	async function deleteNote(noteId: string){
		const formData = new FormData();
		formData.append('id', noteId);
		const deleteNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks/', {
			method: 'DELETE',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + token
			}
		});
		let data1 = await deleteNote.json();
		if (data1['Status'] == 'error') {
			throw redirect(302, '/logout');
		}
		const responseAll = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + token,
				'Content-Type': 'application/json'
			}
		});
		let content2 = await responseAll.json();

		let arr2 = [];

		for (const item in content2['Data']) {

		arr2.push([content2['Data'][item]['ID'], content2['Data'][item]['Title']]);
	}
		$arrS = arr2
		
	};

</script>

<div
	on:pointerenter={() => {
		hover = 'w-[33%]';
			del = 'Delete';
	}}
	on:pointerleave={() => {

		hover = '';
		del = '';
	}}
	class=" text-ellipsis  group mx-auto relative min-w-full h-full {pointer} max-w-sm py-6  border  rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700 "
>
	{#if hover == 'w-[33%]'}
		<button
			transition:fly={{ x: 12, duration: 400 }}
			class="cursor-pointer absolute z-100 flex top-0 h-full transition-width transition-slowest ease duration-150  m-auto right-0 bg-red-400  {hover} rounded-md text-white "
		>
			<div class="m-auto">
				<button on:click={()=>{
					deleteNote(id)
					toast.push('Notebook Deleted!', {
                        theme: {
                            '--toastColor': 'mintcream',

                            '--toastBarBackground': '#C70039'
                        }
                    });
				}} class=" group-hover:flex text-[#ECF2FF] ">Delete</button>
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

		<p
			class="w-fit flex-none mx-auto text-2xl  overflow-hidden text-clip font-bold tracking-tight text-gray-900 {pointer} dark:text-white"
		>
			{#if content.length < 35}
				{content}
			{:else}
				{content.slice(0, 34)}...
			{/if}
		</p>
	
	</button>
</div>
