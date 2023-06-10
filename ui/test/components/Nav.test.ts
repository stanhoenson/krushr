// @vitest-environment jsdom
import "../env-mock";
import { fireEvent, render, waitFor } from "@testing-library/svelte";
import {
  assert,
  expect,
  test,
  Mock,
  beforeAll,
  beforeEach,
  afterAll,
  afterEach,
} from "vitest";

import { setupServer } from "msw/node";
import { rest } from "msw";
import {
  GET_ALL_CATEGORIES_ENDPOINT,
  GET_ALL_STATUSES_ENDPOINT,
  GET_ME_USER_ENDPOINT,
} from "../../src/requests/endpoints";
import type { Category, Status, User } from "../../src/types/models";
import {
  applicationState,
  loadStateFromApi,
  resetApplicationState,
} from "../../src/stores/application-state";

let nonAdmin = true;

const server = setupServer(
  rest.get(GET_ME_USER_ENDPOINT, (req, res, ctx) => {
    console.log({ nonAdmin });
    if (nonAdmin) {
      return res(
        ctx.json({
          id: 1,
          email: "test@test.com",
          role: { id: 1, name: "Creator" },
          roleId: 1,
        } as User)
      );
    }
    return res(
      ctx.json({
        id: 1,
        email: "test@test.com",
        role: { id: 2, name: "Admin" },
        roleId: 2,
      } as User)
    );
  })
);
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
  nonAdmin = true;
});

afterAll(() => server.close());

import Nav from "../../src/components/Nav.svelte";

test("should render sign in link", async () => {
  const button = container.querySelector('a.button.primary[href="/sign-in"]');

  expect(button).toBeTruthy();
});

//TODO this doesnt want to work maybe implement mock window.location
test("should reroute when clicking the link", async () => {
  const link = container.querySelector('a.button.primary[href="/sign-in"]');

  if (!link) {
    throw new Error("Sign-in link not found");
  }

  fireEvent.click(link);

  await waitFor(
    () => {
      // Wait for the expected href change to occur
      return window.location.href.includes("/sign-in");
    },
    { timeout: 3000, interval: 1000 }
  );

  expect(window.location.href).toContain("/sign-in");
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
  nonAdmin = false;
  await loadStateFromApi();
  const usersButton = container.querySelector('a.button[href="/users"]');

  expect(usersButton).toBeTruthy();
});
test("should not render users link", async () => {
  nonAdmin = true;
  await loadStateFromApi();
  const usersButton = container.querySelector('a.button[href="/users"]');

  expect(usersButton).toBeFalsy();
});
