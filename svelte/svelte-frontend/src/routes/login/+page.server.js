import { redirect } from '@sveltejs/kit';
import { id } from '$lib/notes'
export const prerender = false
export function load({ locals }) {
	if (locals.user) throw redirect(302, '/notes');
}

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request }) => {
		const data = await request.formData();
		const email = data.get('email');
		const password = data.get('password');
		const response = await fetch(import.meta.env.VITE_PUBLIC_URL+'/auth/login', {
			method: 'POST',
			body: JSON.stringify({ Email: email, Password: password }),
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			}
		});

		let dataPost;
		try {
			dataPost = await response.json();
		} catch (error) {
			
			return {
				success: false,
				message: dataPost['Data']
			};
		}
		if (dataPost['Status'] == 'error') {
			return {
				success: false,
				message: dataPost['Data']
			}
		}
		//   let parsed = JSON.parse(total);


		var now = new Date();
		var minutes = 30;
		const expirationTime = new Date(Date.now() + 1800000);
		cookies.set('jwt', dataPost['Data']['WebToken'], { path: '/', expires: expirationTime });
    cookies.set('user', dataPost['Data']['Username'], { path: '/', expires: expirationTime });
	// id.set(dataPost['Data']['UserId'])

		//   return{
		//     success: true,
		//     data: dataPost,

		//   }
		throw redirect(301, '/user/'+dataPost['Data']['Username']);
	}
};
