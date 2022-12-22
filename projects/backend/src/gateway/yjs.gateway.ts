/* eslint-disable @typescript-eslint/no-empty-function */
import {
  WebSocketGateway,
  WebSocketServer,
  OnGatewayConnection,
  OnGatewayDisconnect,
} from '@nestjs/websockets';
import { logger } from '../logger';
import { Server } from 'ws';
import { setupWSConnection } from 'y-websocket/bin/utils';

@WebSocketGateway(3001, {
  cors: {
    origin: '*',
  },
})
export class YjsGateway implements OnGatewayConnection, OnGatewayDisconnect {
  constructor() {}

  @WebSocketServer()
  server: Server;

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  handleConnection(client: any): void {
    logger.log('new connection');
    this.server.on('connection', setupWSConnection);
  }

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  handleDisconnect(client: any): void {
    logger.log('disconnection');
  }
}
