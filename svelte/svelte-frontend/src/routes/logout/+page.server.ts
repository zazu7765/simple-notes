import { redirect } from '@sveltejs/kit';


/** @type {import('./$types').LayoutServerLoad} */
export function load({cookies}) {
    cookies.delete('jwt');
    cookies.delete('user');

    throw redirect (303, "/login");
}