import { io, Socket } from "socket.io-client";
import { post, put, get } from "./instance";
import * as common from "../config";

export const createDoc = async () => {
  const resp = await post("/document");
  const id: string = resp.data.data;
  console.log("create document: ", id);
  return id;
};

export const updateDoc = async (id: string, content: string) => {
  try {
    await put(`/document/${id}`, {
      content,
    });
  } catch (error) {
    console.log("update doc err: ", error);
  }
};

export const queryDoc = async (id: string) => {
  try {
    const resp = await get(`/document/${id}`);
    return resp.data.content as string;
  } catch (error) {
    console.log("update doc err: ", error);
  }
};

type clientCallback = (
  message: common.DocumentMessage<common.InnerDocumentMessage>
) => void;
type serverCallback = (message: common.InnerDocumentMessage) => void;

//
// websocket api
//
interface ServerToClientEvents {
  patch_document: serverCallback;
  fetch_document: serverCallback;
}

interface ClientToServerEvents {
  patch_document: clientCallback;
  fetch_document: clientCallback;
}

export class DocumentPatcher {
  private socket: Socket<ServerToClientEvents, ClientToServerEvents>;

  constructor(url: string) {
    this.socket = io(url, {
      transports: ["websocket"],
      autoConnect: false,
      withCredentials: true,
      extraHeaders: {},
    });
    this.socket.on("connect", () => {
      console.log("new connection: ", this.socket.id);
    });
    this.socket.on("disconnect", () => {
      console.log("disconnect");
    });

    this.socket.on("patch_document", (message) => {
      console.log("patch_document message: ", message);
      if (message.error) {
        console.error("patch document err: ", message.error);
      } else {
        console.log("patch document successful: ", message.id);
      }
    });
    this.socket.on("fetch_document", (message) => {
      console.log("fetch_document message: ", message);
      if (message.error) {
        console.error("fetch document err: ", message.error);
      } else {
        const { id, content } = message;
        console.log("fetch document successful: ", id);
        // TODO merge content
        console.log(content);
      }
    });
  }

  start() {
    this.socket.connect();
  }

  patchDocument(id: string, content: string) {
    // console.log("patch", id, content);
    this.socket.emit("patch_document", {
      event: common.PATCH_DOCUMENT_EVENT,
      data: {
        id: id,
        content: content,
      },
    });
  }

  fetchDocument(id: string) {
    this.socket.emit("fetch_document", {
      event: common.FETCH_DOCUMENT_EVENT,
      data: {
        id: id,
      },
    });
  }
}
