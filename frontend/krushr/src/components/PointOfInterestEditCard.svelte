<script lang="ts">
  import type { LatLng, LatLngTuple } from "leaflet";
  import { createImage } from "../requests/images";
  import type {
    Category,
    Detail,
    Image,
    PointOfInterest,
  } from "../types/models";
  import type { PostPointOfInterestBody } from "../types/request-bodies";
  import XMark from "./icons/XMark.svelte";
  import LeafletMap from "./LeafletMap.svelte";

  let categories: Category[] = [];

  export let positionExchange: (
    sourcePoiIndex: number,
    targetPoiIndex: number
  ) => any;
  export let position: number;
  export let pointOfInterest: PostPointOfInterestBody;
  export let allPointsOfInterestLatLngs: LatLngTuple[];

  function handleCategoryToggle(category: Category) {
    let categoryNames = pointOfInterest.categories.map(
      (category) => category.name
    );
    if (!categoryNames.includes(category.name)) {
      pointOfInterest.categories.push(category);
      pointOfInterest = pointOfInterest;
    }
  }

  function handleNewDetail() {
    pointOfInterest.details.push({ text: "" });
    pointOfInterest = pointOfInterest;
  }

  function handleDeleteDetail(index: number) {}

  function handleNewLink() {
    pointOfInterest.links.push({ url: "" });
    pointOfInterest = pointOfInterest;
  }

  function handleDeleteLink(index: number) {}

  function handleNewImage() {
    const input = document.createElement("input");
    input.type = "file";
    input.onchange = handleFileUpload;
    input.click();
  }

  async function handleFileUpload(event: Event) {
    const input = event.target as HTMLInputElement;
    const file = input.files?.[0];
    if (file) {
      let image = await createImage(file);
      pointOfInterest.imageIds.push(image.id);
      pointOfInterest = pointOfInterest;
    }
  }

  function handleDeleteImage(index: number) {}

  function handleDeletePointOfInterest() {}

  function handlePositionChange(newPosition: number) {
    positionExchange(position, newPosition);
  }
</script>

<section class="card" id={`poi-${position}`}>
  <p class="upper">Point of interest</p>
  <div class="grid coordinates">
    <LeafletMap
      {allPointsOfInterestLatLngs}
      {position}
      bind:latitude={pointOfInterest.latitude}
      bind:longitude={pointOfInterest.longitude}
    />
    <div>
      <div class="flex input">
        <label>Longitude</label>
        <input
          bind:value={pointOfInterest.longitude}
          type="text"
          name="longitude"
        />
      </div>
      <div class="flex input name">
        <label>Latitude</label>
        <input
          bind:value={pointOfInterest.latitude}
          type="text"
          name="latitude"
        />
      </div>
      <hr />
      <div class="position">
        <label>Position</label>
        <div class="three">
          <button
            on:click={handlePositionChange.bind(null, position - 1)}
            type="button"
            class={`button block secondary ${position === 0 ? "disabled" : ""}`}
            href="#">Move to {position - 1}</button
          >
          <span>{position}</span>
          <button
            on:click={handlePositionChange.bind(null, position + 1)}
            type="button"
            class="button block secondary"
            href="#">Move to {position + 1}</button
          >
        </div>
      </div>
    </div>
  </div>
  <hr />
  <div class="grid">
    <div class="images">
      <div class="grid">
        {#each pointOfInterest.imageIds as imageId, i}
          {#if i === 0}
            <img
              class="main"
              src={`${
                import.meta.env.PUBLIC_API_BASE_URL
              }/imagedata/${imageId}`}
            />
          {:else}
            <img
              src={`${
                import.meta.env.PUBLIC_API_BASE_URL
              }/imagedata/${imageId}`}
            />
          {/if}
        {/each}
      </div>
      <div class="grid">
        <button
          type="button"
          on:click={handleNewImage}
          class="button block secondary"
          href="#">New image</button
        >
        <button type="button" class="button block error" href="#"
          >Delete last image</button
        >
      </div>
    </div>
    <div class="info">
      <div class="flex input name">
        <label>Name</label>
        <input bind:value={pointOfInterest.name} type="text" name="name" />
      </div>
      <hr />
      <div class="multiple details">
        <label>Details</label>
        <div class="flex column">
          {#each pointOfInterest.details as detail}
            <textarea bind:value={detail.text} name="details" rows="4" />
            <div class="delete-icon"><XMark /></div>
          {/each}
        </div>
        <div class="grid buttons">
          <button
            type="button"
            on:click={handleNewDetail}
            class="button block secondary"
            href="#">New detail</button
          >
          <button type="button" class="button block error disabled" href="#"
            >Delete last detail</button
          >
        </div>
      </div>
      <hr />
      <div class="multiple links">
        <label>Links</label>
        <div class="flex input">
          {#each pointOfInterest.links as link}
            <input bind:value={link.url} type="text" name="link" />
          {/each}
          <button
            type="button"
            on:click={handleNewLink}
            class="button block secondary"
            href="#">New link</button
          >
        </div>
        <button type="button" class="button block error disabled" href="#"
          >Delete last link</button
        >
      </div>
      <hr />
      <p id="categories">Categories</p>
      <div class="flex categories">
        {#each categories as category}
          <button
            type="button"
            on:click={handleCategoryToggle.bind(null, category)}
            class="category"
            href="#">{category.name}</button
          >
        {/each}
      </div>
    </div>
  </div>
  <hr />
  <button
    type="button"
    on:click={handleDeletePointOfInterest}
    class="button thick block error disabled"
    href="#">Delete point of interest</button
  >
</section>
