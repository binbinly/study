import { Injectable } from "@nestjs/common";
import { JwtService } from "@nestjs/jwt";
import { RedisService } from "nestjs-redis";

@Injectable()
export class ChatService {
  public client;
  // 上线用户前缀
  private readonly prefix = "online:";
  // 离线消息缓存前缀
  private readonly historyPrefix = "history:";

  constructor(
    private readonly redisService: RedisService,
    private readonly jwtService: JwtService
  ) {
    this.getClient();
  }

  async getClient() {
    this.client = await this.redisService.getClient();
  }

  // 用户上线
  async online(token: string): Promise<any> {
    let data
    try {
      data = this.jwtService.decode(token);
      console.log("data", token, data);
    } catch(err) {
      return { msg: "验证失败了，请重新登录" };
    }
    
    if (data["user_id"]) {
      if (!this.client) {
        await this.getClient();
      }
      const t = await this.client.get("user:" + data["user_id"]);
      if (t != token) {
        return { msg: "登录过期，请重新登录" };
      }
      await this.client.hset(this.prefix + data["user_id"], "pid", process.pid);
      await this.client.expire(this.prefix + data["user_id"], 86400);
      return { user_id: data["user_id"] };
    } else {
      return { msg: "验证失败，请重新登录" };
    }
  }

  // 用户下线
  async offline(user_id) {
    await this.client.del(this.prefix + user_id);
  }

  // 保存离线消息
  async saveMessage(user_id, msg: string) {
    await this.client.rpush(this.historyPrefix + user_id, msg);
  }

  // 获取所有离线消息
  async getMessages(user_id): Promise<any> {
    const list = await this.client.lrange(this.historyPrefix + user_id, 0, -1);
    // 清除离线消息
    await this.client.del(this.historyPrefix + user_id);
    return list;
  }

  // 分布式场景下: 处理订阅当前进程消息
  async subscribe(socket: any) {
    const clientId = process.pid + "";
    await this.client.subscribe(clientId as string, (err) => {
      if (err) {
        console.log(`订阅到频道 ${clientId} 失败`, err.message);
        return;
      }
      console.log(`已订阅到频道 ${clientId}`);
    });

    this.client.on("message", (channel, message) => {
      socket.send(message);
    });
  }

  // 分布式场景下: 发布消息到指定进程
  async publish(clientId: string, message: string) {
    await this.client.publish(clientId, message);
  }
}
