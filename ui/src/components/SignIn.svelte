<script lang="ts">
  import { onMount } from "svelte";

  import { sha256 } from "../utils/crypto";
  import { signIn } from "../requests/authentication";
  import { getMeUser } from "../requests/users";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";

  let email = "";
  let password = "";
  let error: string | null = "";

  async function handleSubmit(event: Event) {
    try {
      let token = await signIn({ email, password: await sha256(password) });
      window.location.href = "/";
    } catch (e: any) {
      error = e.response ? e.reponse.data.error : e;
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
