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
server.listen({ onUnhandledRequest: "error" });

let container: HTMLElement;
let component: LeafletMap;
let rerender: (options: any) => void;
beforeEach(async () => {
  // await loadStateFromApi();
  server.resetHandlers();
  await handleRender();
});

let pointsOfInterest = [
  {
    name: "",
    longitude: 4.711,
    latitude: 52.0164830842629,
    imageIds: [],
    details: [{ text: "" }],
    links: [{ url: "", text: "" }],
    categories: [],
    support: false,
  },
  {
    name: "",
    longitude: 4.7,
    latitude: 52.011,
    imageIds: [],
    details: [{ text: "" }],
    links: [{ url: "", text: "" }],
    categories: [],
    support: false,
  },
];

let latitude = 52.0164830842629;
let longitude = 4.711;

async function handleRender() {
  pointsOfInterest = [
    {
      name: "",
      longitude,
      latitude,
      imageIds: [],
      details: [{ text: "" }],
      links: [{ url: "", text: "" }],
      categories: [],
      support: false,
    },
    {
      name: "",
      longitude: 4.7,
      latitude: 52.011,
      imageIds: [],
      details: [{ text: "" }],
      links: [{ url: "", text: "" }],
      categories: [],
      support: false,
    },
  ];

  let renderResult = render(LeafletMap, {
    props: {
      disabled: false,
      allPointsOfInterest: pointsOfInterest,
      position: 0,
      longitude: 4.711,
      latitude: 52.0164830842629,
    },
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

import {
  applicationState,
  loadStateFromApi,
  resetApplicationState,
} from "../../src/stores/application-state";
import LeafletMap from "../../src/components/LeafletMap.svelte";
import { setNonAdmin, setupMockserver } from "../mock-server";
import { tick } from "svelte";

//TODO leaflet map doesn't seem to load, maybe this cant be tested in vite
test("map should have right number of markers", async () => {
  let markers = container.querySelectorAll(".map-marker");
  expect(markers.length).toBe(pointsOfInterest.length);
});
test("marker should be removed after poi delete", async () => {
  let markersBefore = container.querySelectorAll(".map-marker").length;
  pointsOfInterest.splice(0, 1);
  component.$set({ allPointsOfInterest: pointsOfInterest });
  await tick();
  let markersAfter = container.querySelectorAll(".map-marker").length;
  expect(markersAfter).toBeLessThan(markersBefore);
});
test("marker should be added after poi create", async () => {
  let markersBefore = container.querySelectorAll(".map-marker").length;
  pointsOfInterest.push({
    name: "",
    longitude: 4.711,
    latitude: 52.0164830842629,
    imageIds: [],
    details: [{ text: "" }],
    links: [{ url: "", text: "" }],
    categories: [],
    support: false,
  });
  component.$set({ allPointsOfInterest: pointsOfInterest });
  await tick();
  let markersAfter = container.querySelectorAll(".map-marker").length;
  expect(markersAfter).toBeGreaterThan(markersBefore);
});
test("marker should move on poi latlng update", async () => {
  let firstMarker: any = container.querySelector(".map-marker");
  if (!firstMarker) {
    throw new Error("no marker found");
  }

  const positionBefore = firstMarker["_leaflet_pos"];
  pointsOfInterest[0].latitude += 0.1;
  pointsOfInterest[0].longitude += 0.1;
  component.$set({ allPointsOfInterest: pointsOfInterest });
  await tick();
  firstMarker = container.querySelector(".map-marker");
  if (!firstMarker) {
    throw new Error("no marker found");
  }
  const positionAfter = firstMarker["_leaflet_pos"];
  expect(positionBefore).not.toEqual(positionAfter);
});
test("poi latlng should update on map click", async () => {
  let map = container.querySelector(".leaflet-container");
  if (!map) throw new Error("map element not found");

  //hardcoded right context index
  let latitudeBefore = component.$$.ctx[3];
  let longitudeBefore = component.$$.ctx[4];

  await fireEvent.click(map);

  let latitudeAfter = component.$$.ctx[3];
  let longitudeAfter = component.$$.ctx[4];

  expect(latitudeBefore).not.toBe(latitudeAfter);
  expect(longitudeBefore).not.toBe(longitudeAfter);
});
