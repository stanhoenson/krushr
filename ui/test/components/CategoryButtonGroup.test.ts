// @vitest-environment jsdom
import "../env-mock";
import { fireEvent, render } from "@testing-library/svelte";
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
} from "../../src/stores/application-state";

const categories = [
  { id: 1, name: "Default", position: 1 },
  { id: 2, name: "Nature", position: 2 },
] as Category[];
const server = setupServer(
  rest.get(GET_ALL_CATEGORIES_ENDPOINT, (req, res, ctx) => {
    return res(ctx.json(categories));
  })
);
server.listen();

let container: HTMLElement;
let component: CategoryButtonGroup;
beforeEach(async () => {
  await loadStateFromApi();
  server.resetHandlers();

  let renderResult = render(CategoryButtonGroup, {
    props: {
      disabled: false,
      selectedCategories: [],
      handleCategoryToggle: (category: Category) => {},
    },
  });
  component = renderResult.component;
  container = renderResult.container;
});

afterEach(() => {
  component.$destroy();
});

afterAll(() => server.close());

import CategoryButtonGroup from "../../src/components/CategoryButtonGroup.svelte";

test("should render category buttons and have correct text", async () => {
  const categoryButtons = container.querySelectorAll(".categories .category");

  let names: string[] = [];

  for (let button of categoryButtons) {
    if (button.textContent) {
      names.push(button.textContent);
    }
  }

  expect(names).toEqual(categories.map((entry) => entry.name));
  expect(categoryButtons.length).toBe(categories.length);
});
