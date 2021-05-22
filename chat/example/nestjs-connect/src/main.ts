import { NestFactory } from "@nestjs/core";
import { AppModule } from "./app.module";
import { Transport } from "@nestjs/microservices";
import { WsAdapter } from "@nestjs/platform-ws";
import { join } from "path";
import { BOOT, IBoot } from "@nestcloud/common";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  app.useWebSocketAdapter(new WsAdapter(app));

  const boot = app.get<IBoot>(BOOT);
  await app.connectMicroservice({
    transport: Transport.GRPC,
    options: {
      url: `0.0.0.0:${boot.get("service.port")}`,
      package: "message",
      loader: {
        keepCase: true, // 非驼峰形式 默认驼峰
        longs: Number,
        defaults: true, // 输出对象上设置默认值，默认 false
      },
      protoPath: join(__dirname, "../proto/message.proto"),
    },
  });
  await app.startAllMicroservicesAsync();
  await app.listen(boot.get("http.port"));
  console.log(`Application is running on: ${await app.getUrl()}`);
}

bootstrap();
