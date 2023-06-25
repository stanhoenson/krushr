<script lang="ts">
  import { onMount } from "svelte";

  import { sha256, sha256WithSalt } from "../utils/crypto";
  import { signIn } from "../requests/authentication";
  import { getMeUser } from "../requests/users";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  import { getErrorMessage } from "../utils/error";

  let email = "";
  let password = "";
  let error: string | null = "";

  async function handleSubmit(event: Event) {
    try {
      await signIn({ email, password: await sha256WithSalt(password) });
      window.location.href = "/";
    } catch (e: any) {
      error = getErrorMessage(e);
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
    <hr />
    <button class="button block primary">Sign in</button>
  </section>
</form>
<div class="alerts half-width">
  {#if error}
    <Alert
      onClose={() => {
        error = "";
      }}>{error}</Alert
    >
  {/if}
</div>
