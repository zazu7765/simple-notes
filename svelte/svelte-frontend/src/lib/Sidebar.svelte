<script script lang="ts">

	import { write } from './notes';
	import ElementD from './DropDownElement.svelte';
	let notes;

	write.subscribe((value) => {
		notes = value;
	});

	export let name: string[];
	let arrayNote: string[] = [];
	for (const note in notes['Data']) {

		if (notes['Data'][note]['NotebookID'] == name[0]) {
			arrayNote.push(notes['Data'][note]);

		}
	}

    let menu = 'hidden';

</script>

<li>
	<button
		type="button"
        on:click={() => {
            if (menu === 'hidden') {
                menu = 'visible';
            } else {
                menu = 'hidden';
            }
        }}
		class="flex items-center w-full p-2 text-base font-normal text-gray-900  duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700"


	>
		<span class="flex-1 ml-3 text-left whitespace-nowrap" >{name[1]}</span>
		<svg

			class="w-6 h-6"
			fill="currentColor"
			viewBox="0 0 20 20"
			xmlns="http://www.w3.org/2000/svg"
			><path
				fill-rule="evenodd"
				d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
				clip-rule="evenodd"
			/></svg
		>
	</button>
	<ul class="{menu} py-2 space-y-2">
		{#if arrayNote}
			{#each arrayNote as titles}
				<ElementD title={titles['Title']} id={titles['ID']} />
			{/each}
		{/if}
	</ul>
</li>
