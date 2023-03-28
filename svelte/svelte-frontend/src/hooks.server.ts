


/** @type {import('@sveltejs/kit').Handle} */
export function handle({ event, resolve}) {
	console.log('ass1'+ event.cookies.get('user'))
	event.locals.user = event.cookies.get('user');
	event.locals.token = event.cookies.get('jwt');
	if (event.locals.user!= event.cookies.get('user') || event.locals.token != event.cookies.get('jwt')){
		
		console.log('ass')
		return new Response('Redirect', {status: 303, headers: { Location: '/login' }});
	}
	

	return resolve(event);
}