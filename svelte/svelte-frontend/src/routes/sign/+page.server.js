import { redirect } from '@sveltejs/kit';
import {token} from '$lib/notes'
export function load({ locals }) {
	if (locals.user) throw redirect(302, '/user/'+locals.user);
}

/** @type {import('./$types').Actions} */
export const actions = {
	default: async ({ cookies, request, locals }) => {
		const data = await request.formData();
		const name = data.get('name');
		const email = data.get('email');
		const password = data.get('password');
		console.log(import.meta.env.VITE_PUBLIC_URL)
		const url = import.meta.env.VITE_PUBLIC_URL
		const response = await fetch(`${url}/auth/signup`, {
			
			method: 'POST',
			body: JSON.stringify({ Name: name, Email: email, Password: password }),
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			}
		});

		console.log("ggg")
		if (!response) {
			console.log("gggff")
			return {
				
				success: false
				
			};
		}
		let dataPost;
		try {
			dataPost = await response.json();
		} catch (error) {
			console.log("ffff")
			return {
				success: false,
				message: dataPost['Data']
			};
		}
		if (dataPost['Status'] == 'error') {
			console.log('jjjjj')
			return {
				success: false,
				message: dataPost['Data']
			}
		}
		//   let parsed = JSON.parse(total);
		console.log('jjjjj')
		const expirationTime = new Date(Date.now() + 1800000);
		cookies.set('jwt', dataPost['Data']['WebToken'], {
			secure: false,
			path: '/',
			expires: expirationTime
		});
		cookies.set('user', dataPost['Data']['Username'], {
			secure: false,
			path: '/',
			expires: expirationTime
		});
		console.log(cookies.get('jwt')+'kkkkkkkkkkk');
		throw redirect(302, '/user/'+dataPost['Data']['Username']);
	}
};
