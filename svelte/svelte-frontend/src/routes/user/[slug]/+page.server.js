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
	if (!locals.user) 
	if (locals.user) login = true;
	 let responseJ = await response.json()
	 let responseNoteJ = await responseNote.json()
	 if (responseNoteJ["Status"] === "error" || responseJ["Status"] === "error"){
		throw redirect(302, '/logout');
	 }
	return {
		token: locals.token,
		name: locals.user,
		response: responseJ,
		responseNote: responseNoteJ,
		bool: login
	};
}
