<script lang="ts">
  import type { LatLng, LatLngTuple } from "leaflet";
  import { createImage } from "../requests/images";
  import type {
    Category,
    Detail,
    Image,
    PointOfInterest,
    Route,
  } from "../types/models";
  import type {
    PostPointOfInterestBody,
    PostRouteBody,
  } from "../types/request-bodies";
  import CategoryButtonGroup from "./CategoryButtonGroup.svelte";
  import XMark from "./icons/XMark.svelte";
  import LeafletMap from "./LeafletMap.svelte";
  import ImagePlaceholder from "./svg/ImagePlaceholder.svelte";

  let categories: Category[] = [];

  export let positionExchange: (
    sourcePoiIndex: number,
    targetPoiIndex: number
  ) => any;
  export let position: number;
  export let pointOfInterest: PostPointOfInterestBody;
  export let route: PostRouteBody;

  function handleCategoryToggle(category: Category) {
    let matchedCategoryIndex = pointOfInterest.categories.findIndex(
      (innerCategory) => innerCategory.name === category.name
    );
    if (matchedCategoryIndex === -1) {
      pointOfInterest.categories.push(category);
      pointOfInterest = pointOfInterest;
    } else {
      pointOfInterest.categories.splice(matchedCategoryIndex, 1);
      pointOfInterest = pointOfInterest;
    }
  }

  function handleNewDetail() {
    pointOfInterest.details.push({ text: "" });
    pointOfInterest = pointOfInterest;
  }

  function handleDeleteDetail(index: number) {
    pointOfInterest.details.splice(index, 1);
    pointOfInterest = pointOfInterest;
  }

  function handleNewLink() {
    pointOfInterest.links.push({ url: "" });
    pointOfInterest = pointOfInterest;
  }

  function handleDeleteLink(index: number) {
    pointOfInterest.links.splice(index, 1);
    pointOfInterest = pointOfInterest;
  }

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

  function handleDeleteImage(index: number) {
    pointOfInterest.imageIds.splice(index, 1);
    pointOfInterest = pointOfInterest;
  }

  function handleDeletePointOfInterest() {
    route.pointsOfInterest.splice(position, 1);
    route = route;
  }

  function handlePositionChange(newPosition: number) {
    positionExchange(position, newPosition);
  }
</script>

<section class="card" id={`poi-${position}`}>
  <p class="upper">Point of interest</p>
  <div class="grid coordinates">
    <LeafletMap
      allPointsOfInterest={route.pointsOfInterest}
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
        {#if pointOfInterest.imageIds.length === 0}
          <div class="main image-placeholder">
            <ImagePlaceholder />
          </div>
        {/if}
        {#each pointOfInterest.imageIds as imageId, i}
          <div>
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
          </div>
        {/each}
      </div>
      <div class="grid">
        <button
          type="button"
          on:click={handleNewImage}
          class="button block secondary"
          href="#">New image</button
        >
        <!-- <button type="button" class="button block error" href="#" -->
        <!--   >Delete last image</button -->
        <!-- > -->
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
          {#each pointOfInterest.details as detail, i}
            <div class="input-with-delete">
              <textarea bind:value={detail.text} name="details" rows="4" />
              {#if i !== 0}
                <div
                  on:click={handleDeleteDetail.bind(null, i)}
                  class="icon delete-icon"
                >
                  <XMark />
                </div>
              {/if}
            </div>
          {/each}
        </div>
        <div class="grid buttons">
          <button
            type="button"
            on:click={handleNewDetail}
            class="button block secondary"
            href="#">New detail</button
          >
          <!-- <button type="button" class="button block error disabled" href="#" -->
          <!--   >Delete last detail</button -->
          <!-- > -->
        </div>
      </div>
      <hr />
      <div class="multiple links">
        <label>Links</label>
        <div class="flex input">
          {#each pointOfInterest.links as link, i}
            <div class="input-with-delete">
              <input bind:value={link.url} type="text" name="link" />
              {#if i !== 0}
                <div
                  on:click={handleDeleteLink.bind(null, i)}
                  class="icon delete-icon"
                >
                  <XMark />
                </div>
              {/if}
            </div>
          {/each}
          <button
            type="button"
            on:click={handleNewLink}
            class="button block secondary"
            href="#">New link</button
          >
        </div>
        <!-- <button type="button" class="button block error disabled" href="#" -->
        <!--   >Delete last link</button -->
        <!-- > -->
      </div>
      <hr />
      <p id="categories">Categories</p>
      <CategoryButtonGroup
        selectedCategories={pointOfInterest.categories}
        {handleCategoryToggle}
      />
    </div>
  </div>
  <hr />
  <button
    type="button"
    on:click={handleDeletePointOfInterest}
    disabled={route.pointsOfInterest.length <= 2}
    class={`button thick block error ${
      route.pointsOfInterest.length <= 2 ? "disabled" : ""
    }`}
    href="#">Delete point of interest</button
  >
</section>
