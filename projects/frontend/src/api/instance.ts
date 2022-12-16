import axios from "axios";
import { defaultAPIURL } from "../config";

//
// http api
//
const instance = axios.create({
  baseURL: defaultAPIURL,
  timeout: 3000,
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Credentials": true,
  },
});

export const post = instance.post;
export const put = instance.put;
export const get = instance.get;
