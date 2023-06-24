<script lang="ts">
  //TODO CLEAN THIS UP THIS IS DOODOO
  import { afterUpdate, onMount } from "svelte";
  import { deleteUser, getAllUsers, getMeUser } from "../requests/users";
  import { applicationState } from "../stores/application-state";
  import type { User } from "../types/models";
  import Alert from "./Alert.svelte";
  import { getErrorMessage } from "../utils/error";
  export let refresh: boolean = false;

  let users: User[] = [];
  let error: string = "";
  let authenticatedUser: User | null;

  let unsubscribe = applicationState.subscribe((value) => {
    authenticatedUser = value.authenticatedUser;
  });

  async function handleDelete(id: number) {
    try {
      let response = await deleteUser(id);
      users = await getAllUsers();
    } catch (e: any) {
      error = getErrorMessage(e);
    }
  }

  onMount(async () => {
    try {
      users = await getAllUsers();
    } catch (e: any) {}

    return () => {
      unsubscribe();
    };
  });

  afterUpdate(async () => {
    if (refresh) {
      users = await getAllUsers();
      refresh = false;
    }
  });
</script>

<div>
  <div class="card">
    <div class="entity">
      <h2>Users</h2>
      <p>All existing users</p>
    </div>

    {#each users as user}
      <div class="user">
        <input type="email" value={user.email} disabled />
        <button
          on:click={handleDelete.bind(null, user.id)}
          class="button error"
          disabled={user.id === authenticatedUser?.id}>Delete</button
        >
      </div>
    {/each}

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
  </div>
</div>
