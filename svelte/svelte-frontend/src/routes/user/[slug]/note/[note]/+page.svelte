<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { page } from '$app/stores';


	import { currentNote, storage } from '$lib/notes';
	import { redirect } from '@sveltejs/kit';
	import Editor from '@tinymce/tinymce-svelte';
	import { HtmlTag } from 'svelte/internal';
	import { toast } from '@zerodevx/svelte-toast';

let value2 = 'Loading';


	let noteId: string[] = [];
	currentNote.subscribe((value) => {
		noteId = value;
	});
	// console.log(noteId[0]);
	// if (noteId[0] == '0') {
	// 	throw redirect(303, '/logout');
	// }
	$storage = 'notes';
	var text;

	// export let data: any;
	export let data: any;
	let load = false
	let dataNote = data.notes;

	onMount(async () => {
		
		
		// alert((dataNote['Data']['Content']));


	

			// quillE.setContents(dataNote['Data']['Content']);
			value2 = dataNote['Data']['Content']
			load = true
		// quillE.insertEmbed(dataNote['Data']['Content']);
		toast.push('Your note is loading...', {
                        theme: {
                            '--toastColor': 'mintcream',
                            '--toastBackground': '#99C1F0',
                            '--toastBarBackground': '#006AE8'
                        }
                    });

	});

	// onDestroy(async () => {

	//     const formData = new FormData();
	// 	formData.append('id', dataNote["Data"]["ID"]);
	// 	formData.append('Content', content);
	// 	const response = await fetch('/notes', {
	// 		method: 'PUT',
	// 		body: formData,
	// 		headers: {
	// 			Authorization: 'Bearer ' + data.token
	// 		}
	// 	});
	// 	console.log(content+" hamid");
	// })

	let typingTimer; //timer identifier
	let doneTypingInterval = 5000;
	$: doneTyping = async () => {

		var delta = value2

		const formData = new FormData();
		formData.append('id', dataNote['Data']['ID']);
		formData.append('Content', delta);
		const response = await fetch(import.meta.env.VITE_PUBLIC_URL_FRONTEND+'/notes', {
			method: 'PUT',
			body: formData,
			headers: {
				Authorization: 'Bearer ' + data.token
			}
		});
	};
	let timer;
	function startTimer() {
		timer = setTimeout(function () {

			doneTyping();
			toast.push('Note saved!', {
                        theme: {
                            '--toastColor': 'mintcream',
                            '--toastBackground': 'rgba(72,187,120,0.9)',
                            '--toastBarBackground': '#2F855A'
                        }
                    });
		}, 5000);
		// Set the timeout to 5 seconds (5000 milliseconds)
	}



</script>
{#if load==true}
<div class="p-4 ">
	<div class="p-4 border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700">

			<!-- <h1>{@html value2}</h1> -->
			<Editor bind:value={value2} 				on:keydown={() => {
				clearTimeout(timer);
				startTimer();
			}}
			  apiKey="ro357dg4ght7s6z2r3m0f2uslrfl9i9dfogk5mvn3dj2xhr5"
			/>
			<!-- <div
				bind:this={editor}
				on:keydown={() => {
					clearTimeout(timer);
					startTimer();
				}}
				bind:textContent={content}
				contenteditable
			/>-->

	</div>
</div>
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

<!-- <div
	class="editor"
	use:quill={options}
	on:keydown={() => {
		clearTimeout(timer);
		startTimer();
	}}
	on:text-change={(e) => (content2 = e.detail)}
/>
{content2.html} -->
