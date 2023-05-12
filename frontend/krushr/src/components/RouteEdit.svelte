<script lang="ts">
  import type { LatLngTuple } from "leaflet";

  import { afterUpdate, onDestroy, onMount } from "svelte";
  import { goudaCoordinates } from "../constants";
  import {
    createRoute,
    deleteRouteById,
    getRouteById,
    updateRoute,
  } from "../requests/routes";
  import { getAllStatuses } from "../requests/statuses";
  import { getMeUser } from "../requests/users";
  import { authenticatedUser } from "../stores/user";
  import type { PointOfInterest, Route, Status, User } from "../types/models";
  import type {
    PostPointOfInterestBody,
    PostRouteBody,
  } from "../types/request-bodies";
  import Alert from "./Alert.svelte";
  import PointOfInterestEditCard from "./PointOfInterestEditCard.svelte";
  import RouteEditCard from "./RouteEditCard.svelte";
  import RouteEditExplanation from "./RouteEditExplanation.svelte";
  import StatusSelect from "./StatusSelect.svelte";

  let route: PostRouteBody;

  let statuses: Status[];

  //values
  let name: string = "";
  let status: Status;

  let error: string | null = null;
  let successMessage: string | null = null;
  let user: User;
  let allPointsOfInterestLatLngs: LatLngTuple[] = [];

  let poiToScrollToAfterUpdate = -1;

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
  function newPointOfInterest() {
    const newPointOfInterest: PostPointOfInterestBody = {
      name: "",
      longitude: goudaCoordinates.longitude,
      latitude: goudaCoordinates.latitude,
      imageIds: [],
      details: [{ text: "" }],
      links: [{ url: "" }],
      categories: [],
    };
    route.pointsOfInterest.push(newPointOfInterest);
    route = route;
    updateAllPointsOfInterestLatLngs();
  }

  function updateAllPointsOfInterestLatLngs() {
    allPointsOfInterestLatLngs = route.pointsOfInterest.map(
      (pointOfInterest) => {
        return [
          pointOfInterest.latitude,
          pointOfInterest.longitude,
        ] as LatLngTuple;
      }
    );
  }

  function positionExchange(sourcePoiIndex: number, targetPoiIndex: number) {
    let oldPois = route.pointsOfInterest.slice();

    if (!oldPois[targetPoiIndex] || !oldPois[sourcePoiIndex]) return;
    route.pointsOfInterest[sourcePoiIndex] = oldPois[targetPoiIndex];
    route.pointsOfInterest[targetPoiIndex] = oldPois[sourcePoiIndex];
    route = route;
    poiToScrollToAfterUpdate = targetPoiIndex;
  }

  async function handleSave(event: Event) {
    try {
      if (!id) {
        let createdRoute = await createRoute(route);
        successMessage = "Succesfully created route!";
        await existingRouteToEditableRoute(createdRoute.id);
        id = createdRoute.id.toString();
        const url = new URL(window.location.origin + "/routes/edit");
        url.searchParams.set("id", createdRoute.id.toString());
        const newUrl = url.toString();
        window.history.pushState({ path: newUrl }, "", newUrl);
      } else {
        let updatedRoute = await updateRoute(parseInt(id), route);
        successMessage = "Succesfully updated route!";
        await existingRouteToEditableRoute(updatedRoute.id);
        id = updatedRoute.id.toString();
        const url = new URL(window.location.origin + "/routes/edit");
        url.searchParams.set("id", updatedRoute.id.toString());
        const newUrl = url.toString();
        window.history.pushState({ path: newUrl }, "", newUrl);
      }
    } catch (e: any) {
      console.log(e);
      error = e.response.data.error;
    }
  }

  async function existingRouteToEditableRoute(id: number) {
    let existingRoute = await getRouteById(id);
    console.log({ existingRoute });
    route = Object.assign({}, existingRoute, {
      imageIds: existingRoute.images.map((image) => image.id),
      pointsOfInterest: existingRoute.pointsOfInterest.map(
        (poi) =>
          Object.assign({}, poi, {
            imageIds: poi.images.map((image) => image.id),
          }) as PostPointOfInterestBody
      ),
    });
  }

  async function handleDeleteRoute() {
    try {
      if (id) {
        let deletedRoute = await deleteRouteById(parseInt(id));
        window.location.href = "/";
      }
    } catch (e: any) {
      error = e.response.data.error;
    }
  }

  onMount(async () => {
    user = await getMeUser();
    statuses = await getAllStatuses();

    let unpublishedStatus = statuses.find(
      (status) => status.name === "Unpublished"
    );
    let defaultStatusId = 0;
    if (unpublishedStatus) {
      defaultStatusId = unpublishedStatus.id;
    }

    if (!id) {
      //create it is then
      route = {
        name: "",
        imageIds: [],
        details: [{ text: "" }],
        links: [{ url: "" }],
        categories: [],
        statusId: defaultStatusId,
        pointsOfInterest: [],
      };
      newPointOfInterest();
      newPointOfInterest();
    } else {
      //get existing route
      await existingRouteToEditableRoute(parseInt(id));
    }

    // getRouteById(queryParams);
  });

  afterUpdate(async () => {
    console.log(route.statusId);
    updateAllPointsOfInterestLatLngs();
    //scroll if necessary
    if (poiToScrollToAfterUpdate !== -1) {
      let element = document.getElementById(`poi-${poiToScrollToAfterUpdate}`);
      console.log("scrolling", element);
      element?.scrollIntoView({ behavior: "smooth" });
      poiToScrollToAfterUpdate = -1;
    }
  });
</script>

<div class="edit">
  {#if user && route}
    <RouteEditExplanation />
    {#if id}
      <button
        class="button thick block error soft mb-s"
        on:click={handleDeleteRoute}>Delete route</button
      >
    {/if}
    <form on:submit|preventDefault={handleSave}>
      <RouteEditCard bind:route />
      <hr class="soft" />
      {#each route.pointsOfInterest as pointOfInterest, i}
        <PointOfInterestEditCard
          {positionExchange}
          bind:route
          position={i}
          bind:pointOfInterest
        />
      {/each}
      <button
        on:click={newPointOfInterest}
        type="button"
        class="button block secondary thick"
        href="#">New point of interest</button
      >
      <hr class="soft" />
      <div class="fixed">
        <div class="controls">
          <StatusSelect bind:value={route.statusId} />
          <button class="button thick  block primary" href="#">Save</button>
        </div>

        <div class="alerts">
          {#if error}
            <Alert
              type="error"
              onClose={() => {
                error = "";
              }}
            >
              {error}
            </Alert>
          {/if}
          {#if successMessage}
            <Alert
              type="success"
              onClose={() => {
                successMessage = "";
              }}
            >
              {successMessage}
            </Alert>
          {/if}
        </div>
      </div>
    </form>
  {:else}
    <div />
  {/if}
</div>
