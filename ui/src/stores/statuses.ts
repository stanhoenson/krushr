import { writable } from "svelte/store";
import { getAllCategories } from "../requests/categories";
import { getAllStatuses } from "../requests/statuses";
import { getMeUser } from "../requests/users";
import type { Status } from "../types/models";

export const statusesStore = writable<Status[]>([]);

let initialized = false;

export async function initializeStatuses() {
  if (!initialized) {
    statusesStore.set(await getAllStatuses());
    initialized = true;
  }
}

// Call initializeUser() when the module is loaded
initializeStatuses();
