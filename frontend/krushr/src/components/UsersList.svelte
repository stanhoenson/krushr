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

<div>
  <div class="card">
    <div class="entity">
      <h2>Users</h2>
      <p>All existing users</p>
    </div>

    {#each users as user}
      <div class="user">
        <input type="email" value={user.email} disabled />
        <button class="button error">Delete</button>
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
