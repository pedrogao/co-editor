import { State, IdManager } from "./dt";

export class ServerState extends State {
  constructor(...args) {
    super(...args);
    this.uidManager = new IdManager();
  }
}

export const serverState = new ServerState(0, 0);
