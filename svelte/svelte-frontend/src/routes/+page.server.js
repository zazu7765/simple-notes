

/** @type {import('./$types').PageServerLoad} */

export async function load ({ locals}) {
	let login = false;

	if (await locals.user && await locals.token){
		login=true}
    
	return{ user:{
		bool: login
	}}
}