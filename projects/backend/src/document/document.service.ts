import { Injectable } from '@nestjs/common';
import { defaultDocumentStore, Document } from '../store';

@Injectable()
export class DocumentService {
  create(content: string): string {
    return defaultDocumentStore.create(content).id;
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
