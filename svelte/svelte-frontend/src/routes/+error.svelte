<script>
    import { page } from '$app/stores';
	import { redirect } from '@sveltejs/kit';
    console.log(page.error)
    throw redirect(302, '/logout');
  </script>
  
