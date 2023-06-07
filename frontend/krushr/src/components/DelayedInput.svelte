<script lang="ts">
  import { afterUpdate } from "svelte";

  export let value: string = "";
  export let loading: boolean = false;
  let bufferedValue: string = "";
  export let delay: number = 500;

  let timeoutRef: any;

  function handleChange(event: Event) {
    loading = true;
    const target = event.target as HTMLInputElement;
    bufferedValue = target.value;

    clearTimeout(timeoutRef);

    timeoutRef = setTimeout(() => {
      loading = false;
      value = bufferedValue;
    }, delay);
  }
</script>

<input
  type="text"
  bind:value={bufferedValue}
  on:input={handleChange}
  {...$$restProps}
/>
