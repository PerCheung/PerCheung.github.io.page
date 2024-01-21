点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# Windows10解决大小核调度问题

> 该教程是给笔记本电脑用的，经过我实践是成功的。

## 1.开启高性能模式

> 使用管理员模式的PowerShell输入下列指令
>

```bash
powercfg -s 8c5e7fda-e8bf-4a96-9a85-a6e23a8c635c
```

## 2.下载安装PowerSettingsExplorer

**这个软件是必须的**！

*我参考的教程是这么说的：网上让你用CMD或者修改注册表的解锁这三个选项的，对于笔记本无效，笔记本用的是“现代电源计划”，表现为电源选项里只有一个平衡，性能释放只能靠滑块，此时只能用PowerSettingsExplorer*。

我经过实践，发现该软件仅97kb大小，而且用的是windows系统原生配置，不侵入，算得上是非常不错的解决方案了。

下载链接🔗[https://www.mediafire.com/file/wt37sbsejk7iepm/PowerSettingsExplorer.zip](https://www.mediafire.com/file/wt37sbsejk7iepm/PowerSettingsExplorer.zip)

下载后解压，双击运行里面的界面，会给你弹出一个很古朴的页面，后面照我的图做就好。

## 3.修改配置

### 生效的异类策略

> 使用异类策略0

![image-20240116011513598](https://percheung.github.io/blogImg/image-20240116011513598.png)

### 异类线程调度策略

> 首选高性能处理器

![image-20240116011704650](https://percheung.github.io/blogImg/image-20240116011704650.png)

### 异类短时间线程调度策略

> 首选高性能处理器

![image-20240116011735181](https://percheung.github.io/blogImg/image-20240116011735181.png)

## 4.你的电源策略

> 控制面板\所有控制面板项\电源选项

*这里当然要记得勾选你的高性能模式。*

![image-20240116012036339](https://percheung.github.io/blogImg/image-20240116012036339.png)

## 5.CPU展示

> 在任务管理器里`右击`，照我的图做，可以看到你电脑上所有的cpu的运行情况，这里可以看到大核小核都在运行，利用率百分之百。

![image-20240116012426299](https://percheung.github.io/blogImg/image-20240116012426299.png)
