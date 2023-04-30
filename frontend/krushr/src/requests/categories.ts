import axios from "axios";
import type { Category } from "../types/models";
import type {
  PostCategoryBody,
  PutCategoryBody,
} from "../types/request-bodies";
import { BASE_URL } from "./endpoints";

// Get all categories
export async function getAllCategories(): Promise<Category[]> {
  const response = await axios.get<Category[]>(`${BASE_URL}/categories`);
  return response.data;
}

// Delete a category by ID
export async function deleteCategoryById(id: number): Promise<void> {
  await axios.delete(`${BASE_URL}/categories/${id}`, { withCredentials: true });
}

// Create a new category
export async function createCategory(
  category: PostCategoryBody
): Promise<Category> {
  const response = await axios.post<Category>(
    `${BASE_URL}/categories`,
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
    `${BASE_URL}/categories/${id}`,
    category,
    { withCredentials: true }
  );
  return response.data;
}
