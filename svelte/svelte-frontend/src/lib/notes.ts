import { writable } from "svelte/store";


let write = writable(null);
let storage = writable(["notebook", 0]);
export {write}
export {storage}