import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpStatus,
} from '@nestjs/common';
import { HttpAdapterHost } from '@nestjs/core';
import { APIResponse } from 'common';
import { logger } from '../logger';

@Catch()
export class AllExceptionsFilter implements ExceptionFilter {
  constructor(private readonly httpAdapterHost: HttpAdapterHost) {}

  catch(exception: unknown, host: ArgumentsHost): void {
    const { httpAdapter } = this.httpAdapterHost;
    const ctx = host.switchToHttp();
    logger.error(
      `request: ${ctx.getRequest()}, resp: ${ctx.getResponse()}, error: ${exception}`,
    );
    const resp: APIResponse<string> = {
      status: 1,
      message: JSON.stringify(exception),
    };

    httpAdapter.reply(ctx.getResponse(), resp, HttpStatus.OK);
  }
}
