<script lang="ts">
  import { createImage } from "../requests/images";

  import type { Category } from "../types/models";

  import type { PostRouteBody } from "../types/request-bodies";
  import CategoryButtonGroup from "./CategoryButtonGroup.svelte";
  import XMark from "./icons/XMark.svelte";
  import ImagePlaceholder from "./svg/ImagePlaceholder.svelte";

  export let route: PostRouteBody;

  function handleCategoryToggle(category: Category) {
    let matchedCategoryIndex = route.categories.findIndex(
      (innerCategory) => innerCategory.name === category.name
    );
    if (matchedCategoryIndex === -1) {
      route.categories.push(category);
      route = route;
    } else {
      route.categories.splice(matchedCategoryIndex, 1);
      route = route;
    }
  }

  function handleNewDetail() {
    route.details.push({ text: "" });
    route = route;
  }

  function handleDeleteDetail(index: number) {
    route.details.splice(index, 1);
    route = route;
  }

  function handleNewLink() {
    route.links.push({ url: "" });
    route = route;
  }

  function handleDeleteLink(index: number) {
    route.links.splice(index, 1);
    route = route;
  }

  function handleNewImage() {
    const input = document.createElement("input");
    input.type = "file";
    input.onchange = handleFileUpload;
    input.click();
  }

  function handleDeleteImage(index: number) {
    route.imageIds.splice(index, 1);
    route = route;
  }

  async function handleFileUpload(event: Event) {
    const input = event.target as HTMLInputElement;
    const file = input.files?.[0];
    if (file) {
      let image = await createImage(file);
      route.imageIds.push(image.id);
      route = route;
    }
  }
</script>

<section class="card">
  <p class="upper">Route</p>
  <div class="grid">
    <div class="images">
      <div class="grid">
        {#if route.imageIds.length === 0}
          <div class="main image-placeholder">
            <ImagePlaceholder />
          </div>
        {/if}
        {#each route.imageIds as imageId, i}
          <div class={`image-with-delete ${i === 0 ? "main" : ""}`}>
            <img
              src={`${
                import.meta.env.PUBLIC_API_BASE_URL
              }/imagedata/${imageId}`}
            />
            {#if route.imageIds.length > 1}
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
      <div class="grid">
        <button
          type="button"
          on:click={handleNewImage}
          class="button block secondary"
          href="#">New image</button
        >
      </div>
    </div>
    <div class="info">
      <div class="flex input name">
        <label>Name</label>
        <input bind:value={route.name} type="text" name="name" required />
      </div>
      <hr />
      <div class="multiple details">
        <label>Details</label>
        <div class="flex column">
          {#each route.details as detail, i}
            <div class="input-with-delete">
              <textarea
                bind:value={detail.text}
                name="details"
                rows="4"
                required
              />
              {#if route.details.length > 1}
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
          {#each route.links as link, i}
            <div class="input-with-delete">
              <input bind:value={link.url} type="text" name="link" required />
              {#if route.links.length > 0}
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
        selectedCategories={route.categories}
        {handleCategoryToggle}
      />
    </div>
  </div>
</section>
