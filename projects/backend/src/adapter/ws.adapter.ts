import * as WebSocket from 'ws';
import { WebSocketAdapter, INestApplicationContext } from '@nestjs/common';
import { MessageMappingProperties } from '@nestjs/websockets';
import { Observable, fromEvent, EMPTY } from 'rxjs';
import { mergeMap, filter } from 'rxjs/operators';
import { logger } from '../logger';
import { DocumentMessage } from 'common/src/index';

export class WsAdapter implements WebSocketAdapter {
  constructor(private app: INestApplicationContext) {}

  create(port: number, options: any = {}): any {
    logger.info(`ws create from port: ${port}`);
    return new WebSocket.Server({ port, ...options });
  }

  bindClientConnect(server: WebSocket, callback: () => void) {
    logger.info('ws bind client connect');
    server.on('connection', callback);
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  bindClientDisconnect(client: WebSocket, callback: () => void) {
    logger.info('client disconnect');
  }

  bindMessageHandlers(
    client: WebSocket,
    handlers: MessageMappingProperties[],
    process: (data: any) => Observable<any>,
  ) {
    logger.info('new connection from client');

    fromEvent(client, 'message')
      .pipe(
        mergeMap((data) =>
          this.bindMessageHandler(client, data, handlers, process),
        ),
        filter((result) => result),
      )
      .subscribe((response) => client.send(JSON.stringify(response)));
  }

  bindMessageHandler(
    _client: WebSocket,
    buffer,
    handlers: MessageMappingProperties[],
    process: (data: any) => Observable<any>,
  ): Observable<any> {
    let message: DocumentMessage<string>;
    try {
      message = JSON.parse(buffer.data);
    } catch (error) {
      logger.error('parse ws data err: ', error);
      return EMPTY;
    }

    const messageHandler = handlers.find(
      (handler) => handler.message === message.event, // 匹配
    );
    if (!messageHandler) {
      logger.error('message miss match: ', message);
      return EMPTY;
    }
    logger.info('handle message: ', message);
    return process(messageHandler.callback(message.data));
  }

  close(server: WebSocket) {
    logger.info('ws server close');
    server.close();
  }
}
