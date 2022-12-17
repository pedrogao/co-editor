<template>
  <div class="app">
    <header class="header">
      <input class="title" placeholder="请填写标题" v-model="inputId" />
      <div class="right-box">
        <button @click="openFile" class="menu">文件</button>
        <div class="more">↓</div>
        <ul class="more-list">
          <li class="item">写文章</li>
          <li class="item">查看</li>
        </ul>
      </div>
    </header>
    <div class="editor">
      <MarkdownEditor ref="editor" :id="inputId" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import MarkdownEditor from "./components/MarkdownEditor.vue";

const inputId = ref("");
const editor = ref<typeof MarkdownEditor | null>(null);

const openFile = () => {
  console.log(inputId.value);
  editor.value?.fresh();
};
</script>

<style scoped>
.app {
  display: block;
  width: 100%;
  height: 100%;
}

.header {
  display: flex;
  align-items: center;
  padding: 0 27px;
  height: 4.334rem;
  background-color: #fff;
  border-bottom: 1px solid #ddd;
  z-index: 100;
}

.title {
  flex: 1 1 auto;
  height: 100%;
  margin: 0;
  padding: 0;
  font-size: 24px;
  font-weight: 500;
  color: #1d2129;
  border: none;
  outline: none;
}

.right-box {
  display: flex;
  align-items: center;
  justify-content: flex-end;
}

.menu {
  height: 32px;
  padding: 2px 16px;
  font-size: 14px;
  line-height: 22px;
  border: 1px solid #1d7dfa;
  border-radius: 2px;
  cursor: pointer;
  color: #fff;
  box-sizing: border-box;
  background-color: #1d7dfa;
}

.more {
  flex: 0 0 auto;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1.667rem;
  border-left: 1px solid hsla(0, 0%, 100%, 0.1);
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
  box-sizing: border-box;
}

.more-list {
  padding: 0.417rem 0;
  font-size: 1.167rem;
  color: var(--juejin-font-2);
  background-color: #fff;
  display: none;
  z-index: 1;
  box-shadow: 0 0 24px rgb(81 87 103 / 16%);
  border-radius: 6px;
  width: 10rem;
}

.item {
  margin: 0 8px;
  padding: 6px 5px 5px 0.8rem;
  border-radius: 4px;
  line-height: 24px;
  font-weight: 400;
}

.editor {
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: auto;
  font-family: -apple-system, BlinkMacSystemFont, Segoe UI, Helvetica, Arial,
    sans-serif, Apple Color Emoji, Segoe UI Emoji;
  color: #24292e;
  background-color: #fff;
}
</style>
