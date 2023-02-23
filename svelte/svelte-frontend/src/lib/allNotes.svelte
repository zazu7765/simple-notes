<script lang="ts">
	import { storage, write } from './notes';
	import { writable } from 'svelte/store';
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
		const response = await fetch('http://localhost:81/notes', {
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
	$: deleteNote = ()=>{
		
	}
</script>

<div class="p-2 flex mb-20 justify-center ">
	<button on:click={()=>{
		stateChanged();
	}}
		><div
			class="absolute bg-green-500 z-10  grid mt-7  max-w-sm p-2 w-10 text-white hover:bg-green-600 rounded-lg shadow  dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
			+
		</div></button
	>
	<div
		class=" {state} z-{zindex} absolute grid  grid-cols-2 gap-1 mx-auto  max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
		<div>
			<input
				bind:value={valueText}
				class="relative block w-full appearance-none rounded  border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
				placeholder="Add a new note!"
			/>
		</div>
		<div >
			<button
				class="bg-blue-500 rounded-md px-3 py-2 text-white"
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
</div>
{#key $arr}
	<div class=" grid lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-1 gap-1 ">
		{#each $arr.reverse() as title}
			<div class="p-2 mt-5 group">
				<a
					href="notebook-{ids}"
					class=" mx-auto block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
				>
					<!-- TODO: REMOVE NOTE FUNCTION-->
					<button
						class="hidden float-right top-0 right-0 bg-red-500 rounded-md mx-2 px-2 py-1 text-white group-hover:inline-flex"
						>X</button
					>
					<h5 class="mb-2 text-xl font-bold tracking-tight text-gray-900 dark:text-white">
						{title['Title']}
					</h5>
					<p class="font-normal text-gray-700 dark:text-gray-400">
						{title['Content']}
					</p>
				</a>
			</div>
		{/each}
	</div>
{/key}
