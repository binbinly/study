import {
  SubscribeMessage,
  WebSocketGateway,
  WebSocketServer,
} from "@nestjs/websockets";
import { Server } from "ws";
import { ChatService } from "./chat.service";
import * as url from "url";

const port = parseInt(process.env.WS_PORT) || 3010
@WebSocketGateway(port)
export class ChatGateway {
  ws = new Map();

  constructor(private readonly ChatService: ChatService) {
    console.log("pid", process.pid);
  }

  @WebSocketServer()
  server: Server;

  // 网关初始化狗子
  async afterInit(server: Server) {
    console.log("server init");
  }

  // socket连接钩子
  async handleConnection(socket: any, req) {
    const params = url.parse(req.url, true);
    if (params) {
      const token = params.query.token;
      if (token) {
        const res = await this.ChatService.online(token.toString());
        if (res.user_id) {
          socket.id = res.user_id;
          this.ws.set(res.user_id, socket);
          return;
        }
        socket.send(this.failClose(res.msg));
        return socket.close();
      }
    }
    socket.send(this.failClose("非法操作"));
    return socket.close();
  }

  // socket 连接断开
  async handleDisconnect(socket: any) {
    console.log("disconnect", socket.id);
    if (this.ws.has(socket.id)) {
      this.ws.delete(socket.id);
    }
    this.ChatService.offline(socket.id);
  }

  @SubscribeMessage("offline_message")
  async handleEvent(client: any, data: string): Promise<any> {
    const list = await this.ChatService.getMessages(client.id);
    list.forEach((item) => {
      client.send(item);
    });
  }

  sendMsg(user_id, msg: string) {
    console.log("user_id", user_id, msg);
    if (this.ws.has(user_id)) {
      const socket = this.ws.get(user_id);
      try {
        socket.send(msg);
      } catch (error) {
        // 发送失败，保存为历史消息
        this.ChatService.saveMessage(user_id, msg);
      }
    } else {
      // 不在线，保存为历史消息
      this.ChatService.saveMessage(user_id, msg);
    }
  }

  failMsg(msg: string) {
    return JSON.stringify({ event: "fail", data: msg });
  }

  // 客户端不允许自动重连
  failClose(msg: string) {
    return JSON.stringify({ event: "close", data: msg });
  }
}
