<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { getRouteById } from "../requests/routes";
  import { getAllStatuses } from "../requests/statuses";
  import { getMeUser } from "../requests/users";
  import { authenticatedUser } from "../stores/user";
  import type { PointOfInterest, Route, Status, User } from "../types/models";
  import Alert from "./Alert.svelte";
  import PointOfInterestEditCard from "./PointOfInterestEditCard.svelte";
  import RouteEditCard from "./RouteEditCard.svelte";
  import RouteEditExplanation from "./RouteEditExplanation.svelte";
  import StatusSelect from "./StatusSelect.svelte";

  let statuses: Status[];
  let pointsOfInterest: PointOfInterest[];

  //values
  let name: string = "";
  let status: Status;

  let error: string | null = null;
  let user: User | null;

  const unsubscribe = authenticatedUser.subscribe((value) => {
    user = value;
  });

  let queryParams = new URLSearchParams(window.location.search);

  let id = queryParams.get("id");
  if (!id) {
    //create it is then
  }

  onMount(async () => {
    if (!user) {
      window.location.href = "/";
    }

    statuses = await getAllStatuses();

    // getRouteById(queryParams);
  });

  onDestroy(() => {
    unsubscribe();
  });
</script>

<div class="edit">
  <RouteEditExplanation />
  <form>
    <RouteEditCard />
    <hr class="soft" />
    {#each pointsOfInterest as pointOfInterest}
      <PointOfInterestEditCard />
    {/each}
    <a id="save" class="button block secondary" href="#"
      >New point of interest</a
    >
    <hr class="soft" />
    <div class="fixed">
      <StatusSelect />
      <button id="save" class="button block primary" href="#">Save</button>
    </div>
  </form>
</div>
