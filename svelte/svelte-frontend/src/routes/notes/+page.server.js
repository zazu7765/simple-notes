import { redirect } from '@sveltejs/kit';



export function load({ locals}) {
	let login = false;
	if (!locals.user) throw redirect(302, '/login')
	if (locals.user) login=true
	return{
		bool: login,
	}
}
