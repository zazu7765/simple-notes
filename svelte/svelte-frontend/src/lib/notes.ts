import { writable } from "svelte/store";


let write = writable(null);
let storage = writable('notebook');
let state = writable('0');
let currentNote = writable(["0","not Fount"]);
let notebooks = writable(null);
export {currentNote}
export {write}
export {storage}
export {state}