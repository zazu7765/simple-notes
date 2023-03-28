<script lang="ts">
    	import { toast } from '@zerodevx/svelte-toast';
	import { storage, write, currentNote, state,arr } from './notes';
	import { writable } from 'svelte/store';
	import { redirect } from '@sveltejs/kit';
	import { fade, fly } from 'svelte/transition';

	import { onMount } from 'svelte';

    let deleComponent;
	
	let deleInChild;
	
	onMount(function(){
		deleInChild = function(){
			deleComponent.deleteNote()
		}
	})
    export let title: any;
    let notes;
    let ids: string;
    export let token: string;
	state.subscribe((value) => {
		ids = value;
	});

	write.subscribe((value) => {
		notes = value;
	});
    let hover = '';
	let del = '';
    async function dele(noteId: string){
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

		let arrayNote = [];

		for (const note in notes['Data']) {
			if (notes['Data'][note]['NotebookID'] == ids) {
				arrayNote.push(notes['Data'][note]);
			}
		}
		$arr = arrayNote;
		$write = notes;
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
                    toast.push('Note deleted!', {
                        theme: {
                            '--toastColor': 'mintcream',
                            '--toastBarBackground': '#C70039'
                        }
                    });
                    // dele(title['ID']);
                    dele(title['ID']);
                }}
                class=" group-hover:flex text-[#ECF2FF] ">Delete</button
            >
        </div>
    </button>
{/if}

<!-- TODO: REMOVE NOTE FUNCTION-->
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

        {@html title['Content'].slice(0, 25)}
    </p>
</a>
</div>