<template>
  <div class="editor-container">
    <div id="editor"></div>
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
// import { createDoc, DocumentPatcher, queryDoc } from "../api/document";
// import { defaultWSURL } from "../config";
// import { marshal, unmarshal } from "../marshaler";
import Quill from "quill";
import QuillCursors from "quill-cursors";
import { QuillBinding } from "y-quill";
import { WebsocketProvider } from "y-websocket";
import * as Y from "yjs";

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

const init = async (id: string | undefined) => {
  // let docId: string | null = null,
  //   content: string | null = null;
  // if (id && id !== "") {
  //   const { id: id1, content: content1 } = await queryDoc(id);
  //   docId = id1;
  //   content = content1;
  // } else {
  //   const { id: id1, content: content1 } = await createDoc();
  //   docId = id1;
  //   content = content1;
  // }
  // changeUrl(docId);

  // const patcher: DocumentPatcher = new DocumentPatcher(
  //   defaultWSURL,
  //   patchCallback,
  //   fetchCallback
  // );
  // patcher.start();

  const toolbarOptions = [
    ["bold", "italic", "underline", "strike"], // toggled buttons
    ["blockquote", "code-block"],

    [{ header: 1 }, { header: 2 }], // custom button values
    [{ list: "ordered" }, { list: "bullet" }],
    [{ script: "sub" }, { script: "super" }], // superscript/subscript
    [{ indent: "-1" }, { indent: "+1" }], // outdent/indent
    [{ direction: "rtl" }], // text direction

    [{ size: ["small", "large", "huge"] }],
    [{ header: [1, 2, 3, 4, 5, 6, false] }],

    ["clean"], // remove formatting button
  ];

  quill = new Quill(document.getElementById("editor")!, {
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
  quill.setText(""); // TODO

  const ydoc = new Y.Doc();
  const ytext = ydoc.getText("quill");
  const wsProvider = new WebsocketProvider(
    "ws://localhost:1234",
    "test-room",
    ydoc
  );
  wsProvider.on("status", (event: any) => {
    console.log(event.status);
  });
  const awareness = wsProvider.awareness;
  // awareness.setLocalStateField("user", {
  //   name: "pedro",
  //   color: "#ffb61e",
  // });
  const binding = new QuillBinding(ytext, quill, awareness);
  console.log(binding);
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
@import url(vditor/dist/index.css);

.editor-container {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
}

#editor {
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
