import axios from "axios";
import type { User } from "../types/models";
import type { PostUserBody, PutUserBody } from "../types/request-bodies";
import {
  BASE_URL,
  CREATE_USER_ENDPOINT,
  DELETE_USER_BY_ID_ENDPOINT,
  GET_ALL_USERS_ENDPOINT,
  GET_ME_USER_ENDPOINT,
  UPDATE_USER_BY_ID_ENDPOINT,
} from "./endpoints";

export async function getMeUser(): Promise<User> {
  const response = await axios.get<User>(GET_ME_USER_ENDPOINT, {
    withCredentials: true,
  });
  return response.data;
}
export async function getAllUsers(): Promise<User[]> {
  const response = await axios.get<User[]>(GET_ALL_USERS_ENDPOINT, {
    withCredentials: true,
  });
  return response.data;
}

export async function createUser(user: PostUserBody): Promise<User> {
  const response = await axios.post<User>(CREATE_USER_ENDPOINT, user, {
    withCredentials: true,
  });
  return response.data;
}

export async function updateUser(id: number, user: PutUserBody): Promise<User> {
  const response = await axios.put<User>(UPDATE_USER_BY_ID_ENDPOINT(id), user, {
    withCredentials: true,
  });
  return response.data;
}

export async function deleteUser(id: number): Promise<void> {
  await axios.delete<number>(DELETE_USER_BY_ID_ENDPOINT(id), {
    withCredentials: true,
  });
}
