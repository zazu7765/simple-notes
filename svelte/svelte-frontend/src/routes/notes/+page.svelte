

<script lang='ts'>
    import Tiptap from '../../lib/TipTap/tiptapApp.svelte';
    import NavBar from '../../lib/NavBar.svelte';
    export let data:any;
    let login = data;

  </script>
  <div>
    <NavBar {login}/>
    
  </div>
  
<div>
  <Tiptap/>
    
    
</div>