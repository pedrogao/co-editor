import { EventEmitter } from "events";

export class Socket extends EventEmitter {}

export const socket = new Socket();
