


export function load({ locals}) {
	let login = false;

	if (locals.user) login=true
    
	return{
		bool: login,
	}
}