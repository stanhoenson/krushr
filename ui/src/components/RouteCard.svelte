<script lang="ts">
  import { onMount } from "svelte";
  import type { Route } from "../types/models";
  import { url } from "inspector";

  export let route: Route;
  // let error: any = null;
  let kmFormatter = new Intl.NumberFormat("en-US", {
    style: "unit",
    unit: "kilometer",
    unitDisplay: "short",
    maximumFractionDigits: 1,
  });

  onMount(async () => {});
</script>

<!-- <section class="routes"> -->
<!--   <p style="text-align: left" class="upper soft">Published routes</p> -->
<!--   <div class="grid"> -->
<!--   </div> -->
<!-- </section> -->
<div>
  <div class="card item">
    <img
      src={`${import.meta.env.PUBLIC_API_BASE_URL}/imagedata/${
        route.images[0].id
      }`}
    />
    <h1>{route.name}</h1>
    <p class="soft details">
      {kmFormatter.format(route.distance)} <span class="pipe">| </span>
      {#each route.categories as category, i}
        {#if i !== 0}
          <span class="dot">{" "}&#x2022;</span>
        {/if}
        {category.name}
      {/each}

      {#if route.status.name}
        <span class="pipe">|</span>
        {route.status.name}
      {/if}
    </p>
    {#each route.details as detail}
      <p>
        {detail.text}
      </p>
    {/each}
    {#if route.links.length != 0}
      <div class="route-links">
        {#each route.links as link}
          <a class="block" href={link.url}>
            {link.text}
          </a>
        {/each}
      </div>
    {/if}
    <a class="button block shade" href={`/routes/edit?id=${route.id}`}
      >Go to route</a
    >
  </div>
</div>
