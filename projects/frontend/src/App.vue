<template>
  <!-- <img
      alt="Vue logo"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    /> -->

  <!-- <HelloWorld msg="You did it!" /> -->
  <MarkdownEditor />

  <!-- <nav> -->
  <!-- <RouterLink to="/">Home</RouterLink> -->
  <!-- <RouterLink to="/about">About</RouterLink> -->
  <!-- <RouterLink to="/editor">Editor</RouterLink> -->
  <!-- </nav> -->
  <!-- 路由视图 -->
  <!-- <RouterView></RouterView> -->
</template>

<script setup lang="ts">
// import { RouterLink, RouterView } from "vue-router";
import MarkdownEditor from "./components/MarkdownEditor.vue";
</script>

<style scoped>
.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}
</style>
