import { writable } from "svelte/store";
import { getAllCategories } from "../requests/categories";
import { getAllStatuses } from "../requests/statuses";
import { getMeUser } from "../requests/users";
import type { Category, Status, User } from "../types/models";

export type ApplicationState = {
  authenticatedUser: User | null;
  categories: Category[];
  statuses: Status[];
};

export const applicationState = writable<ApplicationState>({
  authenticatedUser: null,
  categories: [],
  statuses: [],
});

let initialized = false;

export async function loadStateFromApi() {
  let authenticatedUser = await loadAuthenticatedUserAction();
  let categories = await loadCategoriesAction();
  let statuses = await loadStatusesAction();

  applicationState.set({ authenticatedUser, categories, statuses });
}

async function loadStatusesAction() {
  try {
    let statuses = await getAllStatuses();
    return statuses;
  } catch (error) {
    return [];
  }
}
async function loadCategoriesAction() {
  try {
    let categories = await getAllCategories();
    return categories;
  } catch (error) {
    return [];
  }
}
async function loadAuthenticatedUserAction() {
  try {
    const user = await getMeUser();
    return user;
  } catch (error) {
    return null;
  }
}

async function loadStateFromLocalStorage() {
  try {
    let storedApplicationStateJSON = localStorage.getItem("applicationState");
    if (!storedApplicationStateJSON) {
      return;
    }
    let storedApplicationState: ApplicationState = JSON.parse(
      storedApplicationStateJSON
    );
    applicationState.set(storedApplicationState);
  } catch (error) {
    console.log(error);
  }
}

export async function resetApplicationState() {
  localStorage.removeItem("applicationState");
  applicationState.set({
    authenticatedUser: null,
    categories: [],
    statuses: [],
  });
}

export async function initializeApplicationState() {
  if (!initialized) {
    await loadStateFromLocalStorage();
    applicationState.subscribe((value) => {
      localStorage.setItem("applicationState", JSON.stringify(value));
    });
    await loadStateFromApi();
    initialized = true;
  }
}

// Call initializeUser() when the module is loaded
initializeApplicationState();
