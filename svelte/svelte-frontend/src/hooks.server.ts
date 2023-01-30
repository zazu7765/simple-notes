
/** @type {import('@sveltejs/kit').Handle} */
export function handle({ event, resolve }) {
	event.locals.user = event.cookies.get('jwt');


	return resolve(event);
}