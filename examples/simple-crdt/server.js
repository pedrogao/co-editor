import * as express from "express";
import cors from "cors";
import * as http from "http";
import * as path from "path";
import { Server } from "socket.io";

import { loadActionsInfo } from "./lib/dt";
import { serverState } from "./lib/server";

const app = express.Router();
const server = http.createServer(app);
const socket = new Server(server, {
  cors: {
    origin: "*",
  },
});

app.use(
  cors({
    origin: "*",
  })
);
app.use(express.static(path.join(__dirname, "public")));

const allActions = [];
socket.on("connection", (socket) => {
  console.log("connection: ", socket.id);
  const uid = serverState.uidManager.getId();
  const timestamp = serverState.timestampManager.getTimestamp();
  const actionsInfo = JSON.stringify(allActions);
  socket.emit("client-init", { uid, timestamp, actionsInfo });

  socket.on("client-actions", (actionsInfo) => {
    const actions = loadActionsInfo(actionsInfo);
    allActions = allActions.concat(actions);
    for (const action of actions) {
      serverState.addAction(action);
      console.log("barodcast", action.toString());
      socket.broadcast.emit("server-actions", actionsInfo);
    }
  });

  socket.on("client-connect", (data) => {
    socket.broadcast.emit("new message", {});
  });

  socket.on("add user", (username) => {});
});

server.listen(8888, () => {
  console.log("listening on *:8888");
});
