import { State, loadActionsInfo } from "./dt";
import { socket } from "./socket";

export class ClientState extends State {
  addAction(action) {
    super.addAction(action);
    const actionsInfo = JSON.stringify([action]);
    socket.emit("client-actions", actionsInfo);
  }

  serverAddAction(action) {
    this.addAction(action);
  }
}

let clientState = null;

socket.on("client-init", (info) => {
  const { uid, timestamp, actionsInfo } = info;
  clientState = new ClientState(timestamp, uid);
  const actions = loadActionsInfo(actionsInfo);
  for (const action of actions) {
    clientState.serverAddAction(action);
  }
});

export const clientInit = () => {
  socket.emit("client-connect");
};

export const getClientState = () => {
  return clientState;
};
