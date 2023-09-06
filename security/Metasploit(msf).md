## Metasploit 渗透测试框架的基本使用
```shell 
# Kail 终端命令启动
msfconsole
# 注：Metasploit 程序需要使用 Postgresql 数据库, 快捷启动命令 msfdb run 可以同时启动 postgresql 数据库和 msf
```

### 核心命令
- banner 显示一个 metasploit 横幅
- cd 更改当前的工作目录
- color 切换颜色
- connect 连接与主机通信
- exit 退出控制台
- get 获取特定于上下文的变量的值
- getg 获取全局变量的值
- grep 另一个命令的输出 如： grep creds help
- help 帮助菜单
- history 显示命令历史
- irb 进入 irb 脚本模式
- load 加载一个框架插件
- quit 退出控制台
- route 通过会话路由流量
- save 保存活动的数据存储
- sessions 转储会话列表并显示有关会话的信息
- set 将特定于上下文的变量设置为一个值
- setg 将全局变量设置为一个值
- sleep 在指定的秒数内不做任何事情
- spool 将控制台输出写入文件以及屏幕
- threads 线程查看和操作后台线程
- unload 卸载框架插件
- unset 取消设置一个或多个特定于上下文的变量
- unsetg 取消设置一个或多个全局变量
- version 显示框架和控制台库版本号

### 模块命令
- advanced 显示一个或多个模块的高级选项
- back 从当前上下文返回
- edit 使用首选编辑器编辑当前模块
- info 显示有关一个或多个模块的信息
- loadpath 路径从路径搜索并加载模块
- options 显示全局选项或一个或多个模块
- popm 将最新的模块从堆栈中弹出并使其处于活动状态
- previous 将之前加载的模块设置为当前模块
- pushm 将活动或模块列表推入模块堆栈
- reload_all 从所有定义的模块路径重新加载所有模块
- search 搜索模块名称和描述
- show 显示给定类型的模块或所有模块
- use 按名称选择模块

### 工作命令
- handler 作为作业启动负载处理程序
- jobs 显示和管理作业
- kill 杀死一个工作
- rename_job 重命名作业

### 资源脚本命令
- makerc 保存从开始到文件输入的命令
- resource 运行存储在文件中的命令

### 数据库后端命令
- db_connect 连接到现有的数据库
- db_disconnect 断开与当前数据库实例的连接
- db_export 导出包含数据库内容的文件
- db_import 导入扫描结果文件（文件类型将被自动检测）
- db_nmap 执行 nmap 并自动记录输出
- db_rebuild_cache 重建数据库存储的模块高速缓存
- db_status 显示当前的数据库状态
- hosts 列出数据库中的所有主机
- loot 列出数据库中的所有战利品
- notes 列出数据库中的所有笔记
- services 列出数据库中的所有服务
- vulns 列出数据库中的所有漏洞
- workspace 在数据库工作区之间切换

### 凭证后端命令
- creds 列出数据库中的所有凭据
