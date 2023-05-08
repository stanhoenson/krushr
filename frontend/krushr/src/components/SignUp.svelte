<script lang="ts">
  import { signIn, signUp } from "../requests/authentication";
  import type { User } from "../types/models";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  let email = "";
  let password = "";
  let confirmPassword = "";
  let error: string | null = "test";
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
    try {
      let token = await signIn({ email, password });
      window.location.href = "/";
    } catch (e: any) {
      error = e.response.data.error;
    }
  }
</script>

<form class="sign-up" on:submit|preventDefault={handleSubmit}>
  <section class="card">
    <InputLabel
      bind:value={email}
      label={"Email"}
      error={""}
      type="email"
      required
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
<div class="alerts">
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
