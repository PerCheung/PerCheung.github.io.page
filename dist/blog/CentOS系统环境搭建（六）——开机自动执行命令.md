点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)

# CentOS系统环境搭建（六）——开机自动执行命令

执行如下命令将`/etc/rc.d/rc.local`文标记为可执行文件

```bash
chmod +x /etc/rc.d/rc.local
```

打开`/etc/rc.d/rc.local`文件，在最后面添加要执行的命令

```bash
vim /etc/rc.d/rc.local
```

例如

`>/dev/null 2>&1` 将所有输出重定向到空设备，即不输出日志。最后的 `&` 符号将命令放到后台运行。

```bash
/srv/system/run.sh >/dev/null 2>&1 &
```