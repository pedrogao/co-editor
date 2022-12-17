<template>
  <div class="editor-container">
    <div id="vditor"></div>
    <div class="status">
      <div class="status-left">
        <span>字数：
          <strong>2</strong>
        </span>
        <span>行数：
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
import * as localforage from "localforage";
import { createDoc, DocumentPatcher, queryDoc } from "../api/document";
import { defaultWSURL } from "../config";
import { marshal, unmarshal } from "../marshaler";

export interface Document {
  content: string;
}

const props = defineProps<{
  id?: string;
}>();

onMounted(async () => {
  console.log("props: ", props);
  await init(props.id);
});

let vditor: Vditor | null = null;

const init = async (id: string | undefined) => {
  let doc: Automerge.Doc<Document> | null = null,
    docId: string | null = null,
    content: string | null = null;
  if (id && id !== "") {
    const { id: id1, content: content1 } = await queryDoc(id);
    docId = id1;
    content = content1;
  } else {
    const { id: id1, content: content1 } = await createDoc();
    docId = id1;
    content = content1;
  }
  changeUrl(docId);
  try {
    doc = Automerge.load(unmarshal(content!));
  } catch (error) {
    console.error(error);
  }

  const patcher: DocumentPatcher = new DocumentPatcher(
    defaultWSURL,
    patchCallback,
    fetchCallback
  );
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
        // vditor.focus();
      }
    },
    input: (val) => {
      const newDoc = Automerge.change(doc!, (doc) => {
        doc.content = val;
      });
      doc = newDoc;
      let binary = Automerge.save(newDoc); // 存储
      patcher.patchDocument(docId!, marshal(binary));
    },
  });
};

const fresh = () => {
  init(props.id);
};

const fetchCallback = (message: string) => {
  const doc = Automerge.load<Document>(unmarshal(message));
  if (vditor && doc && doc.content) {
    console.log("fetch: ", doc.content);
    vditor.setValue(doc.content);
  }
};

const patchCallback = (message: string) => {
  console.log(message);
};

const loadFromLocal = async (docId: string) => {
  return localforage.getItem(docId);
};

const saveToocal = async (docId: string, content: string) => {
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
