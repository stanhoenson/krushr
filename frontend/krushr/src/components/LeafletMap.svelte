<script lang="ts">
  import L, { LeafletEvent, LeafletMouseEvent } from "leaflet";
  import { onMount } from "svelte";
  let element: any;
  export let longitude: number;
  export let latitude: number;
  let map: L.Map;
  let marker: L.Marker;

  function onMapClick(e: LeafletMouseEvent) {
    longitude = e.latlng.lng;
    latitude = e.latlng.lat;
    marker.setLatLng(e.latlng);
  }

  onMount(() => {
    map = L.map(element).setView([51.505, -0.09], 13);
    L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      maxZoom: 19,
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    }).addTo(map);
    map.on("click", onMapClick);
    marker = L.marker([51.5, -0.09]).addTo(map);
  });
</script>

<link
  rel="stylesheet"
  href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
  integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI="
  crossorigin=""
/>

<div class="map  main" bind:this={element} />
