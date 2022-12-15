export type APIResponse<T> = {
  status: number; // 0 => ok
  message?: string;
  data?: T; // data
};
