import { Char } from './types';
import { generateId } from './id';

export const generateSite = (siteId: string = '1') => {
  const startId = generateId();
  const endId = generateId();
  const start: Char = {
    id: startId,
    charId: {
      siteId,
      clock: 0,
    },
    value: '',
    visible: false,
    prevId: null,
    nextId: endId,
  };
  const end: Char = {
    id: endId,
    charId: {
      siteId,
      clock: 1,
    },
    value: '',
    visible: false,
    prevId: startId,
    nextId: null,
  };
  return { start, end, siteId };
};

export const generateChar = (
  siteId: string,
  clock: number,
  alpha: string,
  prevId: string,
  nextId: string
): Char => ({
  id: generateId(),
  charId: {
    siteId,
    clock,
  },
  value: alpha,
  visible: true,
  prevId,
  nextId,
});
