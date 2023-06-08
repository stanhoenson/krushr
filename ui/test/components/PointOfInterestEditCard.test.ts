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
import { getAllCategories } from "../../src/requests/categories";
import { GET_ALL_CATEGORIES_ENDPOINT } from "../../src/requests/endpoints";
import type { Category } from "../../src/types/models";
const server = setupServer(
  rest.get(GET_ALL_CATEGORIES_ENDPOINT, (req, res, ctx) => {
    console.log(GET_ALL_CATEGORIES_ENDPOINT);
    return res(
      ctx.json([{ id: 1, name: "Default", position: 1 }] as Category[])
    );
  })
);
server.listen();

let route: PostRouteBody;
let container: HTMLElement;
let component: PointOfInterestEditCard;
beforeEach(() => {
  server.resetHandlers();
  route = {
    name: "",
    imageIds: [],
    details: [{ text: "" }],
    links: [{ url: "", text: "" }],
    categories: [],
    statusId: 1,
    pointsOfInterest: [
      {
        name: "",
        longitude: 0,
        latitude: 0,
        imageIds: [],
        details: [{ text: "" }],
        links: [{ url: "", text: "" }],
        categories: [],
        support: false,
      },
      {
        name: "",
        longitude: 0,
        latitude: 0,
        imageIds: [],
        details: [{ text: "" }],
        links: [{ url: "", text: "" }],
        categories: [],
        support: false,
      },
    ],
  };

  let renderResult = render(PointOfInterestEditCard, {
    props: {
      viewOnly: false,
      positionExchange: (a, b) => {},
      position: 0,
      route: route,
      pointOfInterest: route.pointsOfInterest[0],
    },
  });
  component = renderResult.component;
  container = renderResult.container;
});

afterEach(() => {
  component.$destroy();
});

afterAll(() => server.close());

import PointOfInterestEditCard from "../../src/components/PointOfInterestEditCard.svelte";
import type { PostRouteBody } from "../../src/types/request-bodies";

test("should delete link when delete link button is clicked", async () => {
  const deleteLinkButton = container.querySelector(
    ".multiple.links .icon.delete-icon"
  );

  if (!deleteLinkButton) {
    throw new Error("link delete button not found");
  }

  let linkCountBefore = route.pointsOfInterest[0].links.length;
  await fireEvent.click(deleteLinkButton);

  let linkCountAfter = route.pointsOfInterest[0].links.length;

  expect(linkCountAfter).lessThan(linkCountBefore);
});

test("should toggle category when category button is clicked", async () => {
  const button = container.querySelector(".categories .category");

  if (!button) {
    throw new Error("category button not found");
  }

  let categoriesBefore = route.pointsOfInterest[0].categories.length;
  await fireEvent.click(button);

  let categoriesAfter = route.pointsOfInterest[0].categories.length;
  let matchedCategory = route.pointsOfInterest[0].categories.find(
    (category) => category.name === button.textContent
  );

  expect(categoriesAfter).greaterThan(categoriesBefore);
  expect(matchedCategory).toBeTruthy();
});

test("should add link when add link button is clicked", async () => {
  const button = container.querySelector(".links .button.block.secondary");

  if (!button) {
    throw new Error("add link button not found");
  }

  let linksBefore = route.pointsOfInterest[0].links.length;
  fireEvent.click(button);
  let linksAfter = route.pointsOfInterest[0].links.length;

  expect(linksBefore).lessThan(linksAfter);
});

test("should add detail when add detail button is clicked", async () => {
  const button = container.querySelector(".details .button.block.secondary");

  if (!button) {
    throw new Error("add detail button not found");
  }

  let detailsBefore = route.pointsOfInterest[0].details.length;
  await fireEvent.click(button);

  let detailsAfter = route.pointsOfInterest[0].details.length;

  expect(detailsBefore).lessThan(detailsAfter);
});
test("should not find detail delete button", async () => {
  const button = container.querySelector(".multiple.details .icon.delete-icon");

  expect(button).toBeNull();
});

test("should delete detail when delete detail button is clicked", async () => {
  //add detail first
  const addDetailButton = container.querySelector(
    ".details .button.block.secondary"
  );

  if (!addDetailButton) {
    throw new Error("add detail button not found");
  }

  await fireEvent.click(addDetailButton);

  const button = container.querySelector(".multiple.details .icon.delete-icon");

  if (!button) {
    throw new Error("detail delete button not found");
  }

  let detailCountBefore = route.pointsOfInterest[0].details.length;
  await fireEvent.click(button);

  let detailCountAfter = route.pointsOfInterest[0].details.length;

  expect(detailCountAfter).lessThan(detailCountBefore);
});
test("should change name when name input is changed", async () => {});
test("should change detail text when detail text input is changed", async () => {});
test("should change link url and text when link url and text input are changed", async () => {});
