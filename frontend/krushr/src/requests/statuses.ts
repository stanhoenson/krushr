import axios from "axios";
import type { Status } from "../types/models";
import { BASE_URL } from "./endpoints";

// Get all details
export async function getAllStatuses(): Promise<Status[]> {
  const response = await axios.get<Status[]>(`${BASE_URL}/statuses`);
  return response.data;
}

// Delete a detail by ID
export async function deleteStatusById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/statuses/${id}`, { withCredentials: true });
}
