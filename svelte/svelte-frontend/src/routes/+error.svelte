<script>
    import { page } from '$app/stores';
	import { redirect } from '@sveltejs/kit';
    throw redirect(302, '/logout');
  </script>
  
