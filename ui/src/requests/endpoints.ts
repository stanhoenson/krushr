export const BASE_URL = import.meta.env.PUBLIC_API_BASE_URL;
export const SIGN_UP_ENDPOINT = `${BASE_URL}/authentication/sign-up`;
export const SIGN_IN_ENDPOINT = `${BASE_URL}/authentication/sign-in`;
export const SIGN_OUT_ENDPOINT = `${BASE_URL}/authentication/sign-out`;

export const GET_ALL_CATEGORIES_ENDPOINT = `${BASE_URL}/categories`;
export const DELETE_CATEGORY_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/categories/${id}`;
export const CREATE_CATEGORY_ENDPOINT = `${BASE_URL}/categories`;

export const UPDATE_CATEGORY_ENDPOINT = (id: number) =>
  `${BASE_URL}/categories/${id}`;

export const GET_ALL_DETAILS_ENDPOINT = `${BASE_URL}/details`;

export const DELETE_DETAIL_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/details/${id}`;

export const GET_IMAGEDATA_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/imagedata/${id}`;

export const GET_IMAGE_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/images/${id}`;
export const DELETE_IMAGE_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/images/${id}`;
export const CREATE_IMAGE_ENDPOINT = `${BASE_URL}/images`;

export const GET_ALL_LINKS_ENDPOINT = `${BASE_URL}/links`;
export const DELETE_LINK_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/links/${id}`;

export const GET_ALL_POINTS_OF_INTEREST_ENDPOINT = `${BASE_URL}/points-of-interest`;

export const GET_POINT_OF_INTEREST_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/points-of-interest/${id}`;

export const DELETE_POINT_OF_INTEREST_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/points-of-interest/${id}`;

export const UPDATE_POINT_OF_INTEREST_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/points-of-interest/${id}`;
export const CREATE_POINT_OF_INTEREST_ENDPOINT = `${BASE_URL}/points-of-interest`;

export const GET_ALL_ROUTES_ENDPOINT = `${BASE_URL}/routes`;
export const GET_ROUTE_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/routes/${id}`;
export const CREATE_ROUTE_ENDPOINT = `${BASE_URL}/routes`;
export const UPDATE_ROUTE_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/routes/${id}`;
export const DELETE_ROUTE_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/routes/${id}`;

export const GET_ALL_STATUSES_ENDPOINT = `${BASE_URL}/statuses`;
export const DELETE_STATUS_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/statuses/${id}`;

export const GET_ME_USER_ENDPOINT = `${BASE_URL}/users/me`;
export const GET_ALL_USERS_ENDPOINT = `${BASE_URL}/users`;
export const CREATE_USER_ENDPOINT = `${BASE_URL}/users`;
export const UPDATE_USER_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/users/${id}`;
export const DELETE_USER_BY_ID_ENDPOINT = (id: number) =>
  `${BASE_URL}/users/${id}`;
