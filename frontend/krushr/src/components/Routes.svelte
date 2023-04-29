<script lang="ts">
  import { onMount } from "svelte";
  import { getAllRoutes } from "../requests/routes";
  import type { Route } from "../types/models";

  let routes: Route[] = [];
  let error: any = null;
  let kmFormatter = new Intl.NumberFormat("en-US", {
    style: "unit",
    unit: "kilometer",
    unitDisplay: "long",
    maximumFractionDigits: 2,
  });

  onMount(async () => {
    try {
      routes = await getAllRoutes();
    } catch (e: any) {
      error = e.response.data.error;
    }
  });
</script>

<h1>All Routes</h1>
<div class="routes">
  {#each routes as route}
    <div>
      <div class="route">
        <img
          src={`${import.meta.env.PUBLIC_API_BASE_URL}/imagedata/${
            route.images[0].id
          }`}
        />
        <h1>{route.name}</h1>
        <p class="details">
          {kmFormatter.format(route.distance)} <span>| </span>
          {#each route.categories as category, i}
            {#if i !== 0}
              <span class="dot">&#x2022;</span>
            {/if}
            {category.name}
          {/each}

          <span>|</span>
          {route.status.name}
        </p>
        <p class="description">
          {#each route.details as detail}
            {detail.text}.{" "}
          {/each}
        </p>
        <a href="/">Go to route</a>
      </div>
    </div>
  {/each}
</div>
