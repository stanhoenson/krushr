<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import { signOut } from "../requests/authentication";
  import {
    applicationState,
    initializeApplicationState,
    resetApplicationState,
  } from "../stores/application-state";
  import type { User } from "../types/models";

  let user: User | null;

  let unsubscribe = applicationState.subscribe((value) => {
    user = value.authenticatedUser;
  });

  async function handleSignOut() {
    try {
      await signOut();
      await resetApplicationState();
      window.location.href = "/";
    } catch (e: any) {}
  }

  onDestroy(() => {
    unsubscribe();
  });
</script>

<nav class="flex">
  <div class="flex">
    <a class="button" href="/">Routes</a>
  </div>
  <div class="flex">
    {#if !user}
      <!-- <a class="button" href="/sign-up">Sign up</a> -->
      <a class="button primary" href="/sign-in">Sign in</a>
    {:else}
      <!-- TODO magic string not good -->
      {#if user.role.name === "Admin"}
        <a class="button" href="/users">Users</a>
      {/if}
      <button on:click={handleSignOut} class="button secondary">Sign out</button
      >
    {/if}
  </div>
</nav>
