import { HttpAdapterHost, NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { AllExceptionsFilter } from './exception/exception.filter';
// import { WsAdapter } from './adapter/ws.adapter';
import { WsAdapter } from '@nestjs/platform-ws';

async function bootstrap() {
  const app = await NestFactory.create(AppModule, { cors: true });
  const { httpAdapter } = app.get(HttpAdapterHost);
  app.useGlobalFilters(new AllExceptionsFilter(httpAdapter as any));
  app.useWebSocketAdapter(new WsAdapter(app));
  await app.listen(3000);
}

process.on('warning', (e) => console.warn(e.stack));

bootstrap();
