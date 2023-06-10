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
  authenticatedUserNonAdminUser: User;
  authenticatedUserAdminUser: User;
  statuses: Status[];
  routes: Route[];
};

const categories = [{ id: 1, name: "Default", position: 1 }] as Category[];
const statuses = [
  { id: 1, name: "Unpublished" },
  { id: 2, name: "Published" },
] as Status[];
const authenticatedUserNonAdminUser = {
  id: 1,
  email: "test@test.com",
  role: { id: 1, name: "Creator" },
  roleId: 1,
} as User;
const authenticatedUserAdminUser = {
  id: 1,
  email: "admin@admin.com",
  role: { id: 1, name: "Admin" },
  roleId: 1,
} as User;
export const defaultMockServerOptions: MockServerOptions = {
  categories: categories,
  authenticatedUserNonAdminUser,
  authenticatedUserAdminUser,
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
      user: authenticatedUserAdminUser,
      userId: authenticatedUserAdminUser.id,
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
          user: authenticatedUserAdminUser,
          userId: authenticatedUserAdminUser.id,
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
          user: authenticatedUserAdminUser,
          userId: authenticatedUserAdminUser.id,
          support: false,
        },
      ] as PointOfInterest[],
    } as Route,
  ],
};

export function setNonAdmin(value: boolean) {
  nonAdmin = value;
}

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
      return res(ctx.json(options.routes));
    }),
    rest.get(GET_ALL_CATEGORIES_ENDPOINT, (req, res, ctx) => {
      return res(ctx.json(options.categories));
    }),
    rest.get(GET_ME_USER_ENDPOINT, (req, res, ctx) => {
      if (nonAdmin) {
        return res(ctx.json(options.authenticatedUserNonAdminUser));
      }
      return res(ctx.json(options.authenticatedUserAdminUser));
    }),
    rest.get(GET_ALL_STATUSES_ENDPOINT, (req, res, ctx) => {
      return res(ctx.json(options.statuses));
    })
  );
}
