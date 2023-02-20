import { redirect } from '@sveltejs/kit';
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
		const response = await fetch('http://localhost:81/auth/login', {
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
		//   console.log(dataPost['token']);

		console.log(dataPost['Data']['Expiration']+"timestamp");
		cookies.set('jwt', dataPost['Data']['WebToken'], { path: '/', maxAge: dataPost['Data']['Expiration'] });
    cookies.set('user', dataPost['Data']['Username'], { path: '/', maxAge: dataPost['Data']['Expiration'] });

		//   return{
		//     success: true,
		//     data: dataPost,

		//   }
		throw redirect(301, '/user/'+dataPost['Data']['Username']);
	}
};
