

/** @type {import('./$types').LayoutServerLoad} */
export function load({ locals }) {

	return {
		user: locals.user && {
			username: locals['user']['name'],
			email: locals['user']['email'],
			notes: locals['user']['Notebook'],
		}
	};
}