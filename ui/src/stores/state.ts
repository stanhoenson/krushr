import { writable } from "svelte/store";
import { getMeUser } from "../requests/users";
import type { User } from "../types/models";

export type ApplicationState = {

};

export const authenticatedUser = writable<User | null>(null);

let initialized = false;

async function fetchUser() {
  try {
    const result = await getMeUser();
    authenticatedUser.set(result);
  } catch (e) {
    console.error(e);
  }
}

function initializeUser() {
  if (!initialized) {
    fetchUser();
    initialized = true;
  }
}

// Call initializeUser() when the module is loaded
initializeUser();
