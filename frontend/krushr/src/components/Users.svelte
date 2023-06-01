<script lang="ts">
  import { onMount } from "svelte";

  import { signIn, signUp } from "../requests/authentication";
  import { getMeUser } from "../requests/users";
  import type { User } from "../types/models";
  import { sha256 } from "../utils/crypto";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  import UsersList from "./UsersList.svelte";
  let refresh: boolean = false;
  let email = "";
  let password = "";
  let confirmPassword = "";
  let error: string | null = "";
  let createdUser: User | null;
  let user: User;

  async function handleSubmit(event: Event) {
    if (password === confirmPassword) {
      try {
        createdUser = await signUp({ email, password: await sha256(password) });
        refresh = true;
      } catch (e: any) {
        error = e.response.data.error;
      }
    } else {
      error = "Passwords dont match";
    }
  }

  onMount(async () => {
    try {
      user = await getMeUser();
      //TODO magic string
      if (user.role.name !== "Admin") {
        window.location.href = "/";
      }
    } catch (e: any) {
      window.location.href = "/";
    }
  });
</script>

<div class="grid">
  <form class="sign-up" on:submit|preventDefault={handleSubmit}>
    <section class="card">
      <InputLabel
        bind:value={email}
        label={"Email"}
        error={""}
        type="email"
        required
        autofocus={true}
      />
      <InputLabel
        bind:value={password}
        label={"Password"}
        error={""}
        name="password"
        type="password"
        required
      />
      <InputLabel
        bind:value={confirmPassword}
        label={"Confirm password"}
        error={""}
        name="confirmPassword"
        type="password"
        required
      />

      <hr />
      <button class="button block primary">Create user</button>
    </section>
    <div class="alerts half-width">
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
      {#if createdUser}
        <Alert
          onClose={() => {
            createdUser = null;
          }}
          type="success">Succesfully created user!</Alert
        >
      {/if}
    </div>
  </form>
  <UsersList bind:refresh />
</div>
