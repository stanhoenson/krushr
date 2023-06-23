import "../env-mock";
import {
  act,
  cleanup,
  fireEvent,
  render,
  screen,
  waitFor,
} from "@testing-library/svelte";
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

  let ownRouteCount = 0;
  let othersRouteCount = 0;
  let publishedCount = 0;
  let unpublishedCount = 0;
  for (const route of defaultMockServerOptions.routes) {
    if (route.statusId === 1) unpublishedCount++;
    else publishedCount++;
    if (route.userId === 2) ownRouteCount++;
    else othersRouteCount++;
  }

  let grids = container.querySelectorAll(".grid");
  if (!grids) throw new Error("no grid containers found");

  let measured = 0;
  let measuredUnpublished = 0;
  let measuredOthers = 0;

  for (let grid of grids) {
    let previousSibling = grid.previousElementSibling;
    if (!previousSibling) continue;
    let textContent = previousSibling.textContent;
    if (!textContent) continue;

    let unpublished = textContent.toLowerCase().includes("unpublished");
    let others = textContent.toLowerCase().includes("others");

    let routeCount = grid.querySelectorAll(".card.item").length;

    measured += routeCount;
    if (unpublished) measuredUnpublished += routeCount;
    if (others) measuredOthers += routeCount;
  }

  expect(measured).toBe(defaultMockServerOptions.routes.length);
  expect(measuredUnpublished).toBe(unpublishedCount);
  expect(measured - measuredUnpublished).toBe(publishedCount);
  expect(measuredOthers).toBe(othersRouteCount);
  expect(measured - measuredOthers).toBe(ownRouteCount);
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
