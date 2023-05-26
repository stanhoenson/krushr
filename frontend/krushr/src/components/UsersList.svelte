<script lang="ts">
  //TODO CLEAN THIS UP THIS IS DOODOO
  import { afterUpdate, onMount } from "svelte";
  import { deleteUser, getAllUsers, getMeUser } from "../requests/users";
  import type { User } from "../types/models";
  import Alert from "./Alert.svelte";
  export let refresh: boolean = false;

  let users: User[] = [];
  let error: string = "";
  let user: User;

  async function handleDelete(id: number) {
    try {
      let resposne = await deleteUser(id);
      users = await getAllUsers();
    } catch (e: any) {
      error = e.response.data.error;
    }
  }

  onMount(async () => {
    try {
      user = await getMeUser();
      users = await getAllUsers();
    } catch (e: any) {}
  });

  afterUpdate(async () => {
    if (refresh) {
      users = await getAllUsers();
      refresh = false;
    }
  });
</script>

<div class="entity">
  <h2>Users</h2>
  <p>All existing users</p>
</div>
<div class="table card">
  <div class="column">
    <p class="field">ID</p>
    {#each users as user}
      <input type="number" value={user.id} disabled />
    {/each}
  </div>
  <div class="column">
    <p class="field">Email</p>
    {#each users as user}
      <input type="email" value={user.email} disabled />
    {/each}
  </div>
  <div class="column">
    <p class="field">Action</p>
    {#each users as innerUser}
      <button
        on:click={handleDelete.bind(null, innerUser.id)}
        disabled={user.id === innerUser.id}
        class="button primary">Delete</button
      >
    {/each}
  </div>
</div>
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
