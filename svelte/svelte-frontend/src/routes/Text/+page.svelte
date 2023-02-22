<script>
	import { onDestroy } from 'svelte';

	import { quill } from 'svelte-quill';

	let options = { placeholder: 'Write something from outside...' };

	let content = { html: '', text: '' };
	onDestroy(() => {});
</script>

<svelte:head>
	<link href="//cdn.quilljs.com/1.3.6/quill.snow.css" rel="stylesheet" />
</svelte:head>
<body>
	<div class="grid lg:grid-cols-5 gap-4 min-w-max p-5 max-h-5 min-h-screen">
		<div />
		<div
			class="min-h-none col-span-3 grid-cols-3 gap-4 py-12 px-4 content-center inline-block sm:px-6 lg:px-8 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
		>
			<div class="col-span-3 p-4 rounded-lg">
				<!-- <h1>SVELTE-QUILL</h1> -->

				<div
					class="editor min-h-screen "
					use:quill={options}
					on:text-change={(e) => (content = e.detail)}
				/>

				<!-- <br /> -->
				<!-- Resulting HTML:
		
	{@html content.html} -->
			</div>
		</div>
	</div>
</body>

<style>

</style>
