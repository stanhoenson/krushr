import { writable } from "svelte/store";
import { getAllCategories } from "../requests/categories";
import { getMeUser } from "../requests/users";
import type { Category, User } from "../types/models";

export const categories = writable<Category[]>([]);

let initialized = false;

export async function initializeCategories() {
  if (!initialized) {
    categories.set(await getAllCategories());
    initialized = true;
  }
}

// Call initializeUser() when the module is loaded
initializeCategories();
