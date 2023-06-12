<script lang="ts">
  import { afterUpdate, onDestroy, onMount } from "svelte";
  import { getAllRoutes } from "../requests/routes";
  import type { Route, Status, User } from "../types/models";
  import Fuse from "fuse.js";
  import {
    groupRoutesByStatus,
    groupRoutesByStatusAndUserId,
  } from "../utils/routes";
  import RouteCard from "./RouteCard.svelte";
  import { applicationState } from "../stores/application-state";
  import DelayedInput from "./DelayedInput.svelte";

  // let routes: Route[] = [];
  let error: any = null;
  let user: User | null;
  let filter: string = "";
  let routesFilteredOn: string = "";
  let filtering: boolean = false;

  let unsubscribe = applicationState.subscribe((value) => {
    user = value.authenticatedUser;
  });

  let routes: Route[] = [];
  let fuse = new Fuse(routes, {
    keys: [
      "name",
      "categories.name",
      // "links.text",
      // "links.url",
      // "details.text",
    ],
    shouldSort: true,
  });

  let groupedRoutes: {
    isUser: { status: Status["name"]; routes: Route[] }[];
    notUser: { status: Status["name"]; routes: Route[] }[];
  } = { isUser: [], notUser: [] };

  async function filterRoutes(routes: Route[], fuseInstance: Fuse<Route>) {
    let filteredRoutes = routes;
    if (filter !== "") {
      fuseInstance.setCollection(routes);

      let fuseResults = fuseInstance.search(filter);
      filteredRoutes = fuseResults.map((result) => result.item as Route);
    }
    routesFilteredOn = filter;
    return filteredRoutes;
  }

  onMount(async () => {
    try {
      routes = await getAllRoutes();
      let filteredRoutes = await filterRoutes(routes, fuse);
      groupedRoutes = groupRoutesByStatus(filteredRoutes, user ? user.id : -1);
    } catch (e: any) {
      error = e;
    }
  });
  onDestroy(() => {
    unsubscribe();
  });

  afterUpdate(async () => {
    if (filter != routesFilteredOn) {
      let filteredRoutes = await filterRoutes(routes, fuse);
      groupedRoutes = groupRoutesByStatus(filteredRoutes, user ? user.id : -1);
    }
  });
</script>

<section class="routes">
  {#if user}
    <a class="button block primary" href="/routes/create">Create a new route</a>
  {/if}
  <label>Filter</label>
  <DelayedInput bind:loading={filtering} bind:value={filter} />
  {#if filtering}
    <span>Loading...</span>
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
          <RouteCard {route} />
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
          <RouteCard {route} />
        {/each}
      </div>
    {/if}
  {/each}
</section>
