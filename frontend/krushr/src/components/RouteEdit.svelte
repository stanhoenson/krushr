<script lang="ts">
  import { onMount } from "svelte";
  import { getMeUser } from "../requests/users";
  import type { User } from "../types/models";
  import Alert from "./Alert.svelte";
  import PointOfInterestEditCard from "./PointOfInterestEditCard.svelte";
  import RouteEditCard from "./RouteEditCard.svelte";
  import RouteEditExplanation from "./RouteEditExplanation.svelte";
  let error: string | null = null;
  let user: User;
  // let queryParams = new URLSearchParams(window.location.search);
  // console.log({ queryParams });

  onMount(async () => {
    try {
      user = await getMeUser();
      console.log(user);
    } catch (e) {}
  });
</script>

<RouteEditExplanation />
<form>
  <RouteEditCard />
  <hr class="soft" />
  <PointOfInterestEditCard />
  <a id="save" class="button block secondary" href="#">New point of interest</a>
  <hr class="soft" />
	<div class="fixed">
		<div class="input status">
			<label for="status">Status</label>
			<select name="status">
				<option value="unpublished">Unpublished</option>
				<option value="published">Published</option>
			</select>
		</div>
		<button id="save" class="button block primary" href="#">Save</button>
	</div>
</form>
