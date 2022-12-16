<template>
  <div
    class="bg-blue-100 rounded-lg py-5 px-6 text-base text-blue-700 mb-3"
    role="alert"
  >
    {{ msg }}
  </div>
</template>

<script setup lang="ts">
defineProps<{
  msg: string;
}>();
</script>

<style scoped></style>
