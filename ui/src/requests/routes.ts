import axios from "axios";
import type { Route } from "../types/models";
import type { PostRouteBody, PutRouteBody } from "../types/request-bodies";
import {
  BASE_URL,
  CREATE_ROUTE_ENDPOINT,
  DELETE_ROUTE_BY_ID_ENDPOINT,
  GET_ALL_ROUTES_ENDPOINT,
  GET_ROUTE_BY_ID_ENDPOINT,
  UPDATE_ROUTE_BY_ID_ENDPOINT,
} from "./endpoints";

axios.defaults.withCredentials = true;
export async function getAllRoutes(): Promise<Route[]> {
  const response = await axios.get<Route[]>(GET_ALL_ROUTES_ENDPOINT);
  return response.data;
}

export async function getRouteById(id: number): Promise<Route> {
  const response = await axios.get<Route>(GET_ROUTE_BY_ID_ENDPOINT(id));
  return response.data;
}

export async function createRoute(requestBody: PostRouteBody): Promise<Route> {
  const response = await axios.post<Route>(CREATE_ROUTE_ENDPOINT, requestBody);
  return response.data;
}

export async function updateRoute(
  id: number,
  requestBody: PutRouteBody
): Promise<Route> {
  const response = await axios.put<Route>(
    UPDATE_ROUTE_BY_ID_ENDPOINT(id),
    requestBody
  );
  return response.data;
}

export async function deleteRouteById(id: number): Promise<void> {
  await axios.delete(DELETE_ROUTE_BY_ID_ENDPOINT(id));
}
