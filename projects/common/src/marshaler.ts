export const marshal = (arr: Uint8Array): string => {
  return Buffer.from(arr).toString("base64");
};

export const unmarshal = (s: string): Uint8Array => {
  const buffer = Buffer.from(s, "base64");
  return new Uint8Array(buffer);
};
