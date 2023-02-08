<script lang="ts">
// 	import { onMount } from "svelte";
//   onMount(() => {
//     document.requestStorageAccess().then(
//   () => { console.log('access granted') },
//   () => { console.log('access denied') }
// );     

//    });

  import NavBar from "../lib/NavBar.svelte"
  
  export let data:any;
  let login = data.bool;

  import { onMount } from "svelte";
		import { Editor, rootCtx, defaultValueCtx, editorViewOptionsCtx } from "@milkdown/core";
		import { commonmark } from "@milkdown/preset-commonmark";
		import { nord } from "@milkdown/theme-nord";


		export let defaultValue = 'Type pls 😘';

		function editor(dom) {
      onMount(async () => {
		  return await Editor.make()
		    .config(ctx => {

          ctx.update(editorViewOptionsCtx, (prev) => ({
      ...prev,
      attributes: { class: 'milkdown-editor mx-auto outline-none margin-top-0 flex justify-center rounded-lg', spellcheck: 'false' },
    }))
		    })
		    .use(nord)

		    .use(commonmark)
       
		    .create()
    
      }) };
		

  
    

</script>


  <div >

    <NavBar {login}/>

  </div>

  <div use:editor />

  <div class="px-4">
      <div
        class="
          flex
          justify-center
          items-center
          bg-white
          mx-auto
          max-w-2xl
          rounded-lg
          my-16
          p-16
        "
      >
        <h1 class="text-2xl font-medium">I love taking notes</h1>
      </div>
    </div>


