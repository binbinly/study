import { Module } from "@nestjs/common";
import { MessageController } from "./message.controller";
import { ChatModule } from "../chat/chat.module";

@Module({
  imports: [ChatModule],
  controllers: [MessageController],
})
export class MessageModule {}
