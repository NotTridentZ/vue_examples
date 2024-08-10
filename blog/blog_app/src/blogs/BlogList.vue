<template>
  <div>
    <h2>{{ title }}</h2>
    <div v-if="blogs.length">
      <div v-for="blog in blogs" :key="blog.id">
        <router-link :to="{ name: 'BlogDetails', params: { id: blog.id } }">
          <h2>{{ blog.title }}</h2>
        </router-link>
        <p>Written by: {{ blog.author }}</p>
      </div>
    </div>
    <div v-else>Loading blogs...</div>
    <div v-if="error">{{ error }}</div>
  </div>
</template>

<script>
export default {
  props: ["blogs", "title"],
  components: {},
  setup() {},
};
</script>

<style>
</style>