import { nanoid } from 'nanoid';

export const generate = () => {
  return nanoid(10).replace('_', '').replace('-', '');
};
