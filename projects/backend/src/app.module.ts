import { Module, NestModule, MiddlewareConsumer } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { DocumentController } from './document/document.controller';
import { DocumentService } from './document/document.service';
import { LoggerMiddleware } from './middleware/logger.middleware';
import { WsDocumentGateway } from './gateway/ws.gateway';

@Module({
  imports: [],
  controllers: [AppController, DocumentController],
  providers: [AppService, DocumentService, WsDocumentGateway],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(LoggerMiddleware).forRoutes('*');
  }
}
