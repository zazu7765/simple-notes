import { redirect } from '@sveltejs/kit';

export async function load({ locals, params }) {
	const name = await params.slug;
	let login = false;
	console.log('sss' + locals.user);
	const response = await fetch('http://localhost:81/notebooks/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	const responseNote = await fetch('http://localhost:81/notes/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	if (!locals.user) throw redirect(302, '/login');
	if (locals.user) login = true;
	
	return {
		name: locals.user,
		response: await response.json(),
		responseNote: await responseNote.json(),
		bool: login
	};
}
