import type { Route, Status } from "../types/models";

export function groupRoutesByStatusAndUserId(
  routes: Route[],
  userId: number
): { status: Status["name"]; isUser: boolean; routes: Route[] }[] {
  const groupedRoutes: {
    [key: string]: { isUser: Route[]; notUser: Route[] };
  } = {};

  routes.forEach((route) => {
    const statusName = route.status.name;
    const isUser = route.userId === userId;

    if (!groupedRoutes[statusName]) {
      groupedRoutes[statusName] = { isUser: [], notUser: [] };
    }

    if (isUser) {
      groupedRoutes[statusName].isUser.push(route);
    } else {
      groupedRoutes[statusName].notUser.push(route);
    }
  });

  return Object.entries(groupedRoutes).map(([status, routes]) => ({
    status,
    isUser: !!routes.isUser.length,
    routes: [...routes.isUser, ...routes.notUser],
  }));
}

export function groupRoutesByStatus(
  routes: Route[],
  userId: number
): {
  isUser: boolean;
  routes: { status: Status["name"]; routes: Route[] }[];
}[] {
  const groupedRoutes: {
    [key: string]: { isUser: Route[]; notUser: Route[] };
  } = {};

  routes.forEach((route) => {
    const statusName = route.status.name;
    const isUser = route.user.id === userId;

    if (!groupedRoutes[statusName]) {
      groupedRoutes[statusName] = { isUser: [], notUser: [] };
    }

    if (isUser) {
      groupedRoutes[statusName].isUser.push(route);
    } else {
      groupedRoutes[statusName].notUser.push(route);
    }
  });

  return Object.entries(groupedRoutes).map(([status, routes]) => ({
    isUser: !!routes.isUser.length,
    routes: [
      {
        status,
        routes: [...routes.isUser, ...routes.notUser],
      },
    ],
  }));
}
