<script lang="ts">
  import XMark from "./icons/XMark.svelte";

  export let onClose: null | (() => any) = null;
  export let type: "error" | "success" = "error";
</script>

<div class={`alert block ${type} soft`} {...$$restProps}>
  <span class="message"><slot /></span>
  {#if onClose !== null}
    <div class="icon" on:click={onClose}><XMark /></div>
  {/if}
</div>
