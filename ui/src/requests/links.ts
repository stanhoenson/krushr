import axios from "axios";
import type { Link } from "../types/models";
import {
  DELETE_LINK_BY_ID_ENDPOINT,
  GET_ALL_LINKS_ENDPOINT,
} from "./endpoints";

// Get all links
export async function getAllLinks(): Promise<Link[]> {
  const response = await axios.get<Link[]>(GET_ALL_LINKS_ENDPOINT);
  return response.data;
}

// Delete a link by ID
export async function deleteLinkById(id: number): Promise<void> {
  await axios.delete(DELETE_LINK_BY_ID_ENDPOINT(id), { withCredentials: true });
}
