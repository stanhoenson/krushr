<script lang="ts">
  import { createImage } from "../requests/images";
  import type { Category } from "../types/models";
  import type {
    PostPointOfInterestBody,
    PostRouteBody,
  } from "../types/request-bodies";
  import CategoryButtonGroup from "./CategoryButtonGroup.svelte";
  import XMark from "./icons/XMark.svelte";
  import LeafletMap from "./LeafletMap.svelte";
  import ImagePlaceholder from "./svg/ImagePlaceholder.svelte";

  let categories: Category[] = [];

  export let viewOnly: boolean;
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
    pointOfInterest.links.push({ url: "", text: "" });
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
    input.multiple = true;
    input.click();
  }

  function handleDeleteImage(index: number) {
    pointOfInterest.imageIds.splice(index, 1);
    pointOfInterest = pointOfInterest;
  }

  async function handleFileUpload(event: Event) {
    const input = event.target as HTMLInputElement;
    const files = input.files;
    if (files) {
      for (let file of files)
        if (file) {
          let image = await createImage(file);
          pointOfInterest.imageIds.push(image.id);
        }

      pointOfInterest = pointOfInterest;
    }
  }

  function handleDeletePointOfInterest() {
    route.pointsOfInterest.splice(position, 1);
    route = route;
  }

  function handlePositionChange(newPosition: number) {
    positionExchange(position, newPosition);
  }

  function handleSupportChange(event: any) {
    if (!pointOfInterest.name && event.target.checked)
      pointOfInterest.name = "support point";
    pointOfInterest = pointOfInterest;
  }
</script>

<section class="card" id={`poi-${position}`}>
  <p class="upper">Point of interest</p>
  <div class="grid coordinates">
    <LeafletMap
      disabled={viewOnly}
      allPointsOfInterest={route.pointsOfInterest}
      {position}
      bind:latitude={pointOfInterest.latitude}
      bind:longitude={pointOfInterest.longitude}
    />
    <div>
      <div class="flex input">
        <label>Longitude</label>
        <input
          disabled={viewOnly}
          required
          bind:value={pointOfInterest.longitude}
          type="text"
          name="longitude"
        />
      </div>
      <div class="flex input name">
        <label>Latitude</label>
        <input
          disabled={viewOnly}
          required
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
            disabled={viewOnly}
            on:click={handlePositionChange.bind(null, position - 1)}
            type="button"
            class={`button block secondary ${position === 0 ? "disabled" : ""}`}
            href="#">Move to {position - 1}</button
          >
          <span>{position}</span>
          <button
            disabled={viewOnly}
            on:click={handlePositionChange.bind(null, position + 1)}
            type="button"
            class="button block secondary"
            href="#">Move to {position + 1}</button
          >
        </div>
        <hr />
        <label>support point</label>
        <input
          type="checkbox"
          bind:checked={pointOfInterest.support}
          on:change={handleSupportChange}
        />
      </div>
    </div>
  </div>
  {#if !pointOfInterest.support}
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
            <div class={`image-with-delete ${i === 0 ? "main" : ""}`}>
              <img
                src={`${
                  import.meta.env.PUBLIC_API_BASE_URL
                }/imagedata/${imageId}`}
              />
              {#if pointOfInterest.imageIds.length > 1 && !viewOnly}
                <div
                  on:click={handleDeleteImage.bind(null, i)}
                  class="icon delete-icon"
                >
                  <XMark />
                </div>
              {/if}
            </div>
          {/each}
        </div>
        <button
          disabled={viewOnly}
          type="button"
          on:click={handleNewImage}
          class="button block secondary"
          href="#">New image</button
        >
      </div>
      <div class="info">
        <div class="flex input name">
          <label>Name</label>
          <input
            disabled={viewOnly}
            required
            bind:value={pointOfInterest.name}
            type="text"
            name="name"
          />
        </div>
        <hr />
        <div class="multiple details">
          <label>Details</label>
          <div class="flex column">
            {#each pointOfInterest.details as detail, i}
              <div class="input-with-delete">
                <textarea
                  bind:value={detail.text}
                  name="details"
                  rows="4"
                  disabled={viewOnly}
                />
                {#if pointOfInterest.details.length > 1 && !viewOnly}
                  <div
                    on:click={handleDeleteDetail.bind(null, i)}
                    class="icon delete-icon"
                  >
                    <XMark />
                  </div>
                {/if}
              </div>
            {/each}
            <button
              disabled={viewOnly}
              type="button"
              on:click={handleNewDetail}
              class="button block secondary"
              href="#">New detail</button
            >
          </div>
        </div>
        <hr />
        <div class="multiple links">
          <label>Links</label>
          <div class="flex input">
            {#each pointOfInterest.links as link, i}
              <div class="card input-with-delete">
                <div class="flex input name">
                  <label>Text</label>
                  <input
                    required
                    bind:value={link.text}
                    type="text"
                    name="text"
                    disabled={viewOnly}
                  />
                </div>
                <div id="url" class="flex input name">
                  <label>URL</label>
                  <input
                    bind:value={link.url}
                    type="text"
                    name="link"
                    required
                    disabled={viewOnly}
                  />
                </div>
                {#if pointOfInterest.links.length > 0 && !viewOnly}
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
              disabled={viewOnly}
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
          disabled={viewOnly}
          selectedCategories={pointOfInterest.categories}
          {handleCategoryToggle}
        />
      </div>
    </div>
  {/if}
  <hr />
  <button
    type="button"
    on:click={handleDeletePointOfInterest}
    disabled={route.pointsOfInterest.length <= 2 || viewOnly}
    class={`button thick block error `}
    href="#">Delete point of interest</button
  >
</section>
