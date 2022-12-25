<template>
  <div class="editor-container">
    <div id="qeditor"></div>
    <div class="status">
      <div class="status-left">
        <span class="status-sub"
          >字数：
          <strong ref="textCounter">0</strong>
        </span>
        <span class="status-sub"
          >行数：
          <strong ref="lineCounter">0</strong>
        </span>
      </div>
      <div class="status-right">
        <span class="top">回到顶部</span>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted, ref } from "vue";
import * as localforage from "localforage";
import Quill from "quill";
import QuillCursors from "quill-cursors";
import { QuillBinding } from "y-quill";
import { WebsocketProvider } from "y-websocket";
import * as Y from "yjs";

import { createDoc, queryDoc } from "../api/document";
import { defaultWSURL } from "../config";
import { Counter } from "../modules/quill/counter";

export interface Document {
  content: string;
}

const props = defineProps<{
  id?: string;
}>();

const textCounter = ref<typeof HTMLElement | null>(null);
const lineCounter = ref<typeof HTMLElement | null>(null);

Quill.register("modules/cursors", QuillCursors);
Quill.register("modules/counter", Counter);

onMounted(async () => {
  await init(props.id);
});

let quill: Quill | null = null;

const toolbarOptions = [
  ["bold", "italic", "underline", "strike"],
  [{ header: 1 }, { header: 2 }, { header: 3 }, { header: 4 }],
  [{ list: "ordered" }, { list: "bullet" }, { indent: "+1" }, { indent: "-1" }],
  ["align", "color", "background"],
  ["blockquote", "code-block", "link", "image"],
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
        cursors: {},
        toolbar: toolbarOptions,
        history: {
          userOnly: true,
        },
        counter: {
          wordElement: textCounter.value!,
          wordUnit: "字",
          lineElement: lineCounter.value!,
          lineUnit: "行",
        },
      },
      placeholder: "Type what you want...",
      theme: "snow",
    });
  }
  quill.setText(content);
  console.log("docId: ", docId);
  const ydoc = new Y.Doc();
  const ytext = ydoc.getText("quill");
  const wsProvider = new WebsocketProvider(defaultWSURL, "test-doc", ydoc);
  wsProvider.on("status", (event: any) => {
    console.log(event.status);
  });
  const awareness = wsProvider.awareness;
  /*const binding = */ new QuillBinding(ytext, quill, awareness);
};

const fresh = () => {
  init(props.id);
};

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
  width: 80%;
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

.status-sub {
  padding-left: 16px;
}

button.ql-header svg {
  display: none;
}

.ql-header[value="4"]:after {
  clear: both;
  content: "H4";
  display: table;
  font-weight: 600;
  margin-top: -2px;
  margin-left: 1px;
  font-size: 14px;
}

.ql-header[value="3"]:after {
  clear: both;
  content: "H3";
  display: table;
  font-weight: 600;
  margin-top: -2px;
  margin-left: 1px;
  font-size: 14px;
}

.ql-header[value="2"]:after {
  clear: both;
  content: "H2";
  display: table;
  font-weight: 600;
  margin-top: -2px;
  margin-left: 1px;
  font-size: 14px;
}

.ql-header[value="1"]:after {
  clear: both;
  content: "H1";
  display: table;
  font-weight: 600;
  margin-top: -2px;
  margin-left: 1px;
  font-size: 14px;
}

button.ql-header.ql-active {
  color: #3891d0;
}
</style>
