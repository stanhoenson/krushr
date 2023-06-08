import axios from "axios";
import type { Category } from "../types/models";
import type {
  PostCategoryBody,
  PutCategoryBody,
} from "../types/request-bodies";
import {
  CREATE_CATEGORY_ENDPOINT,
  DELETE_CATEGORY_BY_ID_ENDPOINT,
  GET_ALL_CATEGORIES_ENDPOINT,
  UPDATE_CATEGORY_ENDPOINT,
} from "./endpoints";

// Get all categories
export async function getAllCategories(): Promise<Category[]> {
  const response = await axios.get<Category[]>(GET_ALL_CATEGORIES_ENDPOINT);
  return response.data;
}

// Delete a category by ID
export async function deleteCategoryById(id: number): Promise<void> {
  await axios.delete(DELETE_CATEGORY_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}

// Create a new category
export async function createCategory(
  category: PostCategoryBody
): Promise<Category> {
  const response = await axios.post<Category>(
    CREATE_CATEGORY_ENDPOINT,
    category,
    { withCredentials: true }
  );
  return response.data;
}

// Update an existing category
export async function updateCategory(
  id: number,
  category: PutCategoryBody
): Promise<Category> {
  const response = await axios.put<Category>(
    UPDATE_CATEGORY_ENDPOINT(id),
    category,
    { withCredentials: true }
  );
  return response.data;
}
