<script>
    import "../app.css";
    import { SvelteToast } from '@zerodevx/svelte-toast'

//   import "flowbite/dist/flowbite.css";
const options = {
    
  }

</script>
<slot />
<SvelteToast {options} />