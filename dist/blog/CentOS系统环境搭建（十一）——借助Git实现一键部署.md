点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
# CentOS系统环境搭建（十一）——借助Git实现一键部署<img src="https://percheung.github.io/blogImg/deploy.png" width="40px" alt="" />

> 我希望只写一个脚本点击一下就能部署最新代码的前后端服务！让我们想想，部署总是很麻烦，一堆一堆的重复操作，如何将重复的步骤去掉，整合进一个sh脚本里，将前后端部署简化为一步操作呢？

## 1.后端

### 1.1 拉取代码

进入后端代码所在路径。

```bash
cd /srv/tencent/code/tencent
```

在此文件夹下执行拉取命令。

```bash
git pull
```

### 1.2 打包

有了最新的源码以后，就可以执行打jar包操作了。

这是一个spring boot项目，可以看到我的项目结构如下

```bash
[root@VM-4-17-centos tencent]# ll
total 52
-rw-r--r-- 1 root root 35149 Nov 21 15:09 LICENSE
-rw-r--r-- 1 root root  4157 Nov 21 15:09 pom.xml
-rw-r--r-- 1 root root  3828 Nov 21 15:09 README.md
drwxr-xr-x 4 root root  4096 Nov 21 15:09 src
```

我们在此文件夹`/srv/tencent/code/tencent`下执行打包命令。

```bash
mvn clean package
```

成功后，我们的`/srv/tencent/code/tencent/target`下就会有一个打好的jar包了。

```bash
cd /srv/tencent/code/tencent/target
```

## 2.前端

### 拉取代码

前端代码保存于

```bash
cd /srv/tencent/code/page/dist
```

更新代码

```bash
git pull origin main
```

## 3.一键部署

思考一下，整个步骤就这么几步。

1. git拉取后端代码。
2. 打包。
3. 终止旧服务的jar包。
4. 复制新jar包替换旧jar包。
5. 启动新jar包。
6. git拉取前端代码。
7. 将新的dist文件夹覆盖旧的dist文件夹。
8. 重启nginx。

```bash
#!/bin/bash

# 后端代码更新和打包
cd /srv/tencent/code/tencent
git pull
mvn clean package

# 停止旧的后端服务
cd /srv/tencent/server
./stop_jar.sh

# 替换jar包
cp /srv/tencent/code/tencent/target/tencent-2023.09.13.jar /srv/tencent/server/tencent-2023.09.13.jar

# 启动新的后端服务
cd /srv/tencent/server
./run_jar.sh

# 前端代码更新
cd /srv/tencent/code/page/dist
git pull origin main

# 替换前端文件夹
rsync -av --delete /srv/tencent/code/page/dist/ /srv/tencent/page/dist/

# 重新加载 Nginx
cd /usr/local/nginx/sbin
./nginx -s reload
```

将脚本放在`/srv/tencent`下

```bash
cd /srv/tencent
```
新建部署脚本
```bash
vim deploy.sh
```
赋予执行权限
```bash
chmod +x deploy.sh
```

执行

```bash
./deploy.sh
```
只需要更新一下你的GitHub仓库，然后运行脚本，就会更新你的服务器服务。