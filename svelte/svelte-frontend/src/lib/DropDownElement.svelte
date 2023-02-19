<script lang="ts">
	import { writable } from "svelte/store";
    import { write } from "../lib/notes"

    export let title: string;
    let notes;
    write.subscribe((arg)=>notes)

</script>


<li>
    <a href={title} class="flex items-center w-full p-2 text-base font-normal text-gray-900 transition duration-75 rounded-lg pl-11 group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700">{title}</a>
 </li>