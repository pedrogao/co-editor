<template>
  <div class="container mx-auto bg-gray-200 rounded-xl shadow border p-8 m-10">
    <p class="text-2xl text-gray-700 font-bold mb-5">Type what you want!</p>
    <MarkdownEditor />
  </div>
</template>

<script setup lang="ts">
import MarkdownEditor from "../components/MarkdownEditor.vue";
</script>
