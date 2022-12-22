import { Logger } from '@nestjs/common';

export const logger = new Logger('Common', {
  timestamp: true,
});
