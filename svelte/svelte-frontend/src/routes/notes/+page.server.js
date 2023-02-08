import { redirect } from '@sveltejs/kit';



export async function load({ locals}) {
	let login = false;
	console.log(locals.user);
	const response = await fetch('http://localhost:81/note',{
		method: 'GET',
		headers: {
			'Authorization': 'Bearer ' + locals.user,
			'Content-Type': 'application/json'
		}
	})
	if (!locals.user) throw redirect(302, '/login')
	if (locals.user) login=true

	return{
		response:(await response.json()),
		bool: login,

	}
	
}
