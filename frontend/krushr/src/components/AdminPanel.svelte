<script lang="ts">
  import { onMount } from "svelte";
  import {
    createCategory,
    deleteCategoryById,
    getAllCategories,
    updateCategory,
  } from "../requests/categories";

  import type { Category } from "../types/models";
  import type {
    PostCategoryBody,
    PutCategoryBody,
  } from "../types/request-bodies";
  import GenericList from "./GenericList.svelte";

  let categories: Category[] = [];

  onMount(async () => {
    try {
      categories = await getAllCategories();
    } catch (e: any) {
      console.log(e);
    }
  });

  //we can throw errors in these functions to catch in the GenericRow

  async function onCategorySave(id: number | null, body: Category) {
    let postCategoryBody: PostCategoryBody = {
      name: body.name,
      position: body.position,
    };
    if (id) {
      await updateCategory(id, postCategoryBody);
    } else {
      await createCategory(postCategoryBody);
    }
  }

  async function onCategoryDelete(id: number) {
    await deleteCategoryById(id);
  }
</script>

<h1>Admin Panel</h1>

<GenericList
  title="categories"
  description="thing"
  items={categories}
  blacklistedKeys={["routes", "pointsOfInterest"]}
  onSave={onCategorySave}
  onDelete={onCategoryDelete}
/>
