<script>
   import NoteBook from './notes.svelte'
   import {arrS, arr} from '$lib/notes'
	import { writable } from 'svelte/store';
   import Sidebar from './Sidebar.svelte'
   export let token;
   export let pointer;
   export let mode;
   let mouse = writable([])
   arrS.subscribe((value)=>{
      $mouse = value
   })
</script>
{#if mode==0}
{#each $mouse as name}
<div 
   class=" mx-auto text-ellipsis min-h-full block w-[100%]  border border-slate-700 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700"
>
   <NoteBook token={token} id={name[0]} content={name[1]} pointer={pointer} pre={name[2]} />
</div>
{/each}
{:else if  mode==1}
{#each $arrS.reverse() as name}
<li>
   <Sidebar {name} />
</li>
{/each}
{/if}