import { Module } from "@nestjs/common";
import { ChatGateway } from "./chat.gateway";
import { ChatService } from "./chat.service";
import { JwtModule } from "@nestjs/jwt";
import { RedisModule } from "nestjs-redis";
import { Boot } from '@nestcloud/boot';
import { BOOT } from "@nestcloud/common";

@Module({
  providers: [ChatGateway, ChatService],
  exports: [ChatGateway],
  imports: [
    JwtModule.registerAsync({
      useFactory: (config: Boot) => ({
        secret: config.get('app.key'),
      }),
      inject: [BOOT],
    }),
    RedisModule.forRootAsync({
      useFactory: (config: Boot) => ({
        host: config.get('redis.host', 'localhost'),
        port: config.get('redis.port', 6379),
        db: config.get('redis.db', 0),
        password: config.get('redis.password', ''),
      }),
      inject: [BOOT],
    }),
  ],
})
export class ChatModule {}
