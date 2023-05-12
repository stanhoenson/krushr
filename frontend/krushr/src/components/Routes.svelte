<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from "svelte";
  import { getAllRoutes } from "../requests/routes";
  import { getMeUser } from "../requests/users";
  import { authenticatedUser } from "../stores/user";
  import type { Route, Status, User } from "../types/models";
  import {
    groupRoutesByStatus,
    groupRoutesByStatusAndUserId,
  } from "../utils/routes";
  import RouteCard from "./RouteCard.svelte";

  // let routes: Route[] = [];
  let error: any = null;
  let user: User | null;

  let unsubscribe = authenticatedUser.subscribe((value) => {
    user = value;
  });

  let routes: Route[];

  let groupedRoutes: {
    isUser: { status: Status["name"]; routes: Route[] }[];
    notUser: { status: Status["name"]; routes: Route[] }[];
  } = { isUser: [], notUser: [] };

  onMount(async () => {
    try {
      routes = await getAllRoutes();
      groupedRoutes = groupRoutesByStatus(routes, user ? user.id : -1);
      console.log(groupedRoutes);
    } catch (e: any) {
      error = e.response.data.error;
    }
  });
  onDestroy(() => {
    unsubscribe();
  });

  afterUpdate(async () => {
    groupedRoutes = groupRoutesByStatus(routes, user ? user.id : -1);
  });
</script>

<section class="routes">
  {#if user}
    <a class="button block primary" href="/routes/create">Create a new route</a>
  {/if}
  {#each groupedRoutes.isUser as innerRouteGroup}
    {#if innerRouteGroup.routes.length > 0}
      <hr class="soft" />
      <p class="status">
        <span>
          {innerRouteGroup.status} routes
        </span>
        by yourself
      </p>
      <div class="grid">
        {#each innerRouteGroup.routes as route}
          <RouteCard loggedIn={user !== null} {route} />
        {/each}
      </div>
    {/if}
  {/each}
  {#each groupedRoutes.notUser as innerRouteGroup}
    {#if innerRouteGroup.routes.length > 0}
      <hr class="soft" />
      <p class="status">
        <span>
          {innerRouteGroup.status} routes
        </span>
        by others
      </p>
      <div class="grid">
        {#each innerRouteGroup.routes as route}
          <RouteCard loggedIn={user !== null} {route} />
        {/each}
      </div>
    {/if}
  {/each}
</section>
