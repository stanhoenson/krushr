import "./env-mock";
import type {
  Category,
  Detail,
  Image,
  Link,
  PointOfInterest,
  Route,
  Status,
  User,
} from "../src/types/models";
import { setupServer } from "msw/node";
import { rest } from "msw";
import {
  BASE_URL,
  GET_ALL_CATEGORIES_ENDPOINT,
  GET_ALL_ROUTES_ENDPOINT,
  GET_ALL_STATUSES_ENDPOINT,
  GET_IMAGEDATA_BY_ID_ENDPOINT,
  GET_IMAGE_BY_ID_ENDPOINT,
  GET_ME_USER_ENDPOINT,
} from "../src/requests/endpoints";
import { readFileSync } from "fs";

export type MockServerOptions = {
  categories: Category[];
  authenticatedNonAdminUser: User;
  authenticatedAdminUser: User;
  statuses: Status[];
  routes: Route[];
};

const categories = [{ id: 1, name: "Default", position: 1 }] as Category[];
const statuses = [
  { id: 1, name: "Unpublished" },
  { id: 2, name: "Published" },
] as Status[];
const authenticatedNonAdminUser = {
  id: 2,
  email: "test@test.com",
  role: { id: 1, name: "Creator" },
  roleId: 1,
} as User;
const authenticatedAdminUser = {
  id: 1,
  email: "admin@admin.com",
  role: { id: 1, name: "Admin" },
  roleId: 1,
} as User;
export const defaultMockServerOptions: MockServerOptions = {
  categories: categories,
  authenticatedNonAdminUser: authenticatedNonAdminUser,
  authenticatedAdminUser: authenticatedAdminUser,
  statuses: statuses,
  routes: [
    {
      id: 1,
      name: "Route 1",
      images: [{ id: 1, path: "data/files/image.png" }] as Image[],
      details: [{ id: 1, text: "detail text" }] as Detail[],
      links: [{ id: 1, url: "www.test.com", text: "testLink" }] as Link[],
      categories: categories,
      status: statuses[0],
      statusId: statuses[0].id,
      user: authenticatedAdminUser,
      userId: authenticatedAdminUser.id,
      distance: 10,
      pointsOfInterest: [
        {
          id: 1,
          name: "Poi 1",
          longitude: 10.1,
          latitude: 5.1,
          images: [{ id: 1, path: "data/files/image.png" }] as Image[],
          details: [{ id: 1, text: "detail text" }] as Detail[],
          links: [] as Link[],
          user: authenticatedAdminUser,
          userId: authenticatedAdminUser.id,
          support: false,
        },
        {
          id: 2,
          name: "Poi 2",
          longitude: 12.1,
          latitude: 6.1,
          images: [{ id: 1, path: "data/files/image.png" }] as Image[],
          details: [{ id: 1, text: "detail text" }] as Detail[],
          links: [] as Link[],
          user: authenticatedAdminUser,
          userId: authenticatedAdminUser.id,
          support: false,
        },
      ] as PointOfInterest[],
    },
    {
      id: 2,
      name: "Route 2",
      images: [{ id: 1, path: "data/files/image.png" }] as Image[],
      details: [{ id: 1, text: "detail text" }] as Detail[],
      links: [{ id: 1, url: "www.test.com", text: "testLink" }] as Link[],
      categories: categories,
      status: statuses[1],
      statusId: statuses[1].id,
      user: authenticatedAdminUser,
      userId: authenticatedAdminUser.id,
      distance: 10,
      pointsOfInterest: [
        {
          id: 1,
          name: "Poi 1",
          longitude: 10.1,
          latitude: 5.1,
          images: [{ id: 1, path: "data/files/image.png" }] as Image[],
          details: [{ id: 1, text: "detail text" }] as Detail[],
          links: [] as Link[],
          user: authenticatedAdminUser,
          userId: authenticatedAdminUser.id,
          support: false,
        },
        {
          id: 2,
          name: "Poi 2",
          longitude: 12.1,
          latitude: 6.1,
          images: [{ id: 1, path: "data/files/image.png" }] as Image[],
          details: [{ id: 1, text: "detail text" }] as Detail[],
          links: [] as Link[],
          user: authenticatedAdminUser,
          userId: authenticatedAdminUser.id,
          support: false,
        },
      ] as PointOfInterest[],
    },
  ] as Route[],
};

export function setNonAdmin(value: boolean) {
  nonAdmin = value;
}

let mockOSMRResponse = {
  code: "Ok",
  waypoints: [
    {
      hint: "abc123",
      distance: 0.0,
      location: [40.712776, -74.005974],
      name: "New York City",
    },
    {
      hint: "def456",
      distance: 0.0,
      location: [34.052235, -118.243683],
      name: "Los Angeles",
    },
  ],
  routes: [
    {
      weight_name: "routability",
      geometry: "polyline",
      distance: 5000,
      duration: 3600,
      legs: [
        {
          summary: "",
          weight: 5000,
          duration: 3600,
          steps: [],
          distance: 5000,
        },
      ],
    },
  ],
  uuid: "1234567890",
};
let nonAdmin = true;
//TODO options maybe not so good of a name
export function setupMockserver(
  options: MockServerOptions = defaultMockServerOptions
) {
  return setupServer(
    rest.get(`${BASE_URL}/imagedata/:id`, (req, res, ctx) => {
      const headers = {
        "Content-Type": "image/png",
      };
      res(ctx.set(headers), ctx.body(readFileSync("./test.png")));
    }),
    rest.get(GET_ALL_ROUTES_ENDPOINT, (req, res, ctx) => {
      console.log("huhhh");
      return res(ctx.json(options.routes));
    }),
    rest.get(GET_ALL_CATEGORIES_ENDPOINT, (req, res, ctx) => {
      return res(ctx.json(options.categories));
    }),
    rest.get(GET_ME_USER_ENDPOINT, (req, res, ctx) => {
      if (nonAdmin) {
        return res(ctx.json(options.authenticatedNonAdminUser));
      }
      return res(ctx.json(options.authenticatedAdminUser));
    }),
    rest.get(GET_ALL_STATUSES_ENDPOINT, (req, res, ctx) => {
      return res(ctx.json(options.statuses));
    }),
    rest.get(/\/route\/v1/, (req, res, ctx) => {
      return res(ctx.json(mockOSMRResponse));
    })
  );
}
