<script lang="ts">
  import { onDestroy } from "svelte";
  import { applicationState } from "../stores/application-state";
  import type { Category } from "../types/models";
  import type { PostCategoryBody } from "../types/request-bodies";

  export let disabled: boolean = false;
  export let selectedCategories: PostCategoryBody[] = [];
  export let handleCategoryToggle: (category: Category) => any;
  let selectedCategoriesStringOnly: string[];
  $: selectedCategoriesStringOnly = selectedCategories.map(
    (category) => category.name
  );

  let categories: Category[] = [];

  let unsubscribe = applicationState.subscribe((value) => {
    categories = value.categories;
  });


  onDestroy(() => {
    unsubscribe();
  });
</script>

<div class="flex categories">
  {#each categories as category}
    <button
      {disabled}
      type="button"
      on:click={handleCategoryToggle.bind(null, category)}
      class={`category ${
        selectedCategoriesStringOnly.includes(category.name) ? "selected" : ""
      }`}
      href="#">{category.name}</button
    >
  {/each}
</div>
