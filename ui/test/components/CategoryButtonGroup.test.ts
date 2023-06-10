// @vitest-environment jsdom
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

import type { Category, Status, User } from "../../src/types/models";

const categories = [
  { id: 1, name: "Default", position: 1 },
  { id: 2, name: "Nature", position: 2 },
] as Category[];

const server = setupMockserver(
  Object.assign({}, defaultMockServerOptions, { categories })
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
      selectedCategories: [{ name: categories[0].name, position: 0 }],
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
import {
  applicationState,
  loadStateFromApi,
} from "../../src/stores/application-state";
import { defaultMockServerOptions, setupMockserver } from "../mock-server";

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
test("should render category buttons with appropriate class when they are selected", async () => {
  const selectedCategoryButton = container.querySelector(
    ".categories .category.selected"
  );

  expect(selectedCategoryButton).toBeTruthy();
  expect(selectedCategoryButton?.textContent).toBe(categories[0].name);
});
