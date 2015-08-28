# redis_client

* 这是一个简单的redis使用客户端，用go语言编写。
* 下载代码后可以构建一下，即可运行
* 使用的是redigo包，采用连接池来连接，
* 使用方法 ./main -host=127.0.0.1:6379 -db=0 -password=mima
* 可以附加自己的方法，目前增加了群删操作
* delkeys 2015_08_27_*
