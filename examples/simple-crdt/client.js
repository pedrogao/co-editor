import { io } from "socket.io-client";
import { draw, clear } from "./tree";
import { State, loadActionsInfo } from "./lib/dt";

export const socket = io("http://localhost:8888", {});

export class ClientState extends State {
  addAction(action) {
    super.addAction(action);
    const actionsInfo = JSON.stringify([action]);
    socket.emit("client-actions", actionsInfo);
  }

  serverAddAction(action) {
    super.addAction(action);
  }
}

// App
import { createApp } from "vue";

createApp({
  created() {
    socket.on("server-actions", (actionsInfo) => {
      const actions = loadActionsInfo(actionsInfo);
      for (let action of actions) {
        this.clientState.serverAddAction(action);
      }
      this.updateMessage();
      draw(this.clientState.tree);
    });
    socket.on("client-init", (info) => {
      console.log("client -init");
      const { uid, timeStamp, actionsInfo } = info;
      this.clientState = new ClientState(timeStamp, uid);
      window.clientState = this.clientState;
      const actions = loadActionsInfo(actionsInfo);
      for (let action of actions) {
        this.clientState.serverAddAction(action);
      }
      this.updateMessage();
      draw(this.clientState.tree);
    });
  },
  data() {
    return {
      start: 0,
      end: 0,
      clientState: null,
      message: "",
    };
  },
  methods: {
    updateMessage() {
      const input = document.getElementsByClassName("input-area")[0];
      const start = input && input.selectionStart;
      const end = input && input.selectionEnd;
      this.message = this.clientState.toString();
      // this.$nextTick(() => {
      //   input.setSelectionRange(start, end);
      // });
    },
    draw() {
      clear();
      draw(this.clientState.tree);
    },
    clear() {
      clear();
    },
    onChange(value) {
      this.clientState.update(this.message);
      draw(this.clientState.tree);
      this.updateMessage();
    },
  },
  render() {},
}).mount("#app");
