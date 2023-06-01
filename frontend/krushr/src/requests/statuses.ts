import axios from "axios";
import type { Status } from "../types/models";
import {
  BASE_URL,
  DELETE_STATUS_BY_ID_ENDPOINT,
  GET_ALL_STATUSES_ENDPOINT,
} from "./endpoints";

// Get all details
export async function getAllStatuses(): Promise<Status[]> {
  const response = await axios.get<Status[]>(GET_ALL_STATUSES_ENDPOINT);
  return response.data;
}

// Delete a detail by ID
export async function deleteStatusById(id: number): Promise<void> {
  await axios.delete(DELETE_STATUS_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}
