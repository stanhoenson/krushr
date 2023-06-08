import axios from "axios";
import type { Image } from "../types/models";
import {
  CREATE_IMAGE_ENDPOINT,
  DELETE_IMAGE_BY_ID_ENDPOINT,
  GET_IMAGE_BY_ID_ENDPOINT,
} from "./endpoints";

// Get an image by ID
export async function getImageById(id: number): Promise<Image> {
  const response = await axios.get<Image>(GET_IMAGE_BY_ID_ENDPOINT(id));
  return response.data;
}

// Delete an image by ID
export async function deleteImageById(id: number): Promise<void> {
  await axios.delete(DELETE_IMAGE_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}

// Create an image
export async function createImage(file: File): Promise<Image> {
  const formData = new FormData();
  formData.append("file", file);

  const response = await axios.post(CREATE_IMAGE_ENDPOINT, formData, {
    withCredentials: true,
  });
  return response.data;
}
