<script lang="ts">
  import { onDestroy } from "svelte";
  import { categories as categoriesStore } from "../stores/categories";
  import type { Category } from "../types/models";
  import type { PostCategoryBody } from "../types/request-bodies";

  export let selectedCategories: PostCategoryBody[] = [];
  export let handleCategoryToggle: (category: Category) => any;
  let selectedCategoriesStringOnly: string[];
  $: selectedCategoriesStringOnly = selectedCategories.map(
    (category) => category.name
  );

  let categories: Category[] = [];

  let unsubscribe = categoriesStore.subscribe((value) => {
    categories = value;
  });

  onDestroy(() => {
    unsubscribe();
  });
</script>

<div class="flex categories">
  {#each categories as category}
    <button
      type="button"
      on:click={handleCategoryToggle.bind(null, category)}
      class={`category ${
        selectedCategoriesStringOnly.includes(category.name) ? "selected" : ""
      }`}
      href="#">{category.name}</button
    >
  {/each}
</div>
