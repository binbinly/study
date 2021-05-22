import { Module } from "@nestjs/common";
import { MessageModule } from "./message/message.module";
import { resolve } from "path";
import { BootModule } from "@nestcloud/boot";
import { BOOT, CONSUL } from "@nestcloud/common";
import { ConsulModule } from "@nestcloud/consul";
import { ServiceModule } from "@nestcloud/service";
import { ConfigModule } from '@nestjs/config';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MessageModule,
    BootModule.forRoot({ filePath: resolve(__dirname, '../conf/config.yaml') }),
    ConsulModule.forRootAsync({ inject: [BOOT] }),
    ServiceModule.forRootAsync({ inject: [BOOT, CONSUL] }),
  ],
})
export class AppModule {}
