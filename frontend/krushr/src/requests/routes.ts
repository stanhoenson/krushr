import axios from "axios";
import type { Route } from "../types/models";
import type { PostRouteBody, PutRouteBody } from "../types/request-bodies";
import { BASE_URL } from "./endpoints";

axios.defaults.withCredentials = true;
export async function getAllRoutes(): Promise<Route[]> {
  const response = await axios.get<Route[]>(`${BASE_URL}/routes`);
  return response.data;
}

export async function getRouteById(id: number): Promise<Route> {
  const response = await axios.get<Route>(`${BASE_URL}/routes/${id}`);
  return response.data;
}

export async function createRoute(requestBody: PostRouteBody): Promise<Route> {
  const response = await axios.post<Route>(`${BASE_URL}/routes`, requestBody);
  return response.data;
}

export async function updateRoute(
  id: number,
  requestBody: PutRouteBody
): Promise<Route> {
  const response = await axios.put<Route>(
    `${BASE_URL}/routes/${id}`,
    requestBody
  );
  return response.data;
}

export async function deleteRouteById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/routes/${id}`);
}
