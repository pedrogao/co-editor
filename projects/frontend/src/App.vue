<template>
  <!-- <img
      alt="Vue logo"
      class="logo"
      src="@/assets/logo.svg"
      width="125"
      height="125"
    /> -->

  <div class="container mx-auto bg-gray-200 rounded-xl shadow border p-8 m-10">
    <p class="text-2xl text-gray-700 font-bold mb-5">Type what you want!</p>
    <MarkdownEditor />
  </div>

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

<style scoped></style>
