import Quill from "quill";
import QuillCursors from "quill-cursors";
import * as Y from "yjs";
import { QuillBinding } from "y-quill";
import { WebrtcProvider } from "y-webrtc";

Quill.register("modules/cursors", QuillCursors);

const quill = new Quill(document.querySelector("#editor"), {
  modules: {
    cursors: true,
    toolbar: [
      // adding some basic Quill content features
      [{ header: [1, 2, false] }],
      ["bold", "italic", "underline"],
      ["image", "code-block"],
    ],
    history: {
      // Local undo shouldn't undo changes
      // from remote users
      userOnly: true,
    },
  },
  placeholder: "Type what you want...",
  theme: "snow",
});

const ydoc = new Y.Doc();
const ytext = ydoc.getText("quill");

const provider = new WebrtcProvider("quill-demo-room", ydoc);
const binding = new QuillBinding(ytext, quill, provider.awareness);

// Remove the selection when the iframe is blurred
window.addEventListener("blur", () => {
  quill.blur();
});
