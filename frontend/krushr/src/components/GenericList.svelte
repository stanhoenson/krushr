<script lang="ts">
  import { afterUpdate } from "svelte";

  import GenericRow from "./GenericRow.svelte";

  export let title: string;
  export let description: string;
  export let items: any[] = [];
  export let blacklistedKeys: string[] = [];
  export let onDelete: (id: number) => Promise<any>;
  export let onSave: (id: number | null, body: any) => Promise<any>;

  let keys: string[] = [];
  if (items.length > 0) {
    keys = Object.keys(items[0]);
    keys = keys.filter((key) => {
      return !blacklistedKeys.includes(key);
    });
  }

  function handleCreate() {}

  afterUpdate(() => {
    if (items.length > 0) {
      keys = Object.keys(items[0]);
      keys = keys.filter((key) => {
        return !blacklistedKeys.includes(key);
      });
    }
  });
</script>

<h2>{title}</h2>
<p>{description}</p>
<div style="width: 100%; overflow-x: scroll;">
  {#if items.length === 0}
    <div>No items</div>
  {:else}
    <button>Create</button>
    <!-- HEADER -->
    <div style="display: flex;">
      {#each keys as key}
        <div>{key}</div>
      {/each}
    </div>

    <!-- ITEMS -->
    <div style="display: flex;">
      {#each items as item}
        <GenericRow {item} {blacklistedKeys} {onSave} {onDelete} />
      {/each}
    </div>
  {/if}
</div>
