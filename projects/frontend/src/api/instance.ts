import axios from "axios";

const instance = axios.create({
  baseURL: "http://localhost:3000/api",
  timeout: 3000,
  headers: {
    "Access-Control-Allow-Origin": "*",
    "Access-Control-Allow-Credentials": true,
  },
});

export const post = instance.post;
export const put = instance.put;
export const get = instance.get;
