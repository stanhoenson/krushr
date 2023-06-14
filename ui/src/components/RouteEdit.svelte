<script lang="ts">
  import type { LatLngTuple } from "leaflet";

  import { afterUpdate, onDestroy, onMount, tick } from "svelte";
  import { goudaCoordinates } from "../constants";
  import {
    createRoute,
    deleteRouteById,
    getRouteById,
    updateRoute,
  } from "../requests/routes";
  import { getAllStatuses } from "../requests/statuses";
  import { getMeUser } from "../requests/users";
  import { applicationState } from "../stores/application-state";
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

  let existingRoute: Route;
  let route: PostRouteBody;
  let viewOnly: boolean = false;

  let statuses: Status[];
  let didSomething: boolean = false;

  //values
  let name: string = "";
  let status: Status;

  let error: string | null = null;
  let successMessage: string | null = null;
  let user: User | null;

  let poiToScrollToAfterUpdate = -1;

  let unsubscribe = applicationState.subscribe((applicationState) => {
    statuses = applicationState.statuses;
    user = applicationState.authenticatedUser;
  });

  let queryParams = new URLSearchParams(window.location.search);
  let id = queryParams.get("id");
  function newPointOfInterest() {
    const newPointOfInterest: PostPointOfInterestBody = {
      name: "",
      longitude: goudaCoordinates.longitude,
      latitude: goudaCoordinates.latitude,
      imageIds: [],
      details: [{ text: "" }],
      links: [{ url: "", text: "" }],
      categories: [],
      support: false,
    };
    route.pointsOfInterest.push(newPointOfInterest);
    route = route;
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
        if (didSomething) {
          window.removeEventListener("beforeunload", beforeUnloadHandler);
          didSomething = false;
        }
      }
    } catch (e: any) {
      console.log(e);
      error = e.response ? e.reponse.data.error : e;
    }
  }

  async function existingRouteToEditableRoute(id: number) {
    existingRoute = await getRouteById(id);
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
        const confirmed = window.confirm(
          "Are you sure you want to delete this route?"
        );

        if (confirmed) {
          let deletedRoute = await deleteRouteById(parseInt(id));
          if (didSomething)
            window.removeEventListener("beforeunload", beforeUnloadHandler);
          window.location.href = "/";
        }
      }
    } catch (e: any) {
      error = e.response ? e.reponse.data.error : e;
    }
  }

  onMount(async () => {
    // try {
    //   user = await getMeUser();
    // } catch (e: any) {}
    // statuses = await getAllStatuses();

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
        links: [{ url: "", text: "" }],
        categories: [],
        statusId: defaultStatusId,
        pointsOfInterest: [],
      };
      newPointOfInterest();
      newPointOfInterest();
    } else {
      //get existing route
      try {
        await existingRouteToEditableRoute(parseInt(id));
      } catch (e: any) {
        window.location.href = "/404";
      }
    }

    viewOnly = !!(
      !user ||
      (existingRoute &&
        !(user && user.role.name === "Admin") &&
        !(user && user.id === existingRoute.userId))
    );

    // viewOnly = !!(
    //   (existingRoute && existingRoute.userId === (user ? user.id : -1)) ||
    //   (user && (user ? user.role.name === "Admin" : false))
    // );

    // getRouteById(queryParams);
    return () => {
      if (didSomething)
        window.removeEventListener("beforeunload", beforeUnloadHandler);
    };
  });

  afterUpdate(async () => {
    viewOnly = !!(
      !user ||
      (existingRoute &&
        !(user && user.role.name === "Admin") &&
        !(user && user.id === existingRoute.userId))
    );
    if (!didSomething && !viewOnly) {
      window.addEventListener("beforeunload", beforeUnloadHandler);
      didSomething = true;
    }
    if (viewOnly) {
      window.removeEventListener("beforeunload", beforeUnloadHandler);
    }
    //scroll if necessary
    if (poiToScrollToAfterUpdate !== -1) {
      let element = document.getElementById(`poi-${poiToScrollToAfterUpdate}`);
      element?.scrollIntoView({ behavior: "smooth" });
      poiToScrollToAfterUpdate = -1;
    }
  });

  onDestroy(() => {
    unsubscribe();
  });

  const beforeUnloadHandler = (e: BeforeUnloadEvent) => {
    e.preventDefault();
    return (e.returnValue =
      "Are you sure you want to leave? All changes will be lost"); // This line is required for some browsers
  };
</script>

<div class="edit">
  {#if route}
    <RouteEditExplanation />
    {#if id}
      <button
        disabled={viewOnly}
        class="button thick block error soft mb-s"
        on:click={handleDeleteRoute}>Delete route</button
      >
    {/if}
    <form on:submit|preventDefault={handleSave}>
      <RouteEditCard bind:route {viewOnly} />
      <hr class="soft" />
      {#each route.pointsOfInterest as pointOfInterest, i}
        <PointOfInterestEditCard
          {viewOnly}
          {positionExchange}
          bind:route
          position={i}
          bind:pointOfInterest
        />
      {/each}
      <button
        disabled={viewOnly}
        on:click={newPointOfInterest}
        type="button"
        class="button block secondary thick"
        href="#">New point of interest</button
      >
      <hr class="soft" />
      <div class="fixed">
        <div class="controls">
          <StatusSelect disabled={viewOnly} bind:value={route.statusId} />
          <button
            disabled={viewOnly}
            class="button thick  block primary"
            href="#">Save</button
          >
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
