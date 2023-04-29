import axios from "axios";
import type { Detail } from "../types/models";
import { BASE_URL } from "./endpoints";

// Get all details
export async function getAllDetails(): Promise<Detail[]> {
  const response = await axios.get<Detail[]>(`${BASE_URL}/details`);
  return response.data;
}

// Delete a detail by ID
export async function deleteDetailById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/details/${id}`);
}
