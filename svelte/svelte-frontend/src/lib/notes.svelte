<script lang="ts">
	export let content: string;
	import { storage, write } from './notes';
	export let id;
	let notes;

	write.subscribe((value) => {
		notes = value;
	});
	let arrayNote: string[] = [];
	for (const note in notes['Data']) {
		if (notes['Data'][note]['NotebookID'] == id) {
			arrayNote.push(notes['Data'][note]);
		}
	}

</script>

<div class="mt-5">
	<a
		href="notebook-{id}"
		class=" mx-auto block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
		<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{content}</h5>
		<p class="font-normal text-gray-700 dark:text-gray-400">
			Here are the biggest enterprise technology acquisitions of 2021 so far, in reverse
			chronological order.
		</p>
	</a>
</div>
