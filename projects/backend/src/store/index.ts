import { MemoryDocumentStore } from './memory.store';

export const defaultDocumentStore = new MemoryDocumentStore();
export * from './document.store';
