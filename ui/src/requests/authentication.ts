import axios, { AxiosError } from "axios";
import type { User } from "../types/models";
import type { SignInBody, SignUpBody } from "../types/request-bodies";
import {
  BASE_URL,
  SIGN_IN_ENDPOINT,
  SIGN_OUT_ENDPOINT,
  SIGN_UP_ENDPOINT,
} from "./endpoints";

export async function signUp(signUpBody: SignUpBody): Promise<User> {
  const response = await axios.post<User>(SIGN_UP_ENDPOINT, signUpBody, {
    withCredentials: true,
  });
  return response.data;
}

export async function signIn(signInBody: SignInBody) {
  await axios.post(SIGN_IN_ENDPOINT, signInBody, {
    withCredentials: true,
  });
}

export async function signOut(): Promise<any> {
  const response = await axios.get(SIGN_OUT_ENDPOINT, {
    withCredentials: true,
  });
  return response.data;
}
