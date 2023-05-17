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
  import AdminPanelExplanation from "./AdminPanelExplanation.svelte";

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

<div class="admin">
	<AdminPanelExplanation />

	<GenericList
		title="categories"
		description="Categories can be applied to routes and points of interest"
		items={categories}
		blacklistedKeys={["routes", "pointsOfInterest"]}
		onSave={onCategorySave}
		onDelete={onCategoryDelete}
	/>

	<div class="table card">
		<div class="column">
			<p class="field">ID</p>
			<input type="number" name="id" value="0" disabled>
			<input type="number" name="id" value="1" disabled>
			<input type="number" name="id" value="2" disabled>
		</div>
		<div class="column">
			<p class="field">Name</p>
			<input type="text" name="name" value="Nature">
			<input type="text" name="name" value="Shopping" disabled>
			<input type="text" name="name" value="Architecture" disabled>
		</div>
		<div class="column">
			<p class="field">Position</p>
			<input type="number" name="position" value="2">
			<input type="number" name="position" value="3" disabled>
			<input type="number" name="position" value="1" disabled>
		</div>
		<div class="column">
			<p class="field">Action</p>
			<button class="button primary">Save</button>
			<button class="button secondary">Edit</button>
			<button class="button secondary">Edit</button>
		</div>
	</div>
	<hr class="soft">
	<GenericList
		title="categories"
		description="Categories can be applied to routes and points of interest"
		items={categories}
		blacklistedKeys={["routes", "pointsOfInterest"]}
		onSave={onCategorySave}
		onDelete={onCategoryDelete}
	/>

	<div class="table card">
		<div class="column">
			<p class="field">ID</p>
			<input type="number" name="id" value="0" disabled>
			<input type="number" name="id" value="1" disabled>
			<input type="number" name="id" value="2" disabled>
		</div>
		<div class="column">
			<p class="field">Name</p>
			<input type="text" name="name" value="Nature">
			<input type="text" name="name" value="Shopping" disabled>
			<input type="text" name="name" value="Architecture" disabled>
		</div>
		<div class="column">
			<p class="field">Position</p>
			<input type="number" name="position" value="2">
			<input type="number" name="position" value="3" disabled>
			<input type="number" name="position" value="1" disabled>
		</div>
		<div class="column">
			<p class="field">Action</p>
			<button class="button primary">Save</button>
			<button class="button secondary">Edit</button>
			<button class="button secondary">Edit</button>
		</div>
	</div>
</div>
