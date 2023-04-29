import axios from "axios";
import type { User } from "../types/models";
import type { PostUserBody, PutUserBody } from "../types/request-bodies";
import { BASE_URL } from "./endpoints";

export async function getAllUsers(): Promise<User[]> {
  const response = await axios.get<User[]>(`${BASE_URL}/users`);
  return response.data;
}

export async function createUser(user: PostUserBody): Promise<User> {
  const response = await axios.post<User>(`${BASE_URL}/users`, user);
  return response.data;
}

export async function updateUser(user: PutUserBody): Promise<User> {
  const response = await axios.put<User>(`${BASE_URL}/users`, user);
  return response.data;
}

export async function deleteUser(userId: number): Promise<void> {
  await axios.delete<void>(`${BASE_URL}/users/${userId}`);
}
