import { writable } from "svelte/store";


let write = writable(null);
let storage = writable('notebook');
let state = writable('0');
let currentNote = writable(["0","not Fount"]);
let notebooks = writable(null);
let arr = writable(['0']);
let arrS = writable(['0']);
let id = writable(null);
let token = writable(null);
let env = "localhost/";
export {currentNote}
export {write}
export {storage}
export {state}
export {arr}
export {arrS}
export {id}
export{token}