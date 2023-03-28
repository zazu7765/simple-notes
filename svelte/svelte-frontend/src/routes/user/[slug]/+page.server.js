import { redirect } from '@sveltejs/kit';


export async function load({ locals, params,cookies, }) {
	const name = await params.slug;
	let login = false;
	console.log('fdffff'+locals.token)
	const response = await fetch(import.meta.env.VITE_PUBLIC_URL+'/notebooks/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	console.log('+')
	const responseNote = await fetch(import.meta.env.VITE_PUBLIC_URL+'/notes/all', {
		method: 'GET',
		headers: {
			Authorization: 'Bearer ' + locals.token,
			'Content-Type': 'application/json'
		}
	});
	console.log('fdffff')
	if (!locals.user) throw redirect(302, '/logout');
	if (locals.user) login = true;
	console.log('fdffff')
	 let responseJ = await response.json()
	 let responseNoteJ = await responseNote.json()
	 if (responseNoteJ["Status"] === "error" || responseJ["Status"] === "error"){
		
		console.log('eror'+cookies.get('jwt'))
		throw redirect(302, '/logout');
	 }
	 console.log('fdffff')
	 console.log('suc'+locals.token)
	return {
		token: locals.token,
		name: locals.user,
		response: responseJ,
		responseNote: responseNoteJ,
		bool: login
	};
}
