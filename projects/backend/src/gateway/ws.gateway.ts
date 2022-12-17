import {
  ConnectedSocket,
  MessageBody,
  OnGatewayConnection,
  OnGatewayDisconnect,
  OnGatewayInit,
  SubscribeMessage,
  WebSocketGateway,
} from '@nestjs/websockets';
import * as Automerge from '@automerge/automerge';
import { logger } from '../logger';
import { Document } from '../store';
import { DocumentService } from '../document/document.service';
import * as WebSocket from 'ws';
import {
  PATCH_DOCUMENT_EVENT,
  FETCH_DOCUMENT_EVENT,
  DocumentMessage,
  InnerDocumentMessage,
  marshal,
  unmarshal,
} from 'common';

@WebSocketGateway(3001, {
  cors: {
    origin: '*',
  },
  transports: ['websocket'],
})
export class WsDocumentGateway
  implements OnGatewayConnection, OnGatewayDisconnect, OnGatewayInit
{
  private wsClients = [];

  constructor(private readonly docService: DocumentService) {}

  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  afterInit(server: any) {
    console.log('after init');
  }

  handleDisconnect(client: any) {
    console.log('dis client: ', client.id);
    for (let i = 0; i < this.wsClients.length; i++) {
      if (this.wsClients[i].id === client.id) {
        this.wsClients.splice(i, 1);
        break;
      }
    }
  }

  handleConnection(client: any, ...args: any[]) {
    console.log('new client: ', client.id);
    this.wsClients.push(client);
  }

  private broadcast(event: string, message: any) {
    const broadCastMessage = JSON.stringify(message);
    for (const c of this.wsClients) {
      console.log('client: ', c);
      console.log('broadCastMessage: ', broadCastMessage);
      c.send(event, broadCastMessage);
    }
  }

  @SubscribeMessage(PATCH_DOCUMENT_EVENT)
  patchDocument(
    client: any,
    message: DocumentMessage<InnerDocumentMessage>,
  ): DocumentMessage<InnerDocumentMessage> {
    // 提交文档更新
    const resp: DocumentMessage<InnerDocumentMessage> = {
      event: FETCH_DOCUMENT_EVENT, // 触发 fetch
      data: {},
    };
    const { id, content } = message.data;
    if (!id) {
      resp.data.error = 'id不可为空';
      return resp;
    }
    const doc = this.docService.query(id);
    if (!doc) {
      resp.data.error = '文档不存在';
      return resp;
    }
    if (!content) {
      return resp; // 啥也不做
    }
    let local;
    if (doc.content && doc.content !== '') {
      local = Automerge.load<Document>(unmarshal(doc.content));
    } else {
      local = Automerge.init<Document>();
    }
    const remote = Automerge.load<Document>(unmarshal(content));
    const newDoc = Automerge.merge(local, remote);
    logger.info('patch document: ', newDoc);
    const binary = Automerge.save(newDoc);
    const newContent = marshal(binary);
    const ok = this.docService.update(id, newContent);
    if (!ok) {
      resp.data.error = '更新文档失败';
      return resp;
    }

    resp.data.content = newContent;
    console.log('client: ', client.broadcast.emit);
    client.broadcast.emit(FETCH_DOCUMENT_EVENT, resp.data);
    return resp;
  }

  @SubscribeMessage(FETCH_DOCUMENT_EVENT)
  fetchDocument(
    @MessageBody() message: DocumentMessage<InnerDocumentMessage>,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    @ConnectedSocket() client: WebSocket,
  ): DocumentMessage<InnerDocumentMessage> {
    logger.info('fetch document: ', message.event);
    // 获取最新文档
    // client.send(JSON.stringify({ event: 'tmp', data: '这里是个临时信息' }));
    const resp: DocumentMessage<InnerDocumentMessage> = {
      event: FETCH_DOCUMENT_EVENT,
      data: {},
    };
    const { id } = message.data;
    if (!id) {
      resp.data.error = 'id不可为空';
      return resp;
    }
    const doc = this.docService.query(id);
    if (!doc) {
      resp.data.error = '文档不存在';
      return resp;
    }
    resp.data.content = doc.content;
    return resp;
  }
}
