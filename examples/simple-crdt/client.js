import { io } from "socket.io-client";
import { State } from "./lib/dt";

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
