点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# Windows系统压缩包方式安装redis

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/Redis.png" width="25%" alt="window" />
</div>

> 安装包经过实践是有bug的，一定要用压缩包安装redis

## 1.下载redis

下载链接🔗[https://github.com/tporadowski/redis/releases/download/v5.0.14.1/Redis-x64-5.0.14.1.zip](https://github.com/tporadowski/redis/releases/download/v5.0.14.1/Redis-x64-5.0.14.1.zip)

获得压缩包后解压到任意文件夹下。

如我放到了`C:\App\redis`文件夹下，请将该文件夹配置为环境变量。

## 2.验证

在`cmd`使用下列指令查看redis版本，验证环境变量是否配置成功。

```bash
redis-server -v
```

```bash
redis-cli -v
```

## 3.修改配置

前往文件夹，手动进入文件夹`C:\App\redis`，然后用记事本打开文件`redis.windows-service.conf`，手动修改两个地方。

第一，搜索`bind 127.0.0.1`，将其注释，这样就去除了本地网络回环，也就是开启了外网访问，这样外部电脑就可以访问你本地的redis了，若你对内网安全有要求，就是不想让外网访问，这里就不要改。

```bash
# bind 127.0.0.1
```

第二，开启访问密码，搜索`# requirepass foobared`，解开注释，填入你的密码，若不解开注释，则可以无密码访问。

```bash
requirepass {你的密码}
```

## 4.启动

打开`cmd`

卸载旧服务

```bash
redis-server --service-uninstall
```

添加服务

```bash
"C:\App\redis\redis-server" --service-install "C:\App\redis\redis.windows-service.conf" --service-name redis
```

启动

```bash
redis-server --service-start
```

停止

```bash
redis-server --service-stop
```

删除redis服务

```bash
sc delete redis
```

