<script module lang="ts">
	var url = import.meta.env.VITE_PUBLIC_URL
	import { toast } from '@zerodevx/svelte-toast';
	import { redirect } from '@sveltejs/kit';
	import Sidebar from '../../../lib/Sidebar.svelte';
	import { write, storage, state, arrS} from '../../../lib/notes';
	import { browser } from '$app/environment';
	import { onDestroy, onMount } from 'svelte';
	import NoteBook from '../../../lib/notes.svelte';
	import AllNotes from '../../../lib/allNotes.svelte';
	import { fade } from 'svelte/transition';
	import { writable, type Writable } from 'svelte/store';
	import NoteBooks from '$lib/Notebooks.svelte'

	let options = { placeholder: 'Write something from outside...' };
	let savestore = false;
	let contentEdit = { html: '', text: '' };
	let ses;
	let id;
	export let data: any;
	let pp;
	let valueText2: String = 'Make a new notebook!';
	$write = data['responseNote'];
	import { navigating } from '$app/stores';
	import Notebooks from '$lib/Notebooks.svelte';

	// Example spinner/loading component is visible (when $navigating != null):

	if ($storage && savestore) {
		window.sessionStorage.setItem('store', $storage);
		window.sessionStorage.setItem('id', $state);
	}
	let loaded = false;
	onMount(async () => {
		ses = window.sessionStorage.getItem('store');
		id = window.sessionStorage.getItem('id');
		savestore = true;
		loaded = true;
	});

	write.subscribe((value) => {
		pp = value;
	});

	let content = data['response']['Data'];
	
	let arr: string[] = [];
	for (const item in content) {
		arr.unshift([content[item]['ID'], content[item]['Title'], content[item]['Description']]);
	}
	arrS.set(arr)

	let ww = 'w-0';
	let blur = '';
	let pointer = '';
	let w = ''
	function openNav() {
		ww = 'w-64';
		blur = 'blur-sm';
		pointer = 'pointer-events-none';
		w='hidden'
	}
	function closeNav() {
		ww = 'w-0';
		w = ''
		blur = '';
		pointer = 'pointer-events-auto';
	}

	if (browser) {
		innerWidth = window.innerWidth;
		if (innerWidth < 375) {
			openNav();
		}
	}

	$: addNoteBook = async () => {
		const formData = new FormData();
		// formData.append('id', $id);
		formData.append('Title', valueText);
		const response = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks', {
			method: 'POST',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + data.token
			}
		});
		const responseAll = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + data.token,
				'Content-Type': 'application/json'
			}
		});
		let content2 = await responseAll.json();
		let lengthA = content2['Data'].length - 1;


		arr.unshift([content2['Data'][lengthA]['ID'], content2['Data'][lengthA]['Title']]);

		arrS.set(arr)
		$arrS = arr
	};
	export const deleteNote = async (noteId: string) => {
		const formData = new FormData();
		formData.append('id', noteId);
		const deleteNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks/', {
			method: 'DELETE',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + data.token
			}
		});
		let data1 = await deleteNote.json();
		if (data1['Status'] == 'error') {
			throw redirect(302, '/logout');
		}
		const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notebooks/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + data.token,
				'Content-Type': 'application/json'
			}
		});
		let content = await responseNote.json();

		let arr2 = [];

		for (const item in content) {
		arr2.unshift([content[item]['ID'], content[item]['Title']]);
	}
		arrS.set(arr)
		$arrS = arr
		
	};
	let valueText = '';
	arrS.subscribe((value)=>{
		arr = value
	})
</script>

<!-- <svelte:window bind:innerWidth={screenWidth} />
{#if screenWidth < 600}
	
{:else}
 
{/if} -->
{#if loaded==true}
<button
	on:click={() => {
		openNav();
	}}
	type="button"
	class="Class

	z-50 inline-flex  items-center p-6 mt-2 ml-3 text-sm text-gray-500 rounded-lg hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
>
	<span class="sr-only">Open sidebar</span>
	<svg
		class="w-6 h-6"
		aria-hidden="true"
		fill="currentColor"
		viewBox="0 0 20 20"
		xmlns="http://www.w3.org/2000/svg"
	>
		<path
			clip-rule="evenodd"
			fill-rule="evenodd"
			d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
		/>
	</svg>
</button>

<div class="ease-in-out duration-300 fixed top-0 left-0 z-40 {ww} h-screen " aria-label="Sidebar">
	<div class="h-full py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
		<ul class="space-y-2">
			<button
				class="flex items-center pl-2.5 mb-5"
				on:click={() => {
					closeNav();
				}}
			>

				<span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white"
					>Simple Notes
				</span>
			</button>
			<ul class="space-y-2">
				<li>
					<a
						href="#"
						class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700"
					>
				
						<span class="flex-1 whitespace-nowrap">Your notebooks:</span>
					</a>
				</li>
				<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />
<!-- 
				{#each $arrS.reverse() as name}
					<li>
						<Sidebar {name} />
					</li>
				{/each} -->
				<Notebooks mode=1 token={data.token} pointer={pointer}/>

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
		</ul>
	</div>
</div>
{#if data}
	<div
		class="min-h-screen p-4 {blur}  inset-0"
		on:keydown={() => {
			closeNav();
		}}
		on:click={() => {
			closeNav();
		}}
	>
		<div class="{pointer} ">
			{#if $storage == 'notebook'}
				<div class=" mx-auto grid min-w-full lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-1 gap-2">
					<div class=" ">
						<div

							class=" mx-auto block max-w-sm min-w-full  border  border-gray-200 rounded-lg shadow  dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
						>
							<div class={pointer}>
								<input
									bind:value={valueText}
									class="relative block w-full appearance-none rounded  border border-gray-300 py-6 px-1 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
									placeholder={valueText2}
								/>
							</div>

							<button
								on:click={() => {
									addNoteBook();
									toast.push('Notebook created!', {
                        theme: {
                            '--toastColor': 'mintcream',
                            '--toastBackground': 'rgba(72,187,120,0.9)',
                            '--toastBarBackground': '#2F855A'
                        }
                    });
								}}
								class="min-w-full hover:bg-blue-200 text-2xl text-center font-bold tracking-tight text-gray-900 dark:text-white"
							>
								<p>+</p>
							</button>
						</div>
						<div />
					</div>
					<Notebooks mode=0 token={data.token} pointer={pointer}/>
					<!-- {#each $arrS as name} -->
						<!-- <div 
							class=" mx-auto text-ellipsis min-h-full block w-[100%]  border border-slate-700 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
							<NoteBook token={data.token} id={name[0]} content={name[1]} pointer={pointer} pre={name[2]} />
						</div>
					{/each} -->
				</div>
			
			{/if}

			{#if $storage == 'notes'}

					<button
						class=" absolute top-5 right-5 z-0 {w} "
						on:click={() => {
							$state = '0';

							$storage = 'notebook';
							$: if ($storage) {
								window.sessionStorage.setItem('store', $storage);
							}
							ses = $storage;
							$: if ($state) {
								window.sessionStorage.setItem('id', $state);
							}
							id = $state;
						}}><p class="border-2 px-2 rounded-md py-2">Go back</p></button
					>
					
						<AllNotes token={data.token} ids={$state} />
					

			{/if}
		</div>
	</div>
{:else}
	<div>wating</div>
{/if}
{:else}
<div class="          grid
content-center 
mx-auto
max-w-2xl
rounded-lg
mt-16 items-center justify-center">
	<div
	  class="inline-block h-8 w-8 animate-spin rounded-full border-4 border-solid border-current border-r-transparent align-[-0.125em] motion-reduce:animate-[spin_1.5s_linear_infinite]"
	  role="status">
	  <span
		class="!absolute !-m-px !h-px !w-px !overflow-hidden !whitespace-nowrap !border-0 !p-0 ![clip:rect(0,0,0,0)]"
		>Loading...</span
	  >
	</div>
  </div>
{/if}

<style>
	:global(body){
		background-color: #ECF2FF;
	}
</style>