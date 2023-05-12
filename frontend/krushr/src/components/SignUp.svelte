<script lang="ts">
  import { onMount } from "svelte";

  import { signIn, signUp } from "../requests/authentication";
  import { getMeUser } from "../requests/users";
  import type { User } from "../types/models";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  let email = "";
  let password = "";
  let confirmPassword = "";
  let error: string | null = "";
  let user: User | null;

  async function handleSubmit(event: Event) {
    try {
      user = await signUp({ email, password });
      console.log({ user });
    } catch (e: any) {
      error = e.response.data.error;
    }
  }

  async function handleSignIn() {
    if (password === confirmPassword) {
      try {
        let token = await signIn({ email, password });
        window.location.href = "/";
      } catch (e: any) {
        error = e.response.data.error;
      }
    } else {
      error = "Passwords dont match";
    }
  }
  onMount(async () => {
    try {
      let user = await getMeUser();
      window.location.href = "/";
    } catch (e: any) {}
  });
</script>

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
    <button class="button block primary">Sign up</button>
  </section>
</form>
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
  {#if user}
    <Alert type="success">
      Succesfully register account! <a on:click={handleSignIn} href="#"
        >sign in</a
      > with your new account.
    </Alert>
  {/if}
</div>
