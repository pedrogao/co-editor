<template>
  <div class="container">
    <div id="vditor"></div>
  </div>
</template>

<script lang="ts" setup>
import { onMounted } from "vue";
import Vditor from "vditor";
import * as Automerge from "@automerge/automerge";
import * as localforage from "localforage";
import { post, put, get } from "../api";

export interface Document {
  content: string;
}

let vditor: Vditor | null = null;
let doc: Automerge.Doc<Document> | null = null;
let docId: string;

onMounted(async () => {
  await init();
});

const init = async () => {
  docId = await createDoc();
  doc = Automerge.init();

  vditor = new Vditor("vditor", {
    height: 460,
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
      localforage.setItem(docId, binary).catch((err) => {
        console.log("store local err: ", err);
      });
      updateDoc(docId, binary.toString());
    },
  });
};

const createDoc = async () => {
  const resp = await post("/document");
  const id: string = resp.data.data;
  console.log("create document: ", id);
  return id;
};

const updateDoc = async (id: string, content: string) => {
  try {
    await put(`/document/${id}`, {
      content,
    });
  } catch (error) {
    console.log("update doc err: ", error);
  }
};

const queryDoc = async (id: string) => {
  try {
    const resp = await get(`/document/${id}`);
    return resp.data.content as string;
  } catch (error) {
    console.log("update doc err: ", error);
  }
};
</script>

<style>
@import url(vditor/dist/index.css);

.container {
  /* text-align: center; */
}
</style>
