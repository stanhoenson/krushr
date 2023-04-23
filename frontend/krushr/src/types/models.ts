export interface Route {
  id: number;
  name: string;
  images: Image[];
  details: Detail[];
  links: Link[];
  categories: Category[];
  status: Status;
  statusId: number;
  pointsOfInterest: PointOfInterest[];
  distance: number;
  user: User;
  userId: number;
}

export interface Image {
  id: number;
  path: string;
  routes: Route[];
  pointsOfInterest: PointOfInterest[];
}

export interface Detail {
  id: number;
  text: string;
  routes: Route[];
  pointsOfInterest: PointOfInterest[];
}

export interface Link {
  id: number;
  url: string;
  routes: Route[];
  pointsOfInterest: PointOfInterest[];
}

export interface Category {
  id: number;
  name: string;
  position: number;
  routes: Route[];
  pointsOfInterest: PointOfInterest[];
}

export interface Status {
  id: number;
  name: string;
}

export interface PointOfInterest {
  id: number;
  name: string;
  longitude: number;
  latitude: number;
  images: Image[];
  details: Detail[];
  links: Link[];
  categories: Category[];
  routes: Route[];
  user: User;
  userId: number;
}

export interface User {
  id: number;
  email: string;
  password: string;
  role: Role;
  roleId: number;
}

export interface Role {
  id: number;
  name: string;
}

// interface RoutesPointsOfInterest {
//   routeId: number;
//   pointOfInterestId: number;
//   position: number;
// }
