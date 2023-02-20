<script script lang="ts">
    import { write } from './notes'
    import ElementD from './DropDownElement.svelte'
    let notes;

	write.subscribe(value => {
		notes = value;
	});
    console.log(notes['Data'][0]['NotebookID']+"dafasdfasdfasdfasd")
    export let name: string[];
    let arrayNote: string[] = [];
    for (const note in notes['Data']){
        console.log(notes["Data"][note]['Title'])
        if (notes["Data"][note]['NotebookID'] == name[0]){
            arrayNote.push(notes["Data"][note])
            console.log("afjklns;hhjlfdshjkldfhjksaglhjklfdjhkldfshjkldfsa")
        }
    }
    console.log("hhjlfdshjkldfhjksaglhjklfdjhkldfshjkldfsaa" + arrayNote)
</script>

<li>
    <button type="button" class="flex items-center w-full p-2 text-base font-normal text-gray-900 transition duration-75 rounded-lg group hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700" aria-controls="dropdown-example" data-collapse-toggle={name[1]}>
          
          <span class="flex-1 ml-3 text-left whitespace-nowrap" sidebar-toggle-item>{name[1]}</span>
          <svg sidebar-toggle-item class="w-6 h-6" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z" clip-rule="evenodd"></path></svg>
    </button>
    <ul id={name[1]} class="hidden py-2 space-y-2">
        {#if arrayNote}
        {#each arrayNote as titles, index}
        <ElementD title = {titles['Title']} id = {titles['ID']}/>
        fsadfdsa
        {/each}
        {/if}
        
    </ul>
 </li>
 
