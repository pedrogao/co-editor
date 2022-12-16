export const defaultAPIURL = "http://localhost:3000/api";
export const defaultWSURL = "ws://localhost:3001";

export const PATCH_DOCUMENT_EVENT: string = "patch_document";
export const FETCH_DOCUMENT_EVENT: string = "fetch_document";
export const EVENT_OK: string = "ok";
export const EVENT_ERR: string = "err";

// websocket消息
export type DocumentMessage<T> = {
  event: string;
  data: T;
};

export type InnerDocumentMessage = {
  id?: string;
  content?: string;
  error?: string;
};

// API响应
export type APIResponse<T> = {
  status: number; // 0 => ok
  message?: string;
  data?: T; // data
};
