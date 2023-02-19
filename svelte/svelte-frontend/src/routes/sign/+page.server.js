import { redirect } from '@sveltejs/kit';

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
		const response = await fetch('http://localhost:81/auth/signup', {
			method: 'POST',
			body: JSON.stringify({ Name: name, Email: email, Password: password }),
			headers: {
				Accept: 'application/json',
				'Content-Type': 'application/json'
			}
		});
		if (!response) {
			return {
				success: false
			};
		}
		let dataPost;
		try {
			dataPost = await response.json();
			console.log(dataPost['Data']);
		} catch (error) {
			console.log(error);
			return {
				success: false
			};
		}
		//   let parsed = JSON.parse(total);
		//   console.log(dataPost['token']);

		cookies.set('jwt', dataPost['Data']['WebToken'], {
			path: '/',
			maxAge: dataPost['Data']['Expiration']
		});
		cookies.set('user', dataPost['Data']['Username'], {
			path: '/',
			maxAge: dataPost['Data']['Expiration']
		});
		throw redirect(302, '/user/'+dataPost['Data']['Username']);
	}
};
