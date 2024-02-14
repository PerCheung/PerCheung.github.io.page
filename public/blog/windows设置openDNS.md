点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}

<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# windows设置openDNS

## 1.什么是openDNS？

今天用谷歌的`chrome`的时候，无意间发现了这个东西。

![image-20240131200100029](https://percheung.github.io/blogImg/image-20240131200100029.png)

谷歌的publicDNS我是知道的，就是很多人经常用的`8.8.8.8`。这个openDNS我搜了一下：**OpenDNS**是一个免费的域名解析服务提供商（[DNS](https://zh.wikipedia.org/wiki/DNS)），并具备[反钓鱼](https://zh.wikipedia.org/w/index.php?title=反釣魚&action=edit&redlink=1)、[内容控制软件](https://zh.wikipedia.org/wiki/内容控制软件)等功能。

看起来这个openDNS很好啊！谷歌浏览器既然可以设置DNS提供商，我为什么不直接给我的操作系统设置一个openDNS呢？

## 2.openDNS的ip是多少？

参考<img src="https://percheung.github.io/blogImg/wiki.png" width="40px" alt="" />[维基百科OpenDNS](https://zh.wikipedia.org/zh-cn/OpenDNS)页面`服务地址`部分。查阅可知。

以下为 OpenDNS 的 [IPv4](https://zh.wikipedia.org/wiki/IPv4 "IPv4") 免费 DNS 解析服务器地址：

- `208.67.222.222` （`Resolver1.OpenDNS.com`）
- `208.67.220.220` （`Resolver2.OpenDNS.com`）
- `208.67.222.220` （`Resolver3.OpenDNS.com`）
- `208.67.220.222` （`Resolver4.OpenDNS.com`）

## 3.设置DNS

### 3.1 设置DNS过程讲解

位置是`控制面板\所有控制面板项\网络连接`，找到你正在用的网络，右击选中属性。

![image-20240131201347688](https://percheung.github.io/blogImg/image-20240131201347688.png)

双击`Internet 协议版本 4(TCP/IPv4)`

![image-20240131201508875](https://percheung.github.io/blogImg/image-20240131201508875.png)

设置`DNS`如图。

![image-20240131201704779](https://percheung.github.io/blogImg/image-20240131201704779.png)

首选DNS服务器为`208.67.222.222`

备用DNS服务器为`208.67.220.220`

### 3.2 DNS解析刷新

打开`cmd`输入

```bash
ipconfig /flushdns
```

![image-20240131202150550](https://percheung.github.io/blogImg/image-20240131202150550.png)
