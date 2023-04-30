import axios from "axios";
import type { PointOfInterest } from "../types/models";
import type {
  PostPointOfInterestBody,
  PutPointOfInterestBody,
} from "../types/request-bodies";
import { BASE_URL } from "./endpoints";

// Get all points of interest
export async function getAllPointsOfInterest(): Promise<PointOfInterest[]> {
  const response = await axios.get<PointOfInterest[]>(
    `${BASE_URL}/points-of-interest`
  );
  return response.data;
}

// Get a point of interest by ID
export async function getPointOfInterestById(
  id: string
): Promise<PointOfInterest> {
  const response = await axios.get<PointOfInterest>(
    `${BASE_URL}/points-of-interest/${id}`
  );
  return response.data;
}

// Delete a point of interest by ID
export async function deletePointOfInterestById(id: string): Promise<void> {
  await axios.delete(`${BASE_URL}/points-of-interest/${id}`, {
    withCredentials: true,
  });
}

// Update a point of interest by ID
export async function updatePointOfInterestById(
  id: string,
  requestBody: PutPointOfInterestBody
): Promise<PointOfInterest> {
  const response = await axios.put<PointOfInterest>(
    `${BASE_URL}/points-of-interest/${id}`,
    requestBody,
    { withCredentials: true }
  );
  return response.data;
}

// Create a new point of interest
export async function createPointOfInterest(
  requestBody: PostPointOfInterestBody
): Promise<PointOfInterest> {
  const response = await axios.post<PointOfInterest>(
    `${BASE_URL}/points-of-interest`,
    requestBody,
    { withCredentials: true }
  );
  return response.data;
}
