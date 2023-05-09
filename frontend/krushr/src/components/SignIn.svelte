<script lang="ts">
  import { signIn } from "../requests/authentication";
  import Alert from "./Alert.svelte";
  import InputLabel from "./InputLabel.svelte";
  let email = "";
  let password = "";
  let error: string | null = "test";

  async function handleSubmit(event: Event) {
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
{#if error}
  <Alert
    onClose={() => {
      error = "";
    }}
    message={error}
  />
{/if}
