<script lang="ts">
  import L, { LatLngTuple, LeafletMouseEvent } from "leaflet";
  import { onMount, afterUpdate } from "svelte";
  let element: any;
  export let allMarkers: L.Marker[];
  export let longitude: number;
  export let latitude: number;
  let map: L.Map;
  let marker: L.Marker;

  const initialLatLng: LatLngTuple = [52.0164830842629, 4.710044860839845];
  const initialZoom = 13;

  function onMapClick(e: LeafletMouseEvent) {
    longitude = e.latlng.lng;
    latitude = e.latlng.lat;
    marker.setLatLng(e.latlng);
  }

  function handleMarkerUpdates(map: L.Map, markers: L.Marker[]) {
    for (let marker of markers) {
      marker.addTo(map);
    }
  }

  onMount(() => {
    map = L.map(element).setView(initialLatLng, initialZoom);
    L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      maxZoom: 19,
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    }).addTo(map);
    map.on("click", onMapClick);
    marker = L.marker(initialLatLng, {
      icon: L.divIcon({ html: "1", className: "map-marker" }),
    }).addTo(map);
    if (map && allMarkers.length > 0) {
      handleMarkerUpdates(map, allMarkers);
    }
  });

  afterUpdate(() => {
    if (map && allMarkers.length > 0) {
      handleMarkerUpdates(map, allMarkers);
    }
  });
</script>

<link
  rel="stylesheet"
  href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
  integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI="
  crossorigin=""
/>

<div class="map  main" bind:this={element} />
