<script lang="ts">
  import { onMount } from "svelte";
  import { signOut } from "../requests/authentication";

  import { getMeUser } from "../requests/users";

  let loggedIn: boolean;

  async function handleSignOut() {
    try {
      await signOut();
      window.location.reload();
    } catch (e: any) {}
  }

  onMount(async () => {
    try {
      let user = await getMeUser();
      loggedIn = true;
    } catch (e) {}
  });
</script>

<nav class="flex">
  <div class="flex">
    <a id="active" class="button" href="/">Routes</a>
  </div>
  <div class="flex">
    {#if !loggedIn}
      <a class="button" href="/sign-up">Sign up</a>
      <a class="button primary" href="/sign-in">Sign in</a>
    {:else}
      <button on:click={handleSignOut} class="button secondary">Sign out</button
      >
    {/if}
  </div>
</nav>
