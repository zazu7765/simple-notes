<script lang="ts">
	import { storage, write, currentNote } from './notes';
	import { writable } from 'svelte/store';
	import { redirect } from '@sveltejs/kit';
	export let ids: string;
	export let token: string;
	let notes;


	write.subscribe((value) => {
		notes = value;
	});

	let valueText: string = 'New note!';
	let arr = writable(['0']);
	let arrayNote: string[] = [];

	for (const note in notes['Data']) {
		if (notes['Data'][note]['NotebookID'] == ids) {
			arrayNote.push(notes['Data'][note]);
		}
	}
	$arr = arrayNote;

	$: addNote = async () => {
		const formData = new FormData();
		formData.append('id', ids);
		formData.append('Title', valueText);
		const response = await fetch('http://localhost:81/notes/create', {
			method: 'POST',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + token
			}
		});
		const responseNote = await fetch('http://localhost:81/notes/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + token,
				'Content-Type': 'application/json'
			}
		});
		notes = await responseNote.json();

		arrayNote = [];

		for (const note in notes['Data']) {
			if (notes['Data'][note]['NotebookID'] == ids) {
				arrayNote.push(notes['Data'][note]);
			}
		}
		$arr = arrayNote;
		$write = notes;
		let save = await response.json();
		console.log(save['Status'] + save['Message'] + 'mmmmmmm');
	};
	let click = false;
	let state = 'pointer-events-none blur-md';
	let zindex = '0';
	function stateChanged() {
		if (click == false) {
			click = true;
			state = '';
			zindex = '40';
		} else {
			click = false;
			state = 'pointer-events-none blur-md';
			zindex = '0';
		}
	}

	$: deleteNote = async (noteId: string) => {
		const formData = new FormData();
		formData.append('id', noteId);
		const deleteNote = await fetch('http://localhost:81/notes/', {
			method: 'DELETE',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + token
			}
		});
		let data = await deleteNote.json();
		if (data['Status'] == 'error') {
			throw redirect(302, '/logout');
		}
		const responseNote = await fetch('http://localhost:81/notes/all', {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + token,
				'Content-Type': 'application/json'
			}
		});
		notes = await responseNote.json();

		arrayNote = [];

		for (const note in notes['Data']) {
			if (notes['Data'][note]['NotebookID'] == ids) {
				arrayNote.push(notes['Data'][note]);
			}
		}
		$arr = arrayNote;
		$write = notes;
	};
</script>

<div class="p-2 justify-center mt-5 group grid grid-cols-4">
	<div
		class="col-span-3  mx-auto  max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
		<div class="">
			<input
				bind:value={valueText}
				class="relative block w-full appearance-none rounded  border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
				placeholder="Add a new note!"
			/>
		</div>
	</div>
	<!-- <div>
		<button
			on:click={() => {
				stateChanged();
			}}
			><div
				class=" bg-green-500 z-10  grid mt-7  max-w-sm p-2 w-10 text-white hover:bg-green-600 rounded-lg shadow  dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
			>
				+
			</div></button
		>
		</div> -->
	<div>
		<button
			class="place-self-start mt-7 z-10 bg-blue-500  px-6 py-3 text-white group-hover:inline-flex  max-w-sm p-6 border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
			on:click={() => {
				addNote();
				stateChanged();
			}}
			on:keydown={() => {
				addNote();
				stateChanged();
			}}>+</button
		>
	</div>
</div>
{#key $arr}
	{#each $arr.reverse() as title}
		<div class="p-2 mt-5 groupg grid grid-cols-4">
			<div class="col-span-3">
				<a
					on:click={() => {
						$currentNote = [title['ID'], token];
					}}
					href="notebook-{ids}/note/{title['ID']}"
					class=" mx-auto block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
				>
					<!-- TODO: REMOVE NOTE FUNCTION-->

					<h5 class="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white">
						{title['Title']}
					</h5>
					<p class="font-normal text-gray-700 dark:text-gray-400">
						{title['Content'].slice(0,25)}...
					</p>
				</a>
			</div>

			<div
				on:click={() => {
					deleteNote(title['ID']);
				}}
				on:keyup={() => {
					deleteNote(title['ID']);
				}}
				class="place-self-start float-right top-0 right-5  mt-7 z-10 bg-red-500 mx-2 px-6 py-3 text-white group-hover:inline-flex  max-w-sm p-6 border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
			>
				<button
					on:click={() => {
						deleteNote(title['ID']);
					}}
					class="justify-center ">X</button
				>
			</div>
		</div>
	{/each}
{/key}
