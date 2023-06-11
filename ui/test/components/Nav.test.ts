import "../env-mock";
import { cleanup, fireEvent, render, waitFor } from "@testing-library/svelte";
import {
  afterAll,
  afterEach,
  assert,
  beforeAll,
  beforeEach,
  expect,
  Mock,
  test,
} from "vitest";

const server = setupMockserver();
server.listen();

let container: HTMLElement;
let component: Nav;
let rerender: (options: any) => void;
beforeEach(async () => {
  // await loadStateFromApi();
  server.resetHandlers();
  handleRender();
});

function handleRender() {
  let renderResult = render(Nav);
  component = renderResult.component;
  container = renderResult.container;
  rerender = renderResult.rerender;
}

afterEach(async () => {
  component.$destroy();
  await resetApplicationState();
  setNonAdmin(true);
  cleanup();
});

afterAll(() => server.close());

// may your conscience be of good judgement, so thee wont turn down the second boat
import Nav from "../../src/components/Nav.svelte";
import {
  applicationState,
  loadStateFromApi,
  resetApplicationState,
} from "../../src/stores/application-state";
import { setNonAdmin, setupMockserver } from "../mock-server";

test("should render sign in link", async () => {
  const button = container.querySelector('a.button.primary[href="/sign-in"]');

  expect(button).toBeTruthy();
});
test("should not render sign in link", async () => {
  await loadStateFromApi();
  const button = container.querySelector('a.button.primary[href="/sign-in"]');

  expect(button).toBeFalsy();
});
test("should render sign out button", async () => {
  await loadStateFromApi();
  const signOutButton = container.querySelector("button.button.secondary");

  expect(signOutButton).toBeTruthy();
});

test("should render users link", async () => {
  setNonAdmin(false);
  await loadStateFromApi();
  const usersButton = container.querySelector('a.button[href="/users"]');

  expect(usersButton).toBeTruthy();
});
test("should not render users link", async () => {
  setNonAdmin(true);
  await loadStateFromApi();
  const usersButton = container.querySelector('a.button[href="/users"]');

  expect(usersButton).toBeFalsy();
});

//TODO this doesnt want to work maybe implement mock window.location
// test("should reroute when clicking the link", async () => {
//   const link = container.querySelector('a.button.primary[href="/sign-in"]');

//   if (!link) {
//     throw new Error("Sign-in link not found");
//   }

//   fireEvent.click(link);

//   await waitFor(
//     () => {
//       // Wait for the expected href change to occur
//       return window.location.href.includes("/sign-in");
//     },
//     { timeout: 3000, interval: 1000 }
//   );

//   expect(window.location.href).toContain("/sign-in");
// });
