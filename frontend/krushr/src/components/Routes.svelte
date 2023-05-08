<script lang="ts">
  import { onMount } from "svelte";
  import { getAllRoutes } from "../requests/routes";
  import { getMeUser } from "../requests/users";
  import type { Route, Status, User } from "../types/models";
  import {
    groupRoutesByStatus,
    groupRoutesByStatusAndUserId,
  } from "../utils/routes";
  import RouteCard from "./RouteCard.svelte";

  // let routes: Route[] = [];
  let error: any = null;
  let user: User;
  let groupedRoutes: {
    isUser: boolean;
    routes: { status: Status["name"]; routes: Route[] }[];
  }[] = [];

  onMount(async () => {
    try {
      user = await getMeUser();
      console.log(user);
    } catch (e) {}
    try {
      let routes = await getAllRoutes();
      groupedRoutes = groupRoutesByStatus(routes, user ? user.id : -1);
      console.log(groupedRoutes);
    } catch (e: any) {
      error = e.response.data.error;
    }
  });
</script>

<section class="routes">
  {#if user}
    <a class="button block primary" href="/routes/create">Create a new route</a>
  {/if}
	<hr class="soft">
  {#each groupedRoutes as routeGroup}
    {#each routeGroup.routes as innerRouteGroup}
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
