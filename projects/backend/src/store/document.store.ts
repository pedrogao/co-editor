export class Document {
  content: string; // 文档内容
  id: string; // 文档id
  // creator: string;
  // comments: string[];
  // so many others...
}

export interface UpdateDocumentArgs {
  content: string;
}

export interface DocumentStore {
  create(): Document;

  query(id: string): Document | undefined;

  update(id: string, args: UpdateDocumentArgs): boolean;

  delete(id: string): boolean;
}
