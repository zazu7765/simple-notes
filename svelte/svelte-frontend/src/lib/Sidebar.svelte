<script script lang="ts">
    let call = async ()=>{
		const responseNote = await fetch('http://localhost:81/notebooks/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});

    }
    export let name: String[];

</script>


<li>
    <a href={name[0]} class="flex items-center p-2 text-base font-normal text-gray-900 rounded-lg dark:text-white hover:bg-gray-100 dark:hover:bg-gray-700">
       <svg aria-hidden="true" class="w-6 h-6 text-gray-500 transition duration-75 dark:text-gray-400 group-hover:text-gray-900 dark:group-hover:text-white" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path d="M2 10a8 8 0 018-8v8h8a8 8 0 11-16 0z"></path><path d="M12 2.252A8.014 8.014 0 0117.748 8H12V2.252z"></path></svg>
       <span class="ml-3">{name[1]}</span>
    </a>
 </li>