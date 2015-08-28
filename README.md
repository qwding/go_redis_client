# redis_client

* 这是一个简单的redis使用客户端，用go语言编写。
* 下载代码后可以构建一下，即可运行
* 使用的是redigo包，采用连接池来连接，
* 使用方法 ./main -host=myredis.com:6379 -db=0 -password=mima,默认的话连接127.0.0.1:6379,密码为password
* 输入 `hep` 可查看帮助

### 使用方法
 
* 类似命令行输入 
* `keys *`（查看所有键值对）
* delkeys 2015_08_27_*
