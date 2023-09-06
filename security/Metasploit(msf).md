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

## Metasploit 渗透测试之信息收集

### 基于 tcp 协议收集主机信息

- 用 Metasploit 中的 nmap 和 arp_sweep 收集主机信息
```shell
db_nmap -sV 192.168.1.1

# ARP 扫描
use auxiliary/scanner/discovery/arp_sweep
# 查看一下模块需要配置哪些参数
show options
# 配置 RHOSTS（扫描的目标网络）即可
set RHOSTS 192.168.1.0/24
# SHOST 和 SMAC 是伪造源 IP 和 MAC 地址使用的
# 配置线程数
set THREADS 30
run
#退出一下
back
```
- 使用半连接方式扫描 TCP 端口
```shell
search portscan
use auxiliary/scanner/portscan/syn
# 查看配置项
show options
# 设置扫描的目标
set RHOSTS 192.168.1.1
# 置端口范围使用逗号隔开
set PORTS 80
# 设置线程数
set THREADS 20
run
```
- 使用 auxiliary /sniffer 下的 psnuffle 模块进行密码嗅探
```shell
search psnuffle
use auxiliary/sniffer/psnuffle
# 查看 psnuffle 的模块作用：
info
#这个 psnuffle 模块可以像以前的 dsniff 命令一样，去嗅探密码，只支持 pop3、imap、ftp、HTTP GET 协议。
show options
run
```

### 基于 SNMP 协议收集主机信息
```shell
# 实战-使用 snmp_enum 模块通过 snmp 协议扫描目标服务器信息
use auxiliary/scanner/snmp/snmp_enum
show options
set RHOSTS 192.168.1.180
run
```

### 基于 SMB 协议收集信息
- 使用 smb_version 基于 SMB 协议扫描版本号
```shell 
use auxiliary/scanner/smb/smb_version
# 设置扫描目标，注意多个目标使用逗号+空格隔开
show options
set RHOSTS 192.168.1.56, 192.168.1.180
# 注： 192.168.1.56 后面的逗号和 192.168.1.180 之间是有空格的
run
```
- 使用 smb_enumshares 基于 SMB 协议扫共享文件（账号、密码）
```shell
use auxiliary/scanner/smb/smb_enumshares
show options
# 扫描 192.168.1.53 到 192.168.1.60 的机器
set RHOSTS 192.168.1.53-60
# 如果你不配置用户，就扫描不到信息。配置一下用户信息，我这里用户是默认的管理员用户。
set SMBUser administrator
set SMBPass 123456
run
```
- 使用 smb_lookupsid 扫描系统用户信息（SID 枚举）
```shell
# 注：SID 是 Windows 中每一个用户的 ID，更改用户名 SID 也是不会改变的。
use auxiliary/scanner/smb/smb_lookupsid
show options
set RHOSTS 192.168.1.56
set SMBUser administrator
set SMBPass 123456
run
```

### 基于 SSH 协议收集信息
- 查看 ssh 服务的版本信息
```shell
use auxiliary/scanner/ssh/ssh_version
show options
set RHOSTS 192.168.1.180
run
```
-  对 SSH 暴力破解
```shell
use auxiliary/scanner/ssh/ssh_login
show options
set RHOSTS 192.168.1.180
# 设置字典文件默认的字典文件是不满足实际需求的后期我们使用更强大的字典文件。
set USERPASS_FILE /usr/*/root_userpass.txt
run
# ssh 暴力破解成功后，会自动建立与目标机的连接
# 查看已建立的连接
sessions
```

### 基于 FTP 协议收集信息
- 查看 ftp 服务的版本信息
```shell
# 加载 ftp 服务版本扫描模块
use auxiliary/scanner/ftp/ftp_version
# 查看设置参数
show options
# 设置目标 IP，可以设置多个
set RHOSTS 192.168.1.180
# 执行扫描，输入 exploit 或 run
run
# 扫描出 ftp 服务的版本号，我们可以尝试搜索版本号，看看有没有可以利用的模块
search 2.3.4
```
- ftp 匿名登录扫描
```shell
use auxiliary/scanner/ftp/anonymous
show options
set RHOSTS 192.168.1.180
run
```
- ftp 暴力破解
```shell 
use auxiliary/scanner/ftp/ftp_login
show options
set RHOSTS 192.168.1.180
# 设置字典文件为默认的字典文件是不满足实际需求的，后期我们使用更强大的字典文件。
set USERPASS_FILE /usr/*/root_userpass.txt
run
```
