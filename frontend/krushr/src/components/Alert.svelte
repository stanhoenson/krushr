<script lang="ts">
  import XMark from "./icons/XMark.svelte";

  export let onClose: () => any;
  export let message: string;
</script>

<div id="half-width" class="alert block error soft">
  <span class="message">{message}</span>
  <button class="icon" on:click={onClose}><XMark /></button>
</div>
