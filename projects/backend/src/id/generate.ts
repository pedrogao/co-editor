import { nanoid } from 'nanoid';

export const generate = () => {
  return nanoid(15).replace('_', '').replace('-', '');
};
