import axios, { AxiosError } from "axios";
import type { User } from "../types/models";
import type { SignInBody, SignUpBody } from "../types/request-bodies";
import { BASE_URL } from "./endpoints";

export async function signUp(signUpBody: SignUpBody): Promise<User> {
  const response = await axios.post<User>(
    `${BASE_URL}/authentication/sign-up`,
    signUpBody
  );
  return response.data;
}

export async function signIn(signInBody: SignInBody): Promise<string> {
  const response = await axios.post<string>(
    `${BASE_URL}/authentication/sign-in`,
    signInBody,
    { withCredentials: true }
  );
  return response.data;
}
