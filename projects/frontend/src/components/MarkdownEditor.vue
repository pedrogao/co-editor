<template>
  <div class="container">
    <div id="vditor"></div>
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
    height: 660,
    width: 1200,
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

.container {
  /* text-align: center; */
}
</style>
