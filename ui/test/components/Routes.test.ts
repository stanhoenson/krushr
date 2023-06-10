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
  rest.get(GET_ALL_CATEGORIES_ENDPOINT, (req, res, ctx) => {
    return res(
      ctx.json([{ id: 1, name: "Default", position: 1 }] as Category[])
    );
  }),
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
  }),
  rest.get(GET_ALL_STATUSES_ENDPOINT, (req, res, ctx) => {
    return res(
      ctx.json([
        { id: 1, name: "Unpublished" },
        { id: 2, name: "Published" },
      ] as Status[])
    );
  })
);
server.listen();

let container: HTMLElement;
let component: Routes;
let rerender: (options: any) => void;
beforeEach(async () => {
  // await loadStateFromApi();
  server.resetHandlers();
  handleRender();
});

function handleRender() {
  let renderResult = render(Routes);
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

import Routes from "../../src/components/Routes.svelte";

test("routes should be grouped correctly", async () => {});
test("routes should render correctly", async () => {});
test("create route button should be hidden for non authenticated user", async () => {});
test("create route button should be visible for authenticated user", async () => {});
