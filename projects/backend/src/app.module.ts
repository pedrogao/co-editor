import { Module, NestModule, MiddlewareConsumer } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { DocumentController } from './document/document.controller';
import { DocumentService } from './document/document.service';
import { LoggerMiddleware } from './middleware/logger.middleware';
import { YjsGateway } from './gateway/yjs.gateway';

@Module({
  imports: [],
  controllers: [AppController, DocumentController],
  providers: [AppService, DocumentService, YjsGateway],
})
export class AppModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(LoggerMiddleware).forRoutes('*');
  }
}
