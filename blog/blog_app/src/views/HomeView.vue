<template>
  <div class="home">
    <h1>Home</h1>
    <img alt="Vue logo" src="../assets/logo.png" />
    <div v-if="error">{{ error }}</div>
    <div v-if="blogs.length">
      <BlogList :blogs="blogs" :title="title" />
    </div>
    <div v-else>Loading...</div>
  </div>
</template>

<script>
import { ref } from "vue";
import BlogList from "../blogs/BlogList.vue";
import getBlogs from "../composables/getBlogs";

export default {
  name: "HomeView",
  components: { BlogList },
  setup() {
    const title = ref("All Blogs");
    const { blogs, error, load } = getBlogs();

    load();

    return { title, blogs, error };
  },
};
</script>
