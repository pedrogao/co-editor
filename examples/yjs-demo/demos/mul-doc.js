import * as Y from "yjs";

// 新建一个文本
const doc = new Y.Doc();

// 新建 root map
const root = doc.getMap("root");

// 新建 point
const point = new Y.Map();
point.set("x", 0);
point.set("y", 0);

root.set("point", point);

const name = new Y.Text();
name.insert(0, "pedrogao");

root.set("name", name);

console.log(doc);
