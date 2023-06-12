import "../env-mock";
import {
  fireEvent,
  render,
  waitFor,
  screen,
  cleanup,
  act,
} from "@testing-library/svelte";
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

let server = setupMockserver();

server.listen();

let container: HTMLElement;
let component: Routes;
let rerender: (options: any) => void;
beforeEach(async () => {
  // await loadStateFromApi();
  server.resetHandlers();
  await handleRender();
});

async function handleRender() {
  let renderResult = render(Routes);
  await tick();
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

import {
  applicationState,
  loadStateFromApi,
  resetApplicationState,
} from "../../src/stores/application-state";
import Routes from "../../src/components/Routes.svelte";
import {
  defaultMockServerOptions,
  setNonAdmin,
  setupMockserver,
} from "../mock-server";
import { tick } from "svelte";

test("routes should all be rendered", async () => {
  setNonAdmin(false);
  await loadStateFromApi();
  let cardItems = container.querySelectorAll(".card.item");
  expect(cardItems.length).toBe(defaultMockServerOptions.routes.length);
});
test("routes should be grouped correctly", async () => {
  setNonAdmin(false);
  await loadStateFromApi();
  let cardItems = container.querySelectorAll(".card.item");
  expect(cardItems.length).toBe(defaultMockServerOptions.routes.length);
});
test("create route button should be hidden for non authenticated user", async () => {
  let createRouteButton = container.querySelector(".button.block.primary");
  expect(createRouteButton).toBeFalsy();
});
test("create route button should be visible for authenticated user", async () => {
  await loadStateFromApi();
  let createRouteButton = container.querySelector(".button.block.primary");
  expect(createRouteButton).toBeTruthy();
});
