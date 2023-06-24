<script lang="ts">
  import { getErrorMessage } from "../utils/error";

  export let item: any;
  export let blacklistedKeys: string[] = [];
  export let onDelete: (id: number) => Promise<any>;
  export let onSave: (id: number | null, body: any) => Promise<any>;

  let editableItem: any = structuredClone(item);

  let editing = false;
  let error = "";

  let keys = Object.keys(item);

  keys = keys.filter((key) => {
    return !blacklistedKeys.includes(key);
  });

  function handleEdit() {
    editing = true;
  }

  function handleCancel() {
    editing = false;
    editableItem = structuredClone(item);
  }

  async function handleSave() {
    try {
      onSave(item.id, editableItem);
    } catch (e: any) {
      error = getErrorMessage(e);
    }
  }
  async function handleDelete() {}
</script>

{#if !editing}
  <!-- ITEMS -->
  {#each keys as key}
    <div>{editableItem[key]}</div>
  {/each}
  <button on:click={handleEdit}>Edit</button>
  <button on:click={handleDelete}>Delete</button>
{:else}
  <!-- ITEMS -->
  {#each keys as key}
    <input bind:value={editableItem[key]} />
  {/each}
  <button on:click={handleSave}>Save</button>
  <button on:click={handleCancel}>Cancel</button>
  {#if error}
    <div>{error}</div>
  {/if}
{/if}
