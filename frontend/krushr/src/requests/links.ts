import axios from "axios";
import type { Link } from "../types/models";
import { BASE_URL } from "./endpoints";

// Get all links
export async function getAllLinks(): Promise<Link[]> {
  const response = await axios.get<Link[]>(`${BASE_URL}/links`);
  return response.data;
}

// Delete a link by ID
export async function deleteLinkById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/links/${id}`, { withCredentials: true });
}
