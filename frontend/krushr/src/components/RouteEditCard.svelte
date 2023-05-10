<script lang="ts">
  import {createImage} from "../requests/images";

  import type {Category} from "../types/models";

  import type { PostRouteBody } from "../types/request-bodies";

  export let route: PostRouteBody;

  function handleCategoryToggle(category: Category) {
    let categoryNames = route.categories.map(
      (category) => category.name
    );
    if (!categoryNames.includes(category.name)) {
      route.categories.push(category);
      route = route;
    }
  }

  function handleNewDetail() {
    route.details.push({ text: "" });
    route = route;
  }

  function handleDeleteDetail(index: number) {}

  function handleNewLink() {
    route.links.push({ url: "" });
    route = route;
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
      route.imageIds.push(image.id);
      route = route;
    }
  }

  function handleDeleteImage(index: number) {}

  function handleDeletePointOfInterest() {}
</script>

<section class="card">
  <p class="upper">Route</p>
  <div class="grid">
    <div class="images">
      <div class="grid">
        {#each route.imageIds as imageId, i}
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
        <input type="text" name="name" />
      </div>
      <hr />
      <div class="multiple details">
        <label>Details</label>
        <div class="flex column">
          <textarea name="details" rows="4" />
        </div>
        <div class="grid buttons">
          <a class="button block secondary" href="#">New detail</a>
          <a class="button block error disabled" href="#">Delete last detail</a>
        </div>
      </div>
      <hr />
      <div class="multiple links">
        <label>Links</label>
        <div class="flex input">
          <input type="text" name="link" />
        </div>
        <div class="flex input">
          <input type="text" name="link" />
          <a class="button block secondary" href="#">New link</a>
        </div>
        <a class="button block error" href="#">Delete last link</a>
      </div>
      <hr />
      <p id="categories">Categories</p>
      <div class="flex categories">
        <a class="category" href="#">Architecture</a>
        <a class="category selected" href="#">Nature</a>
        <a class="category" href="#">Culture</a>
        <a class="category" href="#">Dining</a>
        <a class="category" href="#">Shopping</a>
        <a class="category" href="#">History</a>
      </div>
    </div>
  </div>
</section>
