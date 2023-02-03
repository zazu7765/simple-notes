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
		import { Editor, rootCtx, defaultValueCtx } from "@milkdown/core";
		import { commonmark } from "@milkdown/preset-commonmark";
		import { nord } from "@milkdown/theme-nord";


		export let defaultValue;

		function editor(dom) {
      onMount(async () => {
		  Editor.make()
		    .config(ctx => {
		      ctx.set(rootCtx, dom);
		      ctx.set(defaultValueCtx, defaultValue);
		    })
		    .use(nord)

		    .use(commonmark)
       
		    .create();
        
        )};catch(error){
  
  console.log(error)
    }
		}

  
    

</script>


  <div>

    <NavBar {login}/>

  </div>
  <style>
  </style>
  
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


