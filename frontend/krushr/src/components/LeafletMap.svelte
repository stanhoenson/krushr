<script lang="ts">
  import L, { LatLngTuple, LeafletMouseEvent } from "leaflet";
  import { onMount, afterUpdate, onDestroy } from "svelte";
  import { goudaCoordinates } from "../constants";
  import type {
    CoordinatesWithPosition,
    ExtendedMarkerOptions,
  } from "../types/misc";
  let element: any;
  export let allPointsOfInterestLatLngs: LatLngTuple[] = [];
  export let position: number;
  export let longitude: number;
  export let latitude: number;
  let map: L.Map;
  let marker: L.Marker;

  const initialLatLng: LatLngTuple = [
    goudaCoordinates.latitude,
    goudaCoordinates.longitude,
  ];
  const initialZoom = 13;

  function onMapClick(e: LeafletMouseEvent) {
    longitude = e.latlng.lng;
    latitude = e.latlng.lat;
    marker.setLatLng(e.latlng);
  }

  function handlePoisUpdate(map: L.Map, pois: LatLngTuple[]) {
    // for (let poi of pois) {
    //   marker = L.marker(initialLatLng, {
    //     icon: L.divIcon({
    //       html: `<div>${position}</div>`,
    //       className: "map-marker",
    //     }),
    //     position,
    //   } as ExtendedMarkerOptions).addTo(map);
    // }
    map.eachLayer((layer) => {
      if (layer instanceof L.Marker) {
        let options = layer.options as ExtendedMarkerOptions;
        if (options.position) {
          let poi = pois[options.position];
          layer.setLatLng(poi);
        }
      }
    });
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
      icon: L.divIcon({
        html: position.toString(),
        className: "map-marker",
      }),
      position,
    } as ExtendedMarkerOptions).addTo(map);
  });

  afterUpdate(() => {
    handlePoisUpdate(map, allPointsOfInterestLatLngs);
  });
</script>

<link
  rel="stylesheet"
  href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
  integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI="
  crossorigin=""
/>

<div class="map  main" id={`map${position}`} bind:this={element} />
