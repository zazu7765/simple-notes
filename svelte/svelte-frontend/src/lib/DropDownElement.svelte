<script lang="ts">

    export let title: string;
    export let id: any;


</script>


<li>
    <a href="notebook/note/{id}" target="_blank" rel="noreferrer" class="flex items-center w-full p-2 text-base font-normal text-gray-900 transition duration-75 rounded-lg pl-11 group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700">{title.slice(0,20)}</a>
</li>