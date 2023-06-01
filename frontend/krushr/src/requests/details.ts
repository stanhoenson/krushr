import axios from "axios";
import type { Detail } from "../types/models";
import {
  DELETE_CATEGORY_BY_ID_ENDPOINT,
  GET_ALL_CATEGORIES_ENDPOINT,
} from "./endpoints";

// Get all details
export async function getAllDetails(): Promise<Detail[]> {
  const response = await axios.get<Detail[]>(GET_ALL_CATEGORIES_ENDPOINT);
  return response.data;
}

// Delete a detail by ID
export async function deleteDetailById(id: number): Promise<void> {
  await axios.delete(DELETE_CATEGORY_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}
