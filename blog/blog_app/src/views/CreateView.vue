<template>
  <div>
    <h2>Create New Blog</h2>
    <form @submit.prevent="handleSubmit">
      <label>Blog title:</label>
      <input type="text" required v-model="title" />
      <label>Blog body</label>
      <textarea required v-model="body"></textarea>
      <label>Blog author:</label>
      <select v-model="author">
        <option value="mario">mario</option>
        <option value="luigi">luigi</option>
        <option value="yoshi">yoshi</option>
      </select>
      <button :disabled="isSubmitting">Add Blog</button>
      <div v-if="error">{{ error }}</div>
    </form>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";
export default {
  setup() {
    const title = ref("");
    const body = ref("");
    const author = ref("mario");
    const isSubmitting = ref(false);
    const error = ref(null);

    const router = useRouter();

    const handleSubmit = async () => {
      isSubmitting.value = true;
      error.value = null;

      const blog = {
        title: title.value,
        body: body.value,
        author: author.value,
      };

      try {
        const response = await fetch("http://localhost:8000/blogs", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify(blog),
        });

        if (!response.ok) {
          throw new Error("Network response was not ok");
        }

        router.push("/");
      } catch (error) {
        console.error("Error:", error);
        error.value = error.message;
      } finally {
        isSubmitting.value = false;
      }
    };

    return { title, body, author, isSubmitting, handleSubmit };
  },
};
</script>

<style>
</style>