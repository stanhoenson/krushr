// @vitest-environment jsdom
import "../env-mock";
import { cleanup, fireEvent, render, waitFor } from "@testing-library/svelte";
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
  handleRender();
});

function handleRender() {
  let renderResult = render(RouteEdit, {
    props: {},
  });
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

test("buttons and inputs should be disabled for non authenticated user", async () => {});
test("buttons and inputs should be enabled for authenticated user", async () => {});
test("url should change on save", async () => {});
test("should redirect on route delete", async () => {});
test("poi delete button should be disabled with 2 poi's", async () => {});
test("poi delete button should delete poi", async () => {});
