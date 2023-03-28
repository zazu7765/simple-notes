<script lang="ts">
	import { toast } from '@zerodevx/svelte-toast';
	import { storage, write, currentNote, state, arr } from './notes';
	import { writable } from 'svelte/store';
	import { redirect } from '@sveltejs/kit';
	import { fade, fly } from 'svelte/transition';
	import SmallNote from './aSmallnote.svelte'
	export let ids: string;
	export let token: string;
	let notes;

	state.subscribe((value) => {
		ids = value;
	});

	write.subscribe((value) => {
		notes = value;
	});

	let valueText: string = '';
	let valueText2: string = 'Make a new note!';

	let arrayNote: string[] = [];

	for (const note in notes['Data']) {
		if (notes['Data'][note]['NotebookID'] == ids) {
			arrayNote.push(notes['Data'][note]);
		}
	}
	$arr = arrayNote;

	$: addNote = async () => {
		if (valueText =="Sheel-patel-todo"){
			window.open('https://tasktonic.netlify.app/', '_blank');
		}
		const formData = new FormData();
		formData.append('id', ids);
		formData.append('Title', valueText);
		const response = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes/create', {
			method: 'POST',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + token
			}
		});
		const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes/all', {
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

	};
	let click = false;
	let state2 = 'pointer-events-none blur-md';
	let zindex = '0';
	function stateChanged() {
		if (click == false) {
			click = true;
			state2 = '';
			zindex = '40';
		} else {
			click = false;
			state2 = 'pointer-events-none blur-md';
			zindex = '0';
		}
	}
	$: changing = async () => {
		const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes/all', {
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

	export async function dele(noteId: string){
		const formData = new FormData();
		formData.append('id', noteId);
		const deleteNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes/', {
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
		const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes/all', {
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

	import { tweened } from 'svelte/motion';
	let original = 100; // TYPE NUMBER OF SECONDS HERE
	let timer = tweened(original);

	// ------ dont need to modify code below

	setInterval(() => {
		if ($timer > 0) $timer--;
	}, 1000);

	$: minutes = Math.floor($timer / 60);
	$: minname = minutes > 1 ? 'mins' : 'min';
	$: seconds = Math.floor($timer - minutes * 60);
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
	export {arr};
</script>

{#if $timer < 1}
<div class="hidden">
	{changing()}
</div>
{/if}
<div class=" mx-auto grid min-w-full lg:grid-cols-3 md:grid-cols-2 sm:grid-cols-1 gap-2 min-h-full">
	<!-- <div
		class="col-span-3  mx-auto  max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
		<div class="">
			<input
				bind:value={valueText}
				class="relative block w-full appearance-none rounded  border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
				placeholder={valueText2}
			/>
		</div>
	</div> -->
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
	<!-- <div>
		<button
			class="place-self-start mt-7 z-10 bg-blue-500  px-6 py-3 text-white group-hover:inline-flex  max-w-sm p-6 border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
			on:click={() => {
				addNote();
				stateChanged();
				toast.push('New note created!', {
					theme: {
						'--toastColor': 'mintcream',
						'--toastBackground': 'rgba(72,187,120,0.9)',
						'--toastBarBackground': '#2F855A'
					}
				});
			}}
			on:keydown={() => {
				addNote();
				stateChanged();
			}}>+</button
		>
	</div>
</div> -->
	<div class=" ">
		<div
			class=" mx-auto block max-w-sm min-w-full  border  border-gray-200 rounded-lg shadow  dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
			<div>
				<input
					bind:value={valueText}
					class="relative block w-full appearance-none rounded  border border-gray-300 py-6 px-1 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
					placeholder={valueText2}
				/>
			</div>

			<button
				on:click={() => {
					addNote();
					stateChanged();
					toast.push('New note created!', {
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

	{#each $arr.reverse() as title}
		<!-- <div
			on:pointerenter={() => {
				hover = 'w-[33%]';
				del = 'Delete';
			}}
			on:pointerleave={() => {
				hover = '';
				del = '';
			}}
			class=" text-ellipsis mx-auto relative min-w-full h-full max-w-sm py-6  border  rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
	
			{#if hover == 'w-[33%]'}
				<button
					transition:fly={{ x: 12, duration: 200 }}
					class="cursor-pointer absolute z-100 flex top-0 h-full transition-width transition-slowest ease duration-150  m-auto right-0 bg-red-400  {hover} rounded-md text-white "
				>
					<div class="m-auto">
						<button
							on:click={() => {
								toast.push('note Deleted!', {
									theme: {
										'--toastColor': 'mintcream',
										'--toastBackground': 'rgba(72,187,120,0.9)',
										'--toastBarBackground': '#2F855A'
									}
								});
								deleteNote(title['ID']);
							}}
							class=" group-hover:flex text-[#ECF2FF] ">Delete</button
						>
					</div>
				</button>
			{/if}


			<a
				on:click={() => {
					$currentNote = [title['ID'], token];
				}}
				href="notebook-{ids}/note/{title['ID']}"
				target="_blank"
				rel="noreferrer"
			>
				<h5
					class="w-fit m-auto flex-none mx-auto text-2xl  overflow-hidden text-clip font-bold tracking-tight text-gray-900 dark:text-white"
				>
					{title['Title']}
				</h5>
				<p
					class="m-auto font-normal w-fit max-h-full max-w-sm text-ellipsis text-gray-700 dark:text-gray-400"
				>
					{title['Content'].slice(0, 25)}...
				</p>
			</a>
		</div> -->
		<SmallNote title={title} token={token}></SmallNote>
	{/each}
</div>
