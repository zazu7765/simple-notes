<script lang="ts">
	import NavBar from '../../../lib/NavBar.svelte';
	import Notebook from '../../../lib/notes.svelte';
	export let data: any;
	let login = data.bool;

	let num = 5;
	console.log(data);
	let content = data['response']['Data'];
	let arr: String[] = [];
	for (const item in content) {
		arr.push(content[item]['Title']);
	}
	console.log(arr);
</script>

<div>
	<NavBar {login} />
</div>
<div class="grid lg:grid-cols-3 md:grid-cols-2 xl:grid-cols-4 gap-4 grid-flow-row p-5">
	{#each arr as content, index}
		<Notebook {content} />
	{/each}
</div>
