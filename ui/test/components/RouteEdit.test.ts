import "../env-mock";
import {
  cleanup,
  fireEvent,
  screen,
  render,
  waitFor,
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

const server = setupMockserver();
server.listen();

let container: HTMLElement;
let component: RouteEdit;
let rerender: (options: any) => void;
beforeEach(async () => {
  // await loadStateFromApi();
  server.resetHandlers();
  await handleRender();
});

async function handleRender() {
  let renderResult = render(RouteEdit, {});
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

import RouteEdit from "../../src/components/RouteEdit.svelte";
import {
  applicationState,
  loadStateFromApi,
  resetApplicationState,
} from "../../src/stores/application-state";
import { setNonAdmin, setupMockserver } from "../mock-server";
import { tick } from "svelte";

test("buttons and inputs should be disabled for non authenticated user", async () => {
  let inputs = container.querySelectorAll<HTMLButtonElement | HTMLInputElement>(
    "input button"
  );
  for (let input of inputs) {
    expect(input.disabled).toBeTruthy();
  }
});
test("buttons and inputs should be enabled for authenticated user", async () => {
  await loadStateFromApi();
  await tick();
  let inputs = container.querySelectorAll<HTMLButtonElement | HTMLInputElement>(
    "input button"
  );
  for (let input of inputs) {
    expect(input.disabled).toBeTruthy();
  }
});
test("poi delete button should be disabled with 2 poi's", async () => {
  let buttons = container.querySelectorAll("button");
  let poiDeleteButtonCount = 0;
  for (let button of buttons) {
    if (button.textContent === "Delete point of interest") {
      poiDeleteButtonCount++;
      expect(button.disabled).toBeTruthy();
    }
  }
  expect(poiDeleteButtonCount).toEqual(2);
});
test("poi delete button should delete poi", async () => {
  await loadStateFromApi();
  await tick();
  let buttons = container.querySelectorAll("button");
  let newPoiButton = null;
  for (let button of buttons) {
    if (button.textContent === "New point of interest") {
      newPoiButton = button;
    }
  }
  if (!newPoiButton) throw new Error("no new poi button found");

  await fireEvent.click(newPoiButton);
  await tick();

  let poisBefore = 0;
  let firstPoiButton = null;
  for (let button of buttons) {
    if (button.textContent === "Delete point of interest") {
      poisBefore++;
      if (!firstPoiButton) firstPoiButton = button;
    }
  }

  if (!firstPoiButton) throw new Error("no poi delete buttons found");

  await fireEvent.click(firstPoiButton);
  await tick();

  let poisAfter = 0;
  for (let button of buttons) {
    if (button.textContent === "Delete point of interest") {
      poisAfter++;
    }
  }
  expect(poisBefore).toBeGreaterThan(poisAfter);
});
