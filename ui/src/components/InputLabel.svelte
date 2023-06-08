<script lang="ts">
  export let label: string;
  export let value: string;
  export let error: string;
</script>

<div class="input-label">
  <label >{label}</label>
  <div class="flex">
    <input bind:value {...$$restProps} />
  </div>
  {#if error}
    <label class="invalid">{error}</label>
  {/if}
</div>
