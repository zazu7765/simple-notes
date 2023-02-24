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
	let valueText = "New notebook!";

</script>

<div class="mt-5 pointer-events-none">
	<a
		href="notebook-{id}"
		class=" mx-auto block max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
	>
	{#if id == "0"}
	<div class="pointer-events-auto">
		<input
			bind:value={valueText}
			class="relative block w-full appearance-none rounded  border border-gray-300 px-3 py-2 text-gray-900 placeholder-gray-500 focus:z-10 focus:border-indigo-500 focus:outline-none focus:ring-indigo-500 sm:text-sm"
			placeholder="Add a new note!"
		/>
	</div>
	{/if}
		<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{content}</h5>
		<p class="font-normal text-gray-700 dark:text-gray-400 pointer-events-auto">
			Here are the biggest enterprise technology acquisitions of 2021 so far, in reverse
			chronological order.
		</p>
	</a>
</div>
