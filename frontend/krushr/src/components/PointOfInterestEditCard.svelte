<script lang="ts">
  import LeafletMap from "./LeafletMap.svelte";
  export let longitude: number;
  export let latitude: number;
</script>

<section class="card">
  <p class="upper">Point of interest</p>
  <div class="grid coordinates">
    <LeafletMap bind:latitude bind:longitude />
    <div>
      <div class="flex input">
        <label>Longitude</label>
        <input bind:value={longitude} type="text" name="longitude" />
      </div>
      <div class="flex input name">
        <label>Latitude</label>
        <input bind:value={latitude} type="text" name="latitude" />
      </div>
      <hr />
      <a class="button block primary" href="#">Move up</a>
      <a class="button block primary" href="#">Move down</a>
    </div>
  </div>
  <hr />
  <div class="grid">
    <div class="images">
      <div class="grid">
        <img class="main" src="../../_images/landscape-ile-de-france.jpg" />
        <img src="../../_images/hampton-court-green.jpg" />
        <img src="../../_images/charing-cross-bridge.jpg" />
        <img src="../../_images/charing-cross-bridge.jpg" />
      </div>
      <div class="grid">
        <a class="button block secondary" href="#">New image</a>
        <a class="button block error" href="#">Delete last image</a>
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
          <a class="button block secondary" href="#">New link</a>
        </div>
        <a class="button block error disabled" href="#">Delete last link</a>
      </div>
      <hr />
      <p id="categories">Categories</p>
      <div class="flex categories">
        <a class="category" href="#">Architecture</a>
        <a class="category" href="#">Nature</a>
        <a class="category selected" href="#">Culture</a>
        <a class="category" href="#">Dining</a>
        <a class="category selected" href="#">Shopping</a>
        <a class="category" href="#">History</a>
      </div>
    </div>
  </div>
  <hr />
  <a id="save" class="button block error disabled" href="#"
    >Delete point of interest</a
  >
</section>
