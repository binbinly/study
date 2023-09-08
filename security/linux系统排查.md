## Linux 应急响应-溯源-系统日志排查

### 查看当前已经登录到系统的用户
```shell 
w #是一个命令行工具，它可以展示当前登录用户信息，并且每个用户正在做什么。它同时展示以下信息：系统已经运行多长时间，当前时间，和系统负载。
#USER - 登录用户名
#TTY - 登录用户使用的终端名称
#FROM - 来自登录用户的主机名或者 IP
#LOGIN@ - 用户登录时间
#IDLE - 从用户上次和终端交互到现在的时间，即空闲时间
#JCPU - 依附于 tty 的所有进程的使用时间
#PCPU - 用户当前进程的使用时间。当前进程名称显示在 WHAT
#WHAT - 用户当前进程和选项、参数
```

### 查看所有用户最近一次登录
```shell
lastlog #命令 用于显示系统中所有用户最近一次登录信息。
#lastlog 文件在每次有用户登录时被查询。可以使用 lastlog 命令检查某特定用户上次登录的时间，
#并格式化输出上次登录日志/var/log/lastlog 的内容。它根据 UID 排序显示登录名、端口号（tty）和上
#次登录时间。如果一个用户从未登录过，lastlog 显示**Never logged**。注意需要以 root 身份运行该命令。

# 但是记录中会有很多未登录的用户，可以通过 grep -v 命令进行过滤，不显示没有登录过的用户。
lastlog | grep -v "Never logged in"
```

### 查看历史登录用户以及登录失败的用户
```shell
last # 可以查看所有成功登录到系统的用户记录，
lastb # 查看登录成功和失败的用户记录。
# 单独执行 last 指令时，它会读取位于/var/log/wtmp 的文件，并把该给文件的内容记录的登录系统的用户名单全部显示出来。
# 单独执行 lastb 指令，它会读取位于/var/log/btmp 的文件，并把该文件内容记录的登入系统失败的用户名单，全部显示出来。

# 查看最近 5 个登录的用户
last -n 
# -d ip 地址转换为主机名。该参数可以获取登录到系统的用户所使用的的主机名，如果目标使用的vps 服务器绑定了域名，该参数有可能获取到目标域名。
last -a -d |awk -F' ' '{print $1 "\t" $NF}'
# 对登录系统的用户和 ip 进行排序计数
# 这里去掉-d 参数，因为找不到主机名的地址会显示 error.arpa
last -a |awk -F' ' '{ print $1 "\t" $NF}' |sort |uniq -c |sort -nr
# lastb 查看所有登录记录包含失败
lastb -a |awk -F' ' '{ print $1 "\t" $NF}' |sort |uniq -c |sort -nr
# 登录失败的请求更重要的是 ip 地址信息，所以我们只取 ip 地址进行统计。
lastb -a |awk '{print $NF}' |sort |uniq -c |sort -nr
```
### SSH 登录日志分析
```shell
# 系统用户登录都会在/var/log/secure 日志文件中记录。但是这个日志文件会被系统自动分割。
ll -ld /var/log/secure*
grep Failed /var/log/secure*
# 取出第九列和第十一列。
grep Failed /var/log/secure* |awk -F' ' '{print $9 "\t" $11}'
# 过滤用户名+登录失败的 IP
grep Failed /var/log/secure* |grep -v "invalid"|grep -v "release" |awk -F' ' '{print $9 "\t" $11}' |sort |uniq -c | sort -nr
# 查看登录成功的 ip
grep "Accepted " /var/log/secure* | awk '{print $11}' | sort | uniq -c | sort -nr | more
```

### 查看系统历史命令
```shell
# 系统历史命令一般保存在用户家目录下.bash_history 文件中
find / -name .bash_history
# 查看当前用户的历史命令：
history
```

### 常用系统日志说明
日志目录 作用
- /var/log/message 包括整体系统信息
- /var/log/auth.log 包含系统授权信息，包括用户登录和使用的权限机制等
- /var/log/userlog 记录所有等级用户信息的日志
- /var/log/cron 记录 crontab 命令是否被正确的执行
- /var/log/vsftpd.log 记录 Linux FTP 日志
- /var/log/lastlog 记录登录的用户，可以使用命令 lastlog 查看
- /var/log/secure 记录大多数应用输入的账号与密码，登录成功与否
- /var/log/wtmp 记录登录系统成功的账户信息，等同于命令 last
- /var/log/btmp 记录登录系统失败的用户名单，等同于命令 lastb
- /var/log/faillog 记录登录系统不成功的账号信息，一般会被黑客删除

### 计划任务日志
```shell
# 所有执行过的计划任务都会存在在/var/log/cron 文件中。查看所有执行过的计划任务。
cat /var/log/cron* |awk -F':' '{print $NF}' |grep CMD |sort|uniq -c |sort -rn
# 查看所有用户的计划任务
cat /etc/passwd | cut -f 1 -d : |xargs -i crontab -l -u {}
# 也可以直接查看/var/spool/cron/下的文件内容，所有用户级别的计划任务，都在这里有文件
ls /var/spool/cron/
```

### 检查系统用户
```shell
# Linux 系统用户主要存放于/etc/passwd 文件和/etc/shadow 文件中，还有一个组文件/etc/group。
head -n 1 /etc/passwd

# 正常情况，没有密码，用户是登录不了 linux 的。下面模拟黑客以空口令方式登录系统。
# 创建空口令帐号
vim /etc/shadow #删除 test 密码信息，test 就成为空口令用户
```
> 重点在于 UID 和 GID，root 用户的用户标识为 0，如果一个普通用户的 UID 修改为 0，那么这个用户就成为了 root 用户。


### 中间件日志
Web 攻击的方法多种多样，但是默认情况下 Web 日志中所能记录的内容并不算丰富，最致命的是
web 日志是不会记录 post 内容的，想要从 Web 日志中直接找出攻击者的 webshell 是非常难的，
所以一般来说我们的分析思路都是先通过文件的方式找到 webshell，然后再从日志里找到相应的攻击者ip，再去分析攻击者的整个攻击路径，来回溯攻击者的所有行为。

### 通过时间检查站点被黑客修改过的文件
```shell
# 检查最近 1 天内被修改过的文件
# 注：0 表示 24 小时内修改过的，1 表示昨天修改过的，2 表示前天修改过的。这是个单独的日期想要指定 3 天之前到现在被修改过的文件则需要指定-3
find /www/wwwroot/ecshop.xueshenit.com/ -name "*.php" -mtime -1
# 查看 30 天内修改过的文件示例：-mtime -30
# Linux 文件有 3 个时间属性
atime # access time 访问时间 文件中的数据库最后被访问的时间
mtime # modify time 修改时间 文件内容被修改的最后时间
ctime # change time 变化时间 文件的元数据发生变化。比如权限，所有者等

# stat 命令可以查看文件详细信息。
stat /www/wwwroot/ecshop.xueshenit.com/vulnspy.php
```

### 检查服务器已经建立的网络连接
> 如果黑客已经和服务器建立了连接，可通过查看当前服务器已经建立的链接来分析恶意 ip 和进程。 Linux 中查看网络连接常用 netstat
```shell
# netstat 命令参数
#-a 或--all：显示所有连线中的 Socket；
#-n 或--numeric：直接使用 ip 地址，而不通过域名服务器；
#-p 或--programs：显示正在使用 Socket 的程序识别码和程序名称；
#-t 或--tcp：显示 TCP 传输协议的连线状况；
#-u 或--udp：显示 UDP 传输协议的连线状况；
netstat -anutp
#可以通过第六列筛选已经建立链接的进程。TCP 连接状态详解：
#1、LISTEN #本地服务侦听状态
#2、ESTABLISHED #已经建立链接双方正在通讯状态。
#3、CLOSE_WAIT #对方主动关闭连接或者网络异常导致连接中断，这时我方的状态会变成CLOSE_WAIT 此时我方要调用 close()来使得连接正确关闭。
#4、TIME_WAIT #我方主动调用 close()断开连接，收到对方确认后状态变为TIME_WAIT。
#5、SYN_SENT #半连接状态，原理同 SYN Flood 攻击，攻击者发送 SYN 后服务器端口进入 SYN_SENT 状态等待用户返回 SYN+ACK

# 查看已经建立连接的会话。
netstat -anutp |grep ESTABLISHED
netstat -anutp |grep LISTEN #查看本机处于监听的服务，查看黑客开放了哪些监听端口
```

### 通过 GScan 工具自动排查后门
GScan 是一款为安全应急响应提供便利的工具，自动化监测系统中常见位置。
[入口](https://github.com/grayddq/GScan)


### 巧用 systemd-journald 服务分析系统日志
```shell
# 从新到旧排序使用-r 参数
journalctl -r
# 查看指定服务日志
journalctl -u sshd
# 对比系统日志文件
tail -n 4 /var/log/secure
# 实时查看最新日志。
journalctl -f
# 指定查询时间
journalctl --since "2021-11-05 00:00:00" --until "2021-11-05 17:00:00"
#--since 定义开始时间
#--until 定义结束时间
# 输出详细信息
journalctl -o verbose
```

### 实战清理系统日志后使用 systemd-journald 分析日志
```shell
journalctl --until "2021-11-05 17:47:00" -o short-precise
```
