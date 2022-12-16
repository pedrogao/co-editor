import { Document, DocumentStore, UpdateDocumentArgs } from './document.store';
import { nanoid } from 'nanoid';

export class MemoryDocumentStore implements DocumentStore {
  private docMap = new Map<string, Document>();

  create(): Document {
    const id = nanoid(10);
    this.docMap.set(id, { id, content: '' });
    return this.docMap.get(id);
  }

  query(id: string): Document | undefined {
    return this.docMap.get(id);
  }

  update(id: string, args: UpdateDocumentArgs): boolean {
    let doc: Document;
    if ((doc = this.docMap.get(id)) === undefined) {
      return false;
    }
    this.docMap.set(id, { ...doc, ...args });
  }

  delete(id: string): boolean {
    if (!this.docMap.has(id)) {
      return false;
    }
    this.docMap.delete(id);
  }
}
