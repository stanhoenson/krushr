<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { signOut } from "../requests/authentication";
  import { authenticatedUser } from "../stores/user";
  import type { User } from "../types/models";

  let user: User | null;

  let unsubscribe = authenticatedUser.subscribe((value) => (user = value));

  async function handleSignOut() {
    try {
      await signOut();
      window.location.href = "/";
    } catch (e: any) {}
  }

  onDestroy(() => {
    unsubscribe();
  });
</script>

<nav class="flex">
  <div class="flex">
    <a id="active" class="button" href="/">Routes</a>
  </div>
  <div class="flex">
    {#if !user}
      <!-- <a class="button" href="/sign-up">Sign up</a> -->
      <a class="button primary" href="/sign-in">Sign in</a>
    {:else}
      <!-- TODO magic string not good -->
      {#if user.role.name === "Admin"}
        <a class="button primary" href="/admin">Admin</a>
      {/if}
      <button on:click={handleSignOut} class="button secondary">Sign out</button
      >
    {/if}
  </div>
</nav>
