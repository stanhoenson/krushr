<script lang="ts">
  import L, { LatLngTuple, LeafletMouseEvent } from "leaflet";
  import { onMount, afterUpdate, onDestroy } from "svelte";
  import { goudaCoordinates } from "../constants";
  import type {
    CoordinatesWithPosition,
    ExtendedMarkerOptions,
  } from "../types/misc";
  import type { PutPointOfInterestBody } from "../types/request-bodies";
  let element: any;
  export let allPointsOfInterest: PutPointOfInterestBody[];
  export let position: number;
  export let longitude: number;
  export let latitude: number;
  let map: L.Map;
  let marker: L.Marker;

  const initialLatLng: LatLngTuple = [
    Number(goudaCoordinates.latitude.toFixed(6)),
    Number(goudaCoordinates.longitude.toFixed(6)),
  ];
  console.log(initialLatLng);

  const initialZoom = 13;

  function onMapClick(e: LeafletMouseEvent) {
    console.log("updating position", position);
    longitude = Number(e.latlng.lng.toFixed(6));
    latitude = Number(e.latlng.lat.toFixed(6));
    // marker.setLatLng(e.latlng);
  }

  function handlePoisUpdate(map: L.Map, pois: PutPointOfInterestBody[]) {
    let poiIndexesFound: number[] = [];
    map.eachLayer((layer) => {
      if (layer instanceof L.Marker) {
        let options = layer.options as ExtendedMarkerOptions;
        console.log(options);
        if (
          options.position !== null &&
          !poiIndexesFound.includes(options.position)
        ) {
          let poi = pois[options.position];
          console.log(
            options.position,
            "position found",
            poi,
            " in map",
            position
          );
          layer.setLatLng([poi.latitude, poi.longitude]);
          poiIndexesFound.push(options.position);
        }
      }
    });
    pois
      .filter((value, index) => !poiIndexesFound.includes(index))
      .forEach((value, index) => {
        marker = L.marker([value.latitude, value.longitude], {
          icon: L.divIcon({
            html: `<div>${index}</div>`,
            className: "map-marker",
          }),
          position,
        } as ExtendedMarkerOptions).addTo(map);
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
        html: `<div>${position}</div>`,
        className: "map-marker",
      }),
      position,
    } as ExtendedMarkerOptions).addTo(map);
  });

  afterUpdate(() => {
    handlePoisUpdate(map, allPointsOfInterest);
  });
</script>

<link
  rel="stylesheet"
  href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
  integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI="
  crossorigin=""
/>

<div class="map  main" id={`map${position}`} bind:this={element} />
