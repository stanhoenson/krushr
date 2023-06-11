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

async function handleRender() {
  pointsOfInterest = [
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

//TODO leaflet map doesn't seem to load, maybe this cant be tested in vite
test("map should have right number of markers", async () => {
  // let markers = container.querySelectorAll(".map-marker");
  // for (let marker of markers) {
  //   console.log(marker);
  // }
  // console.log(markers);
  // expect(markers.length).toBe(pointsOfInterest.length);
});
test("marker should be removed after poi delete", async () => {});
test("marker should be added after poi create", async () => {});
test("marker should move on poi latlng update", async () => {});
test("poi latlng should update on map click", async () => {});
