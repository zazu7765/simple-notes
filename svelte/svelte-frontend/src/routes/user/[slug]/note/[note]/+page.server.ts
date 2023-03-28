import { redirect } from '@sveltejs/kit';
import { storage  } from '$lib/notes';
export async function load({ locals, params }) {


	let login = false;
	console.log('sss' + locals.user);
	const response = await fetch(import.meta.env.VITE_PUBLIC_URL+'/notebooks/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL+'/notes/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	if (!locals.user) if (locals.user) login = true;
	let responseJ = await response.json();
	let responseNoteJ = await responseNote.json();
	if (responseNoteJ['Status'] === 'error' || responseJ['Status'] === 'error') {
		throw redirect(302, '/logout');
	}
	const formData = new FormData();
	formData.append('id', params.note);
	const responseNotes = await fetch(import.meta.env.VITE_PUBLIC_URL+'/notes', {
		method: 'POST',
		body: formData,
		headers: {
			Authorization: 'Bearer ' + locals.token
		}
	});
	const datanote = await responseNotes.json()
	if (datanote['Status'] == 'error') {
		throw redirect(302, '/logout');
	}

	return {
		token: locals.token,
		name: locals.user,
		response: responseJ,
		responseNote: responseNoteJ,
		bool: login,
		notes: datanote
	};
}
