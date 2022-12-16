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

type callback = (
  message: common.DocumentMessage<common.InnerDocumentMessage>
) => void;

//
// websocket api
//
interface ServerToClientEvents {
  patch_document: callback;
  fetch_document: callback;
}

interface ClientToServerEvents {
  patch_document: callback;
  fetch_document: callback;
}

export class DocumentPatcher {
  private socket: Socket<ServerToClientEvents, ClientToServerEvents>;
  constructor(url: string) {
    this.socket = io(url, {
      withCredentials: true,
      extraHeaders: {
        "Access-Control-Allow-Origin": "*",
        "Access-Control-Allow-Credentials": "true",
      },
    });
  }

  start() {
    this.socket.on("connect", () => {
      this.socket.on("patch_document", (message) => {
        if (message.data.error) {
          console.error("patch document err: ", message.data.error);
        } else {
          console.log("patch document successful: ", message.data.id);
        }
      });

      this.socket.on("fetch_document", (message) => {
        if (message.data.error) {
          console.error("fetch document err: ", message.data.error);
        } else {
          const { id, content } = message.data;
          console.log("fetch document successful: ", id);
          // TODO merge content
          console.log(content);
        }
      });
    });
  }

  patchDocument(id: string, content: string) {
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
