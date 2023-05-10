import axios from "axios";
import type { Image } from "../types/models";
import { BASE_URL } from "./endpoints";

// Get an image by ID
export async function getImageById(id: number): Promise<Image> {
  const response = await axios.get<Image>(`${BASE_URL}/images/${id}`);
  return response.data;
}

// Delete an image by ID
export async function deleteImageById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/images/${id}`, { withCredentials: true });
}

// Create an image
export async function createImage(file: File): Promise<Image> {
  const formData = new FormData();
  formData.append("file", file);

  const response = await axios.post(`${BASE_URL}/images`, formData, {
    withCredentials: true,
  });
  return response.data;
}
