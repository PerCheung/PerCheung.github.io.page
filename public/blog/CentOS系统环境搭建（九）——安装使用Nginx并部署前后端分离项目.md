点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# CentOS系统环境搭建（九）——安装使用Nginx并部署前后端分离项目

本文使用的nginx版本为`1.22.1`

> Nginx发布版本分为主线版本和稳定版本，区分两个版本也非常简单，主线版本版本号为单数，比如1.19，稳定版本为双数，比如1.18。

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/Nginx.png" width="30%" alt="ffmpeg" />
</div>

## 1.安装

### 1.1下载

nginx下载链接：https://nginx.org/en/download.html

中间的是Linux版本的nginx压缩包，右边的是windows系统用的，我们选择下载中间的1.22.1版本。

### 1.2 检验服务器上是否有nginx

```bash
whereis nginx
```

如果什么都没有输出，说明电脑上没有nginx。

### 1.3 解压安装

将下载好的nginx压缩包放在服务器的任意位置即可。

进入该目录执行解压操作。

```bash
tar -zxvf nginx-1.22.1.tar.gz
```

解压后多出来一个文件夹，名字为nginx-1.22.1。

```bash
[root@VM-4-17-centos nginx]# ll
total 1056
drwxr-xr-x 8 1001 mysql    4096 Oct 19  2022 nginx-1.22.1
-rw-r--r-- 1 root root  1073948 Jun 25 17:33 nginx-1.22.1.tar.gz
```

进入文件夹，展开如下。

```bash
[root@VM-4-17-centos nginx]# cd nginx-1.22.1/
[root@VM-4-17-centos nginx-1.22.1]# ll
total 824
drwxr-xr-x 6 1001 mysql   4096 Jun 25 17:34 auto
-rw-r--r-- 1 1001 mysql 317399 Oct 19  2022 CHANGES
-rw-r--r-- 1 1001 mysql 485035 Oct 19  2022 CHANGES.ru
drwxr-xr-x 2 1001 mysql   4096 Jun 25 17:34 conf
-rwxr-xr-x 1 1001 mysql   2590 Oct 19  2022 configure
drwxr-xr-x 4 1001 mysql   4096 Jun 25 17:34 contrib
drwxr-xr-x 2 1001 mysql   4096 Jun 25 17:34 html
-rw-r--r-- 1 1001 mysql   1397 Oct 19  2022 LICENSE
drwxr-xr-x 2 1001 mysql   4096 Jun 25 17:34 man
-rw-r--r-- 1 1001 mysql     49 Oct 19  2022 README
drwxr-xr-x 9 1001 mysql   4096 Jun 25 17:34 src
```

可以看到有一个叫`configure`的可执行文件。

执行，这里会为你安装https模块，避免以后升级<img src="https://percheung.github.io/blogImg/https-service.png" width="40px" alt="ffmpeg" />有问题。

```bash
./configure --prefix=/usr/local/nginx --with-http_stub_status_module --with-http_ssl_module --with-http_v2_module
```

命令运行完成后输入

```bash
make
```

接着是

```bash
make install
```

随后，我们输入

```bash
whereis nginx
```

你将看到

```bash
[root@VM-4-17-centos nginx-1.22.1]# whereis nginx
nginx: /usr/local/nginx
```

这说明nginx已经安装完成，删掉我们的压缩包和解压文件即可。

```bash
rm -rf nginx-1.22.1
```

```bash
rm -rf nginx-1.22.1.tar.gz
```

### 1.4 验证

进入`/usr/local/nginx/sbin`

```bash
cd /usr/local/nginx/sbin
```

可以看到里面有一个名字叫nginx的可执行文件。

```bash
[root@VM-4-17-centos nginx]# cd /usr/local/nginx/sbin
[root@VM-4-17-centos sbin]# ll
total 3804
-rwxr-xr-x 1 root root 3892016 Jun 25 17:36 nginx
```

运行。

```bash
./nginx
```

确保服务器的80端口已经打开，随后在浏览器栏输入ip，进入即可

## 2.部署

### 2.1基本知识

#### 2.1.1常用命令

启动

```bash
./nginx
```

停止

```bash
./nginx -s stop
```

安全退出

```bash
./nginx -s quit
```

重新加载配置文件

```bash
./nginx -s reload
```

查看nginx进程

```bash
ps aux|grep nginx
```

#### 2.1.2配置文件

配置文件位置如下。

```bash
/usr/local/nginx/conf/nginx.conf
```

```ini
# Nginx所属用户和用户组，这里配置的是nobody。由于Nginx并不会直接向客户端发送数据，只有在请求到达时才加入到Nginx的I/O处理队列中，所以Nginx只需要拥有访问目录的权限即可，无需高额权限
#user  nobody;                     

# 工作进程数，一般设置为系统CPU核心数。对于单核服务器来说，设置为1即可；对于多核服务器，可以设置成CPU核心数
worker_processes  2;               

# 定义错误日志文件路径。Nginx生产环境重要的调试信息将会记录在该文件中
#error_log  logs/error.log;        

#error_log  logs/error.log  notice;
#error_log  logs/error.log  info;

# 定义Nginx主进程的PID文件所在位置。若需要向Nginx主进程发送信号时，必须要知道这个进程的PID，因此需要记录到PID文件中
#pid        logs/nginx.pid;         

events {
	# 最大并发数。当同时有多个客户端访问Nginx时，每个访问者占据一个连接，当连接数到达该配置变量的值时，新的访问请求将会等待，直到空闲连接数不到该值再处理新的请求
    worker_connections  1024;      
}

http {
	 # 引入mime.types文件，它可以为不同的文件扩展名设定不同的MIME类型。Nginx不像Apache这样会根据文件扩展名自动推断文件类型，如果文件没有被正确识别MIME类型，就可能会出现浏览器无法正确解释的问题
    include       mime.types;     
	
	# 默认MIME类型，如果服务器无法识别它，那么就让浏览器自己来识别。当Nginx无法识别当前返回的MIME类型时，会采用该参数作为默认值
    default_type  application/octet-stream;   
	
	# 配置日志格式，该部分被注释掉，不起作用
    #log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
    #                  '$status $body_bytes_sent "$http_referer" '
    #                  '"$http_user_agent" "$http_x_forwarded_for"';

	# 定义访问日志文件的存储位置，该部分被注释掉，不起作用
    #access_log  logs/access.log  main;
	
	# 开启高效文件传输，建议保持默认的开启状态。Nginx的高效文件传输技术可以让其更快地将文件内容传输到客户端，提高文件传输速度
    sendfile        on;     
	
	# 禁用TCP节点推送
	#tcp_nopush     on;
	
	#keepalive_timeout  0;
	# 客户端保持连接的时间，超时后断开连接。连接复用机制，当客户端发送完请求后，连接还会继续保持一段时间，以便快速发起下一次请求
    keepalive_timeout  65;  

	# 开启数据压缩功能
	#gzip  on;	

    server {
		# Nginx监听端口，当有请求到达该端口时，Nginx将会处理该请求。当客户端发送了请求过来后，Nginx会监听到该端口被占用，便会启用该server块去处理该请求
        listen       80;            
		
		# 指定服务端的域名/IP，如果没有域名，则使用IP地址。当Nginx接收到一个请求时，首先会按照请求中的Host（或X-Forwarded-Host）头信息的域名取匹配server_name指令中的值，只要匹配上即可
        server_name  localhost;    

		# 该行是将字符集设置为koi8-r，koi8-r是一种字符编码格式
		#charset koi8-r;
        # 设置字符集为 utf-8
        charset utf-8;

		# 该行是定义了当前这个server块的访问日志文件的存储位置。其中logs/host.access.log表示存储的目录和文件名，main表示使用标准的日志格式。如果该行没有被注释掉，就会在相应的目录下生成一份日志文件，记录该server块的访问日志
        #access_log  logs/host.access.log  main;

		# URL匹配规则。在Nginx中，URL匹配是由location指令提供的
        location / {                
			# 当URL中的路径为空时，会默认从这里找到对应的文件。该块内容的作用是指定处理该请求所使用的文件系统根目录路径
            root   /srv/tencent/page/mytencentpage/dist;            
			# 当URL中路径为空时，默认返回欢迎页面。当请求的URI是目录时，Nginx搜索该目录下的index文件，若搜索失败，则返回403
            index  index.html index.htm;
        }

        # 后端接口配置
        location /api {
            # 设置代理服务器发送的http请求头中Host的值为接收到的客户端请求头中的Host值
            proxy_set_header Host $http_host; 
            # 设置代理服务器发送的http请求头中X-Real-IP的值为当前请求的客户端IP地址
            proxy_set_header X-Real-IP $remote_addr; 
            # 设置代理服务器发送的http请求头中REMOTE-HOST的值为当前请求的客户端IP地址
            proxy_set_header REMOTE-HOST $remote_addr;
            # 设置代理服务器发送的http请求头中X-Forwarded-For的值为所有经过的代理服务器IP地址列表，多个IP地址之间用逗号隔开
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; 
            # 将客户端请求通过代理服务器转发到指定地址
            proxy_pass http://127.0.0.1:8082/; 
            # 设置代理服务器与后端服务器建立连接的超时时间
            proxy_connect_timeout 3600s; 
            # 设置代理服务器从后端服务器读取数据的超时时间
            proxy_read_timeout  3600s; 
            # 设置代理服务器向后端服务器发送数据的超时时间
            proxy_send_timeout  3600s; 
        }

		# 配置 404 页面。当请求的文件不存在，或者因为某种原因无法访问到指定的文件时，返回404
        #error_page  404              /404.html;  

		# 错误页面的配置，当出错时会跳转到该页面。当服务端返回错误状态码时，可以为不同的状态码指定不同的错误页面
        error_page   500 502 503 504  /50x.html;    

		# 当请求 URL 精确匹配 /50x.html 时，根据配置路径返回静态页面。配合error_page使用，可以为错误页面指定特定的URI
        location = /50x.html {        
            root   html;
        }

        # 将 PHP 脚本代理到监听在 127.0.0.1:80 的 Apache 服务器，或将 PHP 脚本代理到监听在 127.0.0.1:9000 的 FastCGI 服务器
        #location ~ \.php$ {
        #    proxy_pass   http://127.0.0.1;
        #}
        #location ~ \.php$ {
        #    root           html;
        #    fastcgi_pass   127.0.0.1:9000;
        #    fastcgi_index  index.php;
        #    fastcgi_param  SCRIPT_FILENAME  /scripts$fastcgi_script_name;
        #    include        fastcgi_params;
        #}

        # deny access to .htaccess files, if Apache's document root concurs with nginx's one
        #location ~ /\.ht {
        #    deny  all;
        #}
    }

    # 另一个使用 IP、名称和端口的虚拟主机配置
    #
    #server {
    #    listen       8000;
    #    listen       somename:8080;
    #    server_name  somename  alias  another.alias;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}

    # HTTPS 服务器
    #
    #server {
    #    listen       443 ssl;
    #    server_name  localhost;

    #    ssl_certificate      cert.pem;
    #    ssl_certificate_key  cert.key;

    #    ssl_session_cache    shared:SSL:1m;
    #    ssl_session_timeout  5m;

    #    ssl_ciphers  HIGH:!aNULL:!MD5;
    #    ssl_prefer_server_ciphers  on;

    #    location / {
    #        root   html;
    #        index  index.html index.htm;
    #    }
    #}
}
```

### 2.2 配置效果

我2.1.2的配置文件中，基本每一个配置都有注释。

其中关键的地方有两个。

#### 前端

```ini
# URL匹配规则。在Nginx中，URL匹配是由location指令提供的
        location / {                
			# 当URL中的路径为空时，会默认从这里找到对应的文件。该块内容的作用是指定处理该请求所使用的文件系统根目录路径
            root   /srv/tencent/page/mytencentpage/dist;            
			# 当URL中路径为空时，默认返回欢迎页面。当请求的URI是目录时，Nginx搜索该目录下的index文件，若搜索失败，则返回403
            index  index.html index.htm;
        }
```

我是使用vue创建的前端，`/srv/tencent/page/mytencentpage/dist`就是产生的dist文件夹的地址，这种静态资源，只需要配置到root后面，`location /`表示，当进入http://ip/，就是进入了/srv/tencent/page/mytencentpage/dist。

#### 后端

```ini
# 后端接口配置
        location /api {
            # 设置代理服务器发送的http请求头中Host的值为接收到的客户端请求头中的Host值
            proxy_set_header Host $http_host; 
            # 设置代理服务器发送的http请求头中X-Real-IP的值为当前请求的客户端IP地址
            proxy_set_header X-Real-IP $remote_addr; 
            # 设置代理服务器发送的http请求头中REMOTE-HOST的值为当前请求的客户端IP地址
            proxy_set_header REMOTE-HOST $remote_addr;
            # 设置代理服务器发送的http请求头中X-Forwarded-For的值为所有经过的代理服务器IP地址列表，多个IP地址之间用逗号隔开
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for; 
            # 将客户端请求通过代理服务器转发到指定地址
            proxy_pass http://127.0.0.1:8082/; 
            # 设置代理服务器与后端服务器建立连接的超时时间
            proxy_connect_timeout 3600s; 
            # 设置代理服务器从后端服务器读取数据的超时时间
            proxy_read_timeout  3600s; 
            # 设置代理服务器向后端服务器发送数据的超时时间
            proxy_send_timeout  3600s; 
        }
```

最关键的地方在于`location /api`和`proxy_pass http://127.0.0.1:8082/; `，这意味着当进入http://ip/api时，就进入http://127.0.0.1:8082/，我本地java服务就运行在8082端口，这就达到了从80端口，/api资源进入8082服务的效果。

> 切记，每次修改nginx.conf文件后，用`./nginx -s reload`命令让nginx重新加载。