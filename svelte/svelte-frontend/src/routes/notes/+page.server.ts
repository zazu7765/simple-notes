import { redirect } from '@sveltejs/kit';


/** @type {import('./$types').LayoutServerLoad} */
export function load({cookies}) {
 

    throw redirect (303, "/user/notebooks");
}