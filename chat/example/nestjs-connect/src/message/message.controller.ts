import { Controller } from "@nestjs/common";
import { GrpcMethod } from "@nestjs/microservices";
import {
  SendRequest,
  NotifyRequest,
  MomentRequest,
} from "./interfaces/message.interface";
import { ChatGateway } from "../chat/chat.gateway";

@Controller()
export class MessageController {
  constructor(private readonly wss: ChatGateway) {}

  @GrpcMethod("MessageService")
  sendChat(req: SendRequest): null {
    console.log("server chat", req);
    const data = { event: "message", data: req };
    this.wss.sendMsg(req.to_id, JSON.stringify(data));
    return;
  }

  @GrpcMethod("MessageService")
  sendRecall(req: SendRequest): null {
    console.log("server recall", req);
    const data = { event: "recall", data: req };
    this.wss.sendMsg(req.to_id, JSON.stringify(data));
    return;
  }

  @GrpcMethod("MessageService")
  sendNotify(req: NotifyRequest): null {
    console.log("server notify", req);
    const data = { event: "notify", data: req };
    this.wss.sendMsg(req.to_id, JSON.stringify(data));
    return;
  }

  @GrpcMethod("MessageService")
  sendMoment(req: MomentRequest): null {
    console.log("server moment", req);
    const data = { event: "moment", data: req };
    this.wss.sendMsg(req.to_id, JSON.stringify(data));
    return;
  }
}
