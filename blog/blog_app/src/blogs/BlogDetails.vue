<template>
  <div>
    <h2>Blog Details</h2>
    <div v-if="error">
      {{ error }}
    </div>
    <template v-else-if="!blog">
      <div>Loading blog details...</div>
    </template>
    <template v-else>
      <article>
        <h2>{{ blog.title }}</h2>
        <p>Written by: {{ blog.author }}</p>
        <div>{{ blog.body }}</div>
        <button :disabled="isDeleting" @click="handleClick">Delete</button>
      </article>
    </template>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";
import getBlog from "../composables/getBlog";

export default {
  props: ["id"],
  setup(props) {
    const { blog, error, load } = getBlog(props.id);
    const router = useRouter();
    const isDeleting = ref(false);

    load();

    const handleClick = async () => {
      isDeleting.value = true;

      try {
        const response = await fetch(
          "http://localhost:8000/blogs/" + props.id,
          {
            method: "DELETE",
          }
        );

        if (!response.ok) {
          throw new Error("Failed to delete blog");
        }

        router.push("/");
      } catch (error) {
        console.error("Error deleting blog:", error);
        error.value = "Failed to delete blog. Please try again.";
      } finally {
        isDeleting.value = false;
      }
    };

    return { blog, error, handleClick, isDeleting };
  },
};
</script>

<style>
</style>