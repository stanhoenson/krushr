<script lang="ts">
  import { onDestroy, onMount } from "svelte";
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

  let groupedRoutes: {
    isUser: boolean;
    routes: { status: Status["name"]; routes: Route[] }[];
  }[] = [];

  onMount(async () => {
    try {
      let routes = await getAllRoutes();
      groupedRoutes = groupRoutesByStatus(routes, user ? user.id : -1);
      console.log(groupedRoutes);
    } catch (e: any) {
      error = e.response.data.error;
    }
  });
  onDestroy(() => {
    unsubscribe();
  });
</script>

<section class="routes">
  {#if user}
    <a class="button block primary" href="/routes/create">Create a new route</a>
  {/if}
  {#each groupedRoutes as routeGroup}
    {#each routeGroup.routes as innerRouteGroup}
    <hr class="soft" />
      <p class="status">
        <span>
          {innerRouteGroup.status} routes
        </span>
        {#if routeGroup.isUser && user}
          by yourself
        {:else if user}
          by others
        {/if}
      </p>
      <div class="grid">
        {#each innerRouteGroup.routes as route}
          <RouteCard {route} />
        {/each}
      </div>
    {/each}
  {/each}
</section>
