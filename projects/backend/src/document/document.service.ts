import { Injectable } from '@nestjs/common';
import { nanoid } from 'nanoid';
import { Doc } from './dto';

@Injectable()
export class DocumentService {
  private docMap = new Map<string, Doc>();

  create(): string {
    const id = nanoid(10);
    this.docMap.set(id, { id, content: '' });
    return id;
  }

  query(id: string): Doc | undefined {
    return this.docMap.get(id);
  }

  update(id: string, content: string) {
    this.docMap.set(id, { id, content });
  }
}
