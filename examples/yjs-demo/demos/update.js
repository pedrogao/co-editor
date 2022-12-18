import * as Y from "yjs";

const doc1 = new Y.Doc();
const doc2 = new Y.Doc();
const text1 = doc1.getText();
const text2 = doc2.getText();

doc1.on("update", (update) => Y.applyUpdate(doc2, update));
doc2.on("update", (update) => Y.applyUpdate(doc1, update));

text1.insert(0, "Edwards");
text2.insert(0, "Wilson");

console.log(text1.toJSON());
console.log(text2.toJSON());
