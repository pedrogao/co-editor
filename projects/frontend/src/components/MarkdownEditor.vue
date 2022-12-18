<template>
  <div class="editor-container">
    <div id="qeditor"></div>
    <div class="status">
      <div class="status-left">
        <span
          >字数：
          <strong>2</strong>
        </span>
        <span
          >行数：
          <strong>3</strong>
        </span>
      </div>
      <div class="status-right">
        <span class="top">回到顶部</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import * as localforage from "localforage";
import Quill from "quill";
import QuillCursors from "quill-cursors";
import { QuillBinding } from "y-quill";
import { WebsocketProvider } from "y-websocket";
import * as Y from "yjs";

import { createDoc, queryDoc } from "../api/document";
import { defaultWSURL } from "../config";

export interface Document {
  content: string;
}

const props = defineProps<{
  id?: string;
}>();

Quill.register("modules/cursors", QuillCursors);

onMounted(async () => {
  await init(props.id);
});

let quill: Quill | null = null;

const toolbarOptions = [
  ["bold", "italic", "underline", "strike"],
  ["blockquote", "code-block"],
  [{ header: 1 }, { header: 2 }],
  [{ list: "ordered" }, { list: "bullet" }],
  [{ script: "sub" }, { script: "super" }],
  [{ indent: "-1" }, { indent: "+1" }],
  [{ direction: "rtl" }],
  [{ size: ["small", "large", "huge"] }],
  [{ header: [1, 2, 3, 4, 5, 6, false] }],
  ["clean"],
];

const init = async (id: string | undefined) => {
  let fn = null;
  if (id && id !== "") {
    fn = queryDoc;
  } else {
    fn = createDoc;
  }
  const { id: docId, content } = await fn(id!);
  changeUrl(docId);
  if (!quill) {
    quill = new Quill(document.getElementById("qeditor")!, {
      modules: {
        cursors: true,
        toolbar: toolbarOptions,
        history: {
          // Local undo shouldn't undo changes
          // from remote users
          userOnly: true,
        },
      },
      placeholder: "Type what you want...",
      theme: "snow",
    });
  }
  quill.setText(content);

  const ydoc = new Y.Doc();
  const ytext = ydoc.getText("quill");
  const wsProvider = new WebsocketProvider(defaultWSURL, docId, ydoc);
  wsProvider.on("status", (event: any) => {
    console.log(event.status);
  });
  const awareness = wsProvider.awareness;
  // awareness.setLocalStateField("user", {
  //   name: "pedro",
  //   color: "#ffb61e",
  // });
  /*const binding = */ new QuillBinding(ytext, quill, awareness);
};

const fresh = () => {
  init(props.id);
};

const patchCallback = (message: string) => {
  console.log(message);
};
const fetchCallback = (message: string) => {};

const loadFromLocal = async (docId: string) => {
  return localforage.getItem(docId);
};
const saveToLocal = async (docId: string, content: string) => {
  return localforage.setItem(docId, content);
};

const changeUrl = (docId: string) => {
  const origin = window.location.href;
  const sep = origin.indexOf("/");
  const href = `${origin.substring(0, sep + 1)}${docId}`;
  window.history.pushState({}, "", href);
};

defineExpose({
  fresh,
});
</script>

<style>
.editor-container {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
}

#qeditor {
  height: calc(100vh - 100px) !important;
}

.status {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  font-size: 12px;
  line-height: 24px;
  border-top: 1px solid #e1e4e8;
  user-select: none;
}

.status-left {
  float: left;
}

.status-right {
  float: right;
}

.top {
  padding-right: 10px;
}

span {
  padding-left: 16px;
}
</style>
