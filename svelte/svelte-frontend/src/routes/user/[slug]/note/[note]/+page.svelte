<script lang="ts">
	import { onDestroy, onMount } from 'svelte';
	import { page } from '$app/stores';

	import { quill } from 'svelte-quill';
	import { currentNote, storage } from '$lib/notes';
	import { redirect } from '@sveltejs/kit';
	import Editor from '@tinymce/tinymce-svelte';

let value2 = 'some content';

	let editor;
	export let toolbarOptions = [
		[{ header: 1 }, { header: 2 }, 'blockquote', 'link', 'image', 'video'],
		['bold', 'italic', 'underline', 'strike'],
		[{ list: 'ordered' }, { list: 'ordered' }],
		[{ align: [] }],
		['clean']
	];

	let options = { placeholder: 'Write something from outside...' };

	let content;
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

	let dataNote = data.notes;
	let quillE: any;
	onMount(async () => {
		console.log($page.params.note);
		console.log(dataNote['Data']);
		
		alert((dataNote['Data']['Content']));
		const { default: Quill } = await import('quill');

		quillE = new Quill(editor, {
			modules: {
				toolbar: toolbarOptions
			},
			theme: 'snow',
			placeholder: 'Write your story...'
		});

			// quillE.setContents(dataNote['Data']['Content']);
			value2 = dataNote['Data']['Content']
		
		// quillE.insertEmbed(dataNote['Data']['Content']);

	});

	// onDestroy(async () => {

	//     const formData = new FormData();
	// 	formData.append('id', dataNote["Data"]["ID"]);
	// 	formData.append('Content', content);
	// 	const response = await fetch('http://localhost:81/notes', {
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
		console.log(value2);
		var delta = value2
		console.log(editor);
		const formData = new FormData();
		formData.append('id', dataNote['Data']['ID']);
		formData.append('Content', delta);
		const response = await fetch('http://localhost:81/notes', {
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
			console.log('No user input detected.');
			doneTyping();
		}, 5000);
		// Set the timeout to 5 seconds (5000 milliseconds)
	}
	let content2 = { html: '', text: '' };
</script>

<div class="p-4 ">
	<div class="p-4 border-2 border-gray-200 border-dashed rounded-lg dark:border-gray-700">
		<div class="editor-wrapper editor h-screen">
			<div
				bind:this={editor}
				on:keydown={() => {
					clearTimeout(timer);
					startTimer();
				}}
				bind:textContent={content}
				contenteditable
			/>
		</div>
	</div>
</div>


<main>
	<h1>{value2}</h1>
	<Editor bind:value={value2}
	  apiKey="ro357dg4ght7s6z2r3m0f2uslrfl9i9dfogk5mvn3dj2xhr5"
	/>
  </main>
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
