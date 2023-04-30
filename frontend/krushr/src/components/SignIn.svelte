<script lang="ts">
  import { signIn } from "../requests/authentication";
  let email = "";
  let password = "";
  let error: string | null = null;

  function handleSubmit(event: Event) {
    // perform form validation
    if (!email || !password) {
      error = "Please enter both email and password";
      return;
    }
    signIn({ email, password })
      .then((response) => {
        console.log(response);

      })
      .catch((error) => console.error(error));

    // submit form data via AJAX or navigate to new page
    console.log("Form submitted successfully!");
  }
</script>

<h1>Sign In</h1>
{#if error}
  <p style="color: red">{error}</p>
{/if}
<form on:submit|preventDefault={handleSubmit}>
  <label>
    Email:
    <input type="email" bind:value={email} />
  </label>
  <label>
    Password:
    <input type="password" bind:value={password} />
  </label>
  <button type="submit">Sign In</button>
</form>
