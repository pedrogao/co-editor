import {
  ConnectedSocket,
  MessageBody,
  SubscribeMessage,
  WebSocketGateway,
} from '@nestjs/websockets';
import { logger } from '../logger';
import { DocumentService } from '../document/document.service';
import * as WebSocket from 'ws';
import {
  PATCH_DOCUMENT_EVENT,
  FETCH_DOCUMENT_EVENT,
  EVENT_ERR,
  DocumentMessage,
  InnerDocumentMessage,
} from 'common';

@WebSocketGateway(3001, { cors: true })
export class WsDocumentGateway {
  constructor(private readonly docService: DocumentService) {}

  @SubscribeMessage(PATCH_DOCUMENT_EVENT)
  patchDocument(
    @MessageBody() message: InnerDocumentMessage,
  ): DocumentMessage<InnerDocumentMessage> {
    logger.info('patch document: ', message);
    // 提交文档更新
    const resp: DocumentMessage<InnerDocumentMessage> = {
      event: PATCH_DOCUMENT_EVENT,
      data: {},
    };
    const { id, content } = message;
    if (!id) {
      resp.data.error = EVENT_ERR;
      return resp;
    }
    const doc = this.docService.query(id);
    if (!doc) {
      resp.data.error = EVENT_ERR;
      return resp;
    }
    if (!content) {
      return resp; // 啥也不做
    }

    const ok = this.docService.update(id, content);
    if (!ok) {
      resp.data.error = EVENT_ERR;
      return resp;
    }
    // TODO 引入 automerge
    // 更新完成
    return resp;
  }

  @SubscribeMessage(FETCH_DOCUMENT_EVENT)
  fetchDocument(
    @MessageBody() message: InnerDocumentMessage,
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    @ConnectedSocket() client: WebSocket,
  ): DocumentMessage<InnerDocumentMessage> {
    logger.info('fetch document: ', message);
    // 获取最新文档
    // client.send(JSON.stringify({ event: 'tmp', data: '这里是个临时信息' }));
    // TODO 返回最新的文档数据
    const resp: DocumentMessage<InnerDocumentMessage> = {
      event: FETCH_DOCUMENT_EVENT,
      data: {},
    };
    const { id } = message;
    if (!id) {
      resp.data.error = EVENT_ERR;
      return resp;
    }
    const doc = this.docService.query(id);
    if (!doc) {
      resp.data.error = EVENT_ERR;
      return resp;
    }
    resp.data.content = doc.content;
    return resp;
  }
}
