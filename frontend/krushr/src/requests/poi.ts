import axios from "axios";
import type { PointOfInterest } from "../types/models";
import type {
  PostPointOfInterestBody,
  PutPointOfInterestBody,
} from "../types/request-bodies";
import {
  CREATE_POINT_OF_INTEREST_ENDPOINT,
  DELETE_POINT_OF_INTEREST_BY_ID_ENDPOINT,
  GET_ALL_POINTS_OF_INTEREST_ENDPOINT,
  GET_POINT_OF_INTEREST_BY_ID_ENDPOINT,
  UPDATE_POINT_OF_INTEREST_BY_ID_ENDPOINT,
} from "./endpoints";

// Get all points of interest
export async function getAllPointsOfInterest(): Promise<PointOfInterest[]> {
  const response = await axios.get<PointOfInterest[]>(
    GET_ALL_POINTS_OF_INTEREST_ENDPOINT
  );
  return response.data;
}

// Get a point of interest by ID
export async function getPointOfInterestById(
  id: number
): Promise<PointOfInterest> {
  const response = await axios.get<PointOfInterest>(
    GET_POINT_OF_INTEREST_BY_ID_ENDPOINT(id)
  );
  return response.data;
}

// Delete a point of interest by ID
export async function deletePointOfInterestById(id: number): Promise<void> {
  await axios.delete(DELETE_POINT_OF_INTEREST_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}

// Update a point of interest by ID
export async function updatePointOfInterestById(
  id: number,
  requestBody: PutPointOfInterestBody
): Promise<PointOfInterest> {
  const response = await axios.put<PointOfInterest>(
    UPDATE_POINT_OF_INTEREST_BY_ID_ENDPOINT(id),
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
    CREATE_POINT_OF_INTEREST_ENDPOINT,
    requestBody,
    { withCredentials: true }
  );
  return response.data;
}
