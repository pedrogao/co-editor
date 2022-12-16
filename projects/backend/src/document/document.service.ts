import { Injectable } from '@nestjs/common';
import { defaultDocumentStore, Document } from '../store';

@Injectable()
export class DocumentService {
  create(): string {
    return defaultDocumentStore.create().id;
  }

  query(id: string): Document | undefined {
    return defaultDocumentStore.query(id);
  }

  update(id: string, content: string) {
    return defaultDocumentStore.update(id, {
      content,
    });
  }
}
