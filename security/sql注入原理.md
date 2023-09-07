## sql 注入原理

### 原理
用户登录的基本 SQL 语句：
> select * from users where username = '用户输入的用户名' and password='用户输入的密码'

用户输入的内容是可控的，例如我们可以在用户名中输入' or 1=1 --空格
> select * from users where username = '' or 1=1 --空格 'and password = '用户输入的密码'

此时我们输入的第一个单引号将 username 的单引号闭合，相当于输入了一个空用户，or 表示左右两边只要有一边条件判断成立则该语句返回结果为真，
其中 1=1 永远为真，所以当前 SQL 语句无论怎么执行结果永远为真，--空格表示注释，注释后面所有代码不再执行。   

我们可以看到上面我们闭合的方法是没有输入用户名的，所以并不能成功登录
> select * from users where username = 'admin' or 1=1 --空格 'and password ='用户输入的密码'

我们在单引号前面加上用户名表示我们要登录的用户。这样就成功绕过了用户密码认证。
单引号的作用：在提交数据或者 URL 中添加单引号进行提交如果返回 SQL 错误即可判断当前位置存在 SQL 注入漏洞。原因是没有被过滤。

##  SQL 注入的分类
SQL 注入的分类基本上都是根据注入的方式进行分类，大概分为以下 4 类
1. 布尔注入：可以根据返回页面判断条件真假的注入；
2. 联合注入：可以使用 union 的注入；
3. 延时注入：不能根据页面返回内容判断任何信息，用条件语句查看时间延迟语句是否执行（即页面返回时间是否增加）来判断；
4. 报错注入：页面会返回错误信息，或者把注入的语句的结果直接返回在页面中；

以上是根据常见的注入方式进行分类，但是通常来说 SQL 注入只分为字符型或者数字型比如：
- 数字型 1 or 1=1
- 字符型 1' or '1'='1

### [sqli-labs 学习环境](https://github.com/Audi-1/sqli-lab)

### SQL 注入联合查询-获取数据库数据

#### 爆出字段的显示位置
```sql
-- UNION 操作符用于合并两个或多个 SELECT 语句的结果集。
-- 请注意，UNION 内部的 SELECT 语句必须拥有相同数量的列。列也必须拥有相似的数据类型。同时，每条 SELECT 语句中的列的顺序必须相同。
-- SQL UNION 语法:
SELECT column_name(s) FROM table_name1
UNION
SELECT column_name(s) FROM table_name2
```

#### 在 MySQL 中使用 union 爆出字段 
```sql
-- 先使用单条语句查询： 
select * from users where id=1;
-- 使用两条语句结合查询
select * from users where id=1 union select 1,2,3;
```
> 由此可以看到两条语句结合输出的结果，而 union select 1,2,3 表示，在上一条语句的查询结果中，再输出 1 2 3 到对应的字段中，因此可以利用来做爆出字段的显示位置

### 基于报错的 SQL 注入
