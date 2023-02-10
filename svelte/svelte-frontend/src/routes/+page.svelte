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




	
		

  
    

</script>


  <div >

    <NavBar {login}/>

  </div>



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


