import { writable } from "svelte/store";

export const markers = writable<L.Marker[]>([]);
