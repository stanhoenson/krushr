<script lang="ts">
  import L, { LatLng, LatLngTuple, LeafletMouseEvent } from "leaflet";
  import { onMount, afterUpdate, onDestroy } from "svelte";
  import { goudaCoordinates } from "../constants";
  import "leaflet-routing-machine";
  import "leaflet-routing-machine/dist/leaflet-routing-machine.css";
  import "leaflet.fullscreen/Control.FullScreen.js";
  import "leaflet.fullscreen/Control.FullScreen.css";
  import type {
    CoordinatesWithPosition,
    ExtendedMarkerOptions,
  } from "../types/misc";
  import type { PutPointOfInterestBody } from "../types/request-bodies";
  import Alert from "./Alert.svelte";
  let element: any;

  export let disabled: boolean;
  export let allPointsOfInterest: PutPointOfInterestBody[];
  export let position: number;
  export let longitude: number;
  export let latitude: number;
  let map: L.Map;
  let marker: L.Marker;
  let routingControl: L.Routing.Control;
  let error: string = "";

  const initialLatLng: LatLngTuple = [
    Number(goudaCoordinates.latitude.toFixed(6)),
    Number(goudaCoordinates.longitude.toFixed(6)),
  ];

  const initialZoom = 13;

  function onMapClick(e: LeafletMouseEvent) {
    if (disabled) return;
    console.log("updating position", position);
    longitude = Number(e.latlng.lng.toFixed(6));
    latitude = Number(e.latlng.lat.toFixed(6));
    // marker.setLatLng(e.latlng);
  }

  //TODO sometimes pois are undefined, now i just check but shouldnt really happen look into it! PROB rewrite this
  function handlePoisUpdate(map: L.Map, pois: PutPointOfInterestBody[]) {
    let waypoints: L.LatLng[] = [];
    console.log(position);
    let poiIndexesFound: number[] = [];
    map.eachLayer((layer) => {
      if (layer instanceof L.Marker) {
        let options = layer.options as ExtendedMarkerOptions;
        if (
          options.position !== null &&
          !poiIndexesFound.includes(options.position)
        ) {
          if (options.position >= pois.length) {
            layer.remove();
          }
          let poi = pois[options.position];
          if (poi) {
            layer.setLatLng([poi.latitude, poi.longitude]);
            poiIndexesFound.push(options.position);
          }
        }
      }
    });
    for (let index in pois) {
      let poi = pois[index];
      if (!poi) continue;
      waypoints.push(new LatLng(poi.latitude, poi.longitude));
      if (poiIndexesFound.includes(parseInt(index))) {
        continue;
      }
      console.log("creating", index);
      marker = L.marker([poi.latitude, poi.longitude], {
        icon: L.divIcon({
          html: `<div>${parseInt(index) + 1}</div>`,
          className: "map-marker",
        }),
        position: parseInt(index),
      } as ExtendedMarkerOptions).addTo(map);
    }
    if (!routingControl) {
      routingControl = L.Routing.control({
        fitSelectedRoutes: false,
        waypoints,
        router: import.meta.env.PUBLIC_OSRM_URL
          ? L.routing.osrmv1({
              serviceUrl: "https://osrm.hoenson.xyz/route/v1",
              profile: "walking",
            })
          : undefined,
        plan: new L.Routing.Plan([], {
          createMarker: () => {
            return false;
          },
        }),
      }).addTo(map);

      routingControl.hide();
    } else {
      routingControl;
      routingControl.setWaypoints(waypoints);
    }
  }

  onMount(() => {
    console.log("mouunt??");
    map = L.map(element).setView(initialLatLng, initialZoom);
    L.tileLayer("https://tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution:
        '&copy; <a href="http://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    }).addTo(map);
    map.on("click", onMapClick);
    marker = L.marker(initialLatLng, {
      icon: L.divIcon({
        html: `<div>${position + 1}</div>`,
        className: "map-marker",
      }),
      position,
    } as ExtendedMarkerOptions).addTo(map);
    L.control
      .fullscreen({
        position: "topleft", // change the position of the button can be topleft, topright, bottomright or bottomleft, default topleft
        title: "Show me the fullscreen !", // change the title of the button, default Full Screen
        titleCancel: "Exit fullscreen mode", // change the title of the button when fullscreen is on, default Exit Full Screen
        content: undefined, // change the content of the button, can be HTML, default null
        forceSeparateButton: true, // force separate button to detach from zoom buttons, default false
        forcePseudoFullscreen: true, // force use of pseudo full screen even if full screen API is available, default false
        fullscreenElement: false, // Dom element to render in full screen, false by default, fallback to map._container
      })
      .addTo(map);
    try {
      handlePoisUpdate(map, allPointsOfInterest);
      error = "";
    } catch (e: any) {
      error = e;
    }
  });

  afterUpdate(() => {
    try {
      handlePoisUpdate(map, allPointsOfInterest);
      error = "";
    } catch (e: any) {
      error = e;
    }
  });
</script>

<link
  rel="stylesheet"
  href="https://unpkg.com/leaflet@1.9.3/dist/leaflet.css"
  integrity="sha256-kLaT2GOSpHechhsozzB+flnD+zUyjE2LlfWPgU04xyI="
  crossorigin=""
/>
<div>
  <div class="map main" id={`map${position}`} bind:this={element} />

  {#if error}
    <Alert>{error}</Alert>
  {/if}
</div>
