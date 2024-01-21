点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# CentOS系统环境搭建（十）——运行Java服务<img src="https://percheung.github.io/blogImg/java.png" width="50px" alt="ffmpeg" />

> 我是用jar包的方式运行的Java服务，下面讲讲我是怎么做的，这里的内容是我腾讯服务器部署的README.md文件的内容。

## 以jar包方式部署

> 目录位置

```bash
/srv/tencent/server
```

> 进入目录

```bash
cd /srv/tencent/server
```

### 1.停止服务

```bash
./stop_jar.sh
```

> stop_jar.sh文件内容

```bash
#!/bin/bash

pkill -f /srv/tencent/server/tencent-2023.09.13.jar
```

### 2.运行服务

```bash
./run_jar.sh
```

> run_jar.sh文件内容

```bash
#!/bin/bash

# 分配堆内存大小（总内存的四分之一）
heap_memory=1
# 获取线程数量
thread_count=$(nproc)

# 启动 Java 应用程序
#`-server`：指定Java虚拟机以服务器模式运行，以获得更好的性能。
#`-Dfile.encoding=UTF-8`：设置文件编码为UTF-8，确保正确处理文本数据。
#`-XX:ActiveProcessorCount=8`：指定并行处理器的数目为8，以影响Java虚拟机的线程和CPU利用情况。
#`-Xms2g`：指定Java虚拟机的最小堆内存为2g。
#`-Xmx4g`：指定Java虚拟机的最大堆内存为4g。
#`-XX:+UseParallelGC`：启用并行垃圾回收器，以提高垃圾回收的效率。
#`-XX:+UseCompressedOops`：启用压缩指针，以减少内存消耗。
#`-XX:ParallelGCThreads=8`：指定并行垃圾回收器的线程数为8。
#`-XX:+DisableExplicitGC`：禁用显式的垃圾回收调用
#`-XX:+HeapDumpOnOutOfMemoryError`：当发生堆溢出异常时，自动生成堆转储文件（Heap Dump）。
#`-Djava.net.preferIPv4Stack=true`：该属性为 true 可确保 JVM 在具备 IPv6 支持的情况下仍然优先使用 IPv4。
nohup java -server -Dfile.encoding=UTF-8 -XX:ActiveProcessorCount=$thread_count -Xms${heap_memory}G -Xmx${heap_memory}G -XX:+DisableExplicitGC -XX:+UseParallelGC -XX:+UseCompressedOops -XX:ParallelGCThreads=$thread_count -XX:+HeapDumpOnOutOfMemoryError -Djava.net.preferIPv4Stack=true -jar /srv/tencent/server/tencent-2023.09.13.jar &
```

### 3.查看控制台信息

```bash
./see_log.sh
```

> see_log.sh文件内容

```bash
#!/bin/bash

# 实时查看日志末尾1000行
tail -f -n 1000 nohup.out
```

### 4. 删除nohup.out文件中的旧数据（只保留末尾二十万行）

```bash
./clearlog.sh
```

> clearlog.sh文件内容

```bash
#!/bin/bash

# 保留末尾二十万行
tail -n 200000 nohup.out > temp.txt && mv temp.txt nohup.out
```

### 5.日志位置

```bash
cd /srv/tencent/server/logs
```

### 6.文件位置

```bash
cd /srv/tencent/server/file
```

### 7.nginx部署

```bash
        location /api/ {
            # 设置代理服务器发送的http请求头中Host的值为接收到的客户端请求头中的Host值
            proxy_set_header Host $http_host;
            # 设置代理服务器发送的http请求头中X-Real-IP的值为当前请求的客户端IP地址
            proxy_set_header X-Real-IP $remote_addr;
            # 设置代理服务器发送的http请求头中REMOTE-HOST的值为当前请求的客户端IP地址
            proxy_set_header REMOTE-HOST $remote_addr;
            # 设置代理服务器发送的http请求头中X-Forwarded-For的值为所有经过的代理服务器IP地址列表，多个IP地址之间用逗号隔开
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            # 将客户端请求通过代理服务器转发到指定地址
            proxy_pass http://127.0.0.1:8080/api/;
            # 设置代理服务器与后端服务器建立连接的超时时间
            proxy_connect_timeout 3600s;
            # 设置代理服务器从后端服务器读取数据的超时时间
            proxy_read_timeout  3600s;
            # 设置代理服务器向后端服务器发送数据的超时时间
            proxy_send_timeout  3600s;
        }
```

> nginx重新加载配置文件

```bash
/usr/local/nginx/sbin/nginx -s reload
```