<script lang="ts">
  import { onDestroy } from "svelte";
  import { applicationState } from "../stores/application-state";

  import type { Status } from "../types/models";

  export let disabled: boolean = false;
  export let value: number;

  let statuses: Status[] = [];

  let unsubscribe = applicationState.subscribe((value) => {
    statuses = value.statuses;
  });

  onDestroy(() => {
    unsubscribe();
  });
</script>

<div class="input status">
  <label for="status">Status</label>
  <select {disabled} bind:value name="status">
    {#each statuses as status}
      <option value={status.id}>{status.name}</option>
    {/each}
  </select>
</div>
