<template>
  <div class="editor-container">
    <div id="vditor"></div>
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
import Vditor from "vditor";
import * as Automerge from "@automerge/automerge";
// import * as localforage from "localforage";
import { createDoc, DocumentPatcher } from "../api/document";
import { defaultWSURL } from "../config";

export interface Document {
  content: string;
}

let vditor: Vditor | null = null;
let doc: Automerge.Doc<Document> | null = null;
let docId: string;
let patcher: DocumentPatcher;

onMounted(async () => {
  await init();
});

const init = async () => {
  docId = await createDoc();
  doc = Automerge.init();
  patcher = new DocumentPatcher(defaultWSURL);
  patcher.start();

  vditor = new Vditor("vditor", {
    width: "100%",
    toolbarConfig: {
      pin: true,
    },
    cache: {
      enable: false,
    },
    after: () => {
      if (vditor && doc && doc.content) {
        vditor.setValue(doc.content);
      }
    },
    input: (val) => {
      const newDoc = Automerge.change(doc!, (doc) => {
        doc.content = val;
      });
      doc = newDoc;
      let binary = Automerge.save(newDoc); // 存储
      // localforage.setItem(docId, binary).catch((err) => {
      //   console.log("store local err: ", err);
      // });
      patcher.patchDocument(docId, binary.toString());
      // updateDoc(docId, binary.toString());
    },
  });
};
</script>

<style>
@import url(vditor/dist/index.css);

.editor-container {
  width: 100%;
  height: 100%;
  padding: 0;
  margin: 0;
}

#vditor {
  height: calc(100vh - 93px) !important;
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
