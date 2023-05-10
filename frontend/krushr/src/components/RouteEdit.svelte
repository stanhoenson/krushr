<script lang="ts">
  import type { LatLngTuple } from "leaflet";

  import { afterUpdate, onDestroy, onMount } from "svelte";
  import { goudaCoordinates } from "../constants";
  import { getAllStatuses } from "../requests/statuses";
  import { getMeUser } from "../requests/users";
  import { authenticatedUser } from "../stores/user";
  import type { PointOfInterest, Route, Status, User } from "../types/models";
  import Alert from "./Alert.svelte";
  import PointOfInterestEditCard from "./PointOfInterestEditCard.svelte";
  import RouteEditCard from "./RouteEditCard.svelte";
  import RouteEditExplanation from "./RouteEditExplanation.svelte";
  import StatusSelect from "./StatusSelect.svelte";

  let route: Route;

  let statuses: Status[];

  //values
  let name: string = "";
  let status: Status;

  let error: string | null = null;
  let user: User;
  let allPointsOfInterestLatLngs: LatLngTuple[] = [];

  // $: allPointsOfInterestLatLngs = route.pointsOfInterest.map(
  //   (pointOfInterest) => {
  //     return [
  //       pointOfInterest.latitude,
  //       pointOfInterest.longitude,
  //     ] as LatLngTuple;
  //   }
  // );

  let queryParams = new URLSearchParams(window.location.search);

  let id = queryParams.get("id");
  if (!id) {
    //create it is then
  }
  function handleNewPointOfInterest() {
    const newPointOfInterest: PointOfInterest = {
      id: 0,
      name: "",
      longitude: goudaCoordinates.longitude,
      latitude: goudaCoordinates.latitude,
      images: [],
      details: [],
      links: [],
      categories: [],
      routes: [],
      user: user,
      userId: user.id,
    };
    route.pointsOfInterest.push(newPointOfInterest);
    route = route;
  }

  onMount(async () => {
    user = await getMeUser();
    statuses = await getAllStatuses();

    route = {
      id: 1,
      name: "My Awesome Route",
      images: [],
      details: [],
      links: [],
      categories: [],
      status: statuses[0],
      statusId: statuses[0].id,
      pointsOfInterest: [],
      distance: 0,
      user: user,
      userId: user.id,
    };

    // getRouteById(queryParams);
  });

  afterUpdate(async () => {
    allPointsOfInterestLatLngs = route.pointsOfInterest.map(
      (pointOfInterest) => {
        return [
          pointOfInterest.latitude,
          pointOfInterest.longitude,
        ] as LatLngTuple;
      }
    );
  });
</script>

<div class="edit">
  {#if user && route}
    <RouteEditExplanation />
    <form>
      <RouteEditCard />
      <hr  class="soft"/>
      {#each route.pointsOfInterest as pointOfInterest, i}
        <PointOfInterestEditCard
          {allPointsOfInterestLatLngs}
          position={i}
          bind:pointOfInterest
        />
      {/each}
      <button
        on:click={handleNewPointOfInterest}
        type="button"
        class="button block secondary thick"
        href="#">New point of interest</button
      >
      <hr class="soft" />
      <div class="fixed">
        <StatusSelect />
        <button id="save" class="button thick block primary" href="#"
          >Save</button
        >
      </div>
    </form>
  {:else}
    <Alert type="error">Unauthorized</Alert>
  {/if}
</div>
