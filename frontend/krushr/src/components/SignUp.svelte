<script lang="ts">
  import { signIn, signUp } from "../requests/authentication";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  let email = "";
  let password = "";
  let confirmPassword = "";
  let error: string | null = "test";
  let user;

  async function handleSubmit(event: Event) {
    try {
      let user = await signUp({ email, password });
      console.log({ user });
      window.location.href = "/sign-in";
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
{#if error}
  <Alert
    onClose={() => {
      error = "";
    }}
    message={error}
  />
{/if}
