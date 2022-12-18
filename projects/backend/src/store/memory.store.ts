import { Document, DocumentStore, UpdateDocumentArgs } from './document.store';
import { generate } from '../id';

export class MemoryDocumentStore implements DocumentStore {
  private docMap = new Map<string, Document>();

  create(content: string): Document {
    const id = generate();
    this.docMap.set(id, { id, content });
    return this.docMap.get(id);
  }

  query(id: string): Document | undefined {
    return this.docMap.get(id);
  }

  update(id: string, args: UpdateDocumentArgs): boolean {
    this.docMap.set(id, { id, content: args.content });
    return true;
  }

  delete(id: string): boolean {
    if (!this.docMap.has(id)) {
      return false;
    }
    this.docMap.delete(id);
  }
}
