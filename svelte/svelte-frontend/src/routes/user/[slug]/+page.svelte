<script module lang="ts">
	import { redirect } from '@sveltejs/kit';
	import Sidebar from '../../../lib/Sidebar.svelte';
	import { write, storage, state } from '../../../lib/notes';
	import { browser } from '$app/environment';
	import { onDestroy, onMount } from 'svelte';
	import NoteBook from '../../../lib/notes.svelte';
	import AllNotes from '../../../lib/allNotes.svelte';


	import { writable, type Writable } from 'svelte/store';

	let options = { placeholder: 'Write something from outside...' };
	let savestore = false;
	let contentEdit = { html: '', text: '' };
	let ses;
	let id;
	export let data: any;
	let pp;
	$write = data['responseNote'];
	import { navigating } from '$app/stores'
// Example spinner/loading component is visible (when $navigating != null):


	if ($storage && savestore) {
		window.sessionStorage.setItem('store', $storage);
		window.sessionStorage.setItem('id', $state);
	}
	onMount(async () => {
		ses = window.sessionStorage.getItem('store');
		id = window.sessionStorage.getItem('id');
		savestore = true;
	});

	write.subscribe((value) => {
		pp = value;
	});

	let content = data['response']['Data'];
	let arrS = writable(['']);
	let arr: string[] = [];
	for (const item in content) {
		arr.push([content[item]['ID'], content[item]['Title']]);
	}
	$arrS = arr;

	let ww = 'w-0';
	let blur = '';
	let pointer = '';
	function openNav() {
		ww = 'w-64';
		blur = 'blur-sm';
		pointer = 'pointer-events-none';
	}
	function closeNav() {
		ww = 'w-0';
		blur = '';
		pointer = '';
	}

	if (browser) {
		innerWidth = window.innerWidth;
		if (innerWidth < 375) {
			openNav();
		}
	}

	$: addNoteBook = async () => {
		const formData = new FormData();
		formData.append('id', arr[0]['UserID']);
		formData.append('Title', valueText);
		const response = await fetch('http://localhost:81/notebooks', {
			method: 'POST',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + data.token
			}
		});
		const responseAll = await fetch('http://localhost:81/notebooks/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + data.token,
				'Content-Type': 'application/json'
			}
		});
		let content2 = await responseAll.json();
		let lengthA = content2['Data'].length - 1;
		console.log(content2['Data'][lengthA]['ID']);

		arr.push([content2['Data'][lengthA]['ID'], content2['Data'][lengthA]['Title']]);

		$arrS = arr;
	};
	let valueText = 'New Notebook!';
</script>

<!-- <svelte:window bind:innerWidth={screenWidth} />
{#if screenWidth < 600}
	
{:else}
 
{/if} -->
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

<aside class="ease-in-out duration-300 fixed top-0 left-0 z-40 {ww} h-screen " aria-label="Sidebar">
	<div class="h-full py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
		<ul class="space-y-2">
			<button
				class="flex items-center pl-2.5 mb-5"
				on:click={() => {
					closeNav();
				}}
			>
				<img
					src="https://flowbite.com/docs/images/logo.svg"
					class="h-6 mr-3 sm:h-7"
					alt="Flowbite Logo"
				/>
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
						<svg
							class="h-10 w-10"
							fill="#000000"
							viewBox="0 0 1024 1024"
							xmlns="http://www.w3.org/2000/svg"
							><g id="SVGRepo_bgCarrier" stroke-width="0" /><g
								id="SVGRepo_tracerCarrier"
								stroke-linecap="round"
								stroke-linejoin="round"
							/><g id="SVGRepo_iconCarrier"
								><path
									d="M698.8 563.1l69.2-.1V460.8l-69.1-.2-7.7.2c-14.4-.1-18.4-12-18.4-26.5 0-7.2 2.9-13.7 7.6-18.4l48.8-48.8-72.4-72.4-48.9 48.9c-4.7 4.7-11.2 7.6-18.4 7.6-14.5 0-26.2-11.7-26.3-26.1v-69.2H460.8v68.9h-.1c-.1 14.5-11.8 26.2-26.3 26.2-7.1 0-13.5-2.9-18.2-7.4l-49.1-49.1-72.3 72.4 48.8 48.8s0 .1-.1.1c4.7 4.8 7.6 11.3 7.6 18.5 0 14.4-4 26.2-18.4 26.5H256v102.4l69-.2v.2c14.4.1 26.1 11.8 26.1 26.3 0 7.1-2.9 13.5-7.4 18.2l-49 49 72.4 72.4 48.8-48.8s.1 0 .1.1c4.7-4.6 11.2-7.6 18.4-7.6 14.4 0 26.2 4.1 26.4 18.5 0 0-.1 7.5 0 7.5v69.3l102.4-.2v-69.1h.1c.2-14.4 11.8-26 26.2-26 7.2 0 13.6 2.9 18.4 7.5h.1l48.8 48.8 72.4-72.4-48.8-48.8c-4.6-4.7-7.5-11.2-7.5-18.4-.1-14.6 11.5-26.3 25.9-26.4zM512 614c-56.5 0-102.3-45.8-102.3-102.3S455.5 409.3 512 409.3s102.3 45.8 102.3 102.4C614.4 568.2 568.5 614 512 614z"
								/></g
							></svg
						>
						<span class="flex-1 whitespace-nowrap">Settings</span>
					</a>
				</li>
				<ul class="pt-4 mt-4 space-y-2 border-t border-gray-500 dark:border-gray-700" />

				{#each $arrS.reverse() as name}
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
		</ul>
	</div>
</aside>
{#if !$navigating}
	

<div
	class="min-h-screen p-4 {blur} inset-0"
	on:keydown={() => {
		closeNav();
	}}
	on:click={() => {
		closeNav();
	}}
>
	<div class={pointer}>
		{#if ses == 'notebook'}
			<div class=" grid lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-1 gap-1 ">
				<div class="mt-5 ">
					<div
						href="notebook-0"
						class=" mx-auto block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
					>
						<div class="pointer-events-auto">
							<input
								bind:value={valueText}
								class="relative block w-full appearance-none rounded  border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
								placeholder="Add a new note!"
							/>
						</div>

						<button
					on:click={() => {
						addNoteBook();
					}}
					class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white"
				>
					+
				</button>

						<p class="font-normal text-gray-700 dark:text-gray-400 pointer-events-auto">
							Here are the biggest enterprise technology acquisitions of 2021 so far, in reverse
							chronological order.
						</p>
					</div>
					<div>
				
					</div>
				</div>

				{#each $arrS.reverse() as name}
					<button
						on:click={() => {
							$state = name[0];

							$storage = 'notes';
							$: if ($storage) {
								window.sessionStorage.setItem('store', $storage);
							}
							ses = $storage;

							$: if ($state) {
								window.sessionStorage.setItem('id', $state);
							}
							id = $state;
						}}
					>
						<NoteBook id={name[0]} content={name[1]} />
					</button>
				{/each}
			</div>
		{/if}

		{#if ses == 'notes'}
			<div class=" grid lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-1 gap-1 ">
				<button
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
					}}>go back</button
				>
				{#key $write}
					<AllNotes token={data.token} ids={id} />
				{/key}
			</div>
		{/if}
	</div>
</div>
{:else}
<div>wating</div>
{/if}
<style>
</style>
