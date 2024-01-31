点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 如何使用WebDAV云同步

## 1.WebDAV是什么？

### 1.1 发现WebDAV的背景

不久前我写过如何搭建chatGPT的web页面的博客，而且自己也搭建了一个，后面`ChatGPT-Next-Web`更新后，出了一个云端同步的功能，我一看，用的软件是WebDAV，搜了一下，没发现这个软件啊？这怎么用。查阅资料才知道，原来这不是软件，而是一种协议，找一个支持该协议的软件，就可以做云同步了。

### 1.2 WebDAV是什么？

查询自<img src="https://percheung.github.io/blogImg/wiki.png" width="55px" alt="" />维基百科。

**基于Web的分布式编写和版本控制**（英语：**W**eb-based **D**istributed **A**uthoring and **V**ersioning，缩写：**WebDAV**）是[超文本传输协议](https://zh.wikipedia.org/wiki/%E8%B6%85%E6%96%87%E6%9C%AC%E4%BC%A0%E8%BE%93%E5%8D%8F%E8%AE%AE "超文本传输协议")（HTTP）的扩展，有利于用户间协同编辑和管理存储在[万维网](https://zh.wikipedia.org/wiki/%E4%B8%87%E7%BB%B4%E7%BD%91 "万维网")服务器文档。WebDAV由[互联网工程任务组](https://zh.wikipedia.org/wiki/%E4%BA%92%E8%81%94%E7%BD%91%E5%B7%A5%E7%A8%8B%E4%BB%BB%E5%8A%A1%E7%BB%84 "互联网工程任务组")的工作组在[RFC](https://zh.wikipedia.org/wiki/RFC "RFC") [4918](https://tools.ietf.org/html/rfc4918 "rfc:4918")中定义。

WebDAV协议为用户在[服务器](https://zh.wikipedia.org/wiki/%E6%9C%8D%E5%8A%A1%E5%99%A8 "服务器")上创建、更改和移动文档提供了一个框架。WebDAV协议最重要的功能包括作者或修改日期等属性的维护、[命名空间](https://zh.wikipedia.org/wiki/%E5%91%BD%E5%90%8D%E7%A9%BA%E9%97%B4 "命名空间")管理、集合和覆盖保护。为属性维护所提供的功能包括创建、删除和查询文件信息等；**命名空间管理**处理在服务器名称空间内复制和移动网页的能力；**集合**（Collections）处理各种资源的创建、删除和列举；**覆盖保护**处理与锁定文件相关的问题。WebDAV协议利用[TLS](https://zh.wikipedia.org/wiki/TLS "TLS")、[HTTP摘要认证](https://zh.wikipedia.org/wiki/HTTP%E6%91%98%E8%A6%81%E8%AE%A4%E8%AF%81 "HTTP摘要认证")、[XML](https://zh.wikipedia.org/wiki/XML "XML")等技术来满足这些需求。

许多现代[操作系统](https://zh.wikipedia.org/wiki/%E6%93%8D%E4%BD%9C%E7%B3%BB%E7%BB%9F "操作系统")为WebDAV提供了内置的客户端支持。

链接🔗[WebDAV - 维基百科，自由的百科全书](https://zh.wikipedia.org/zh-sg/WebDAV)

## 2.使用WebDAV

### 国内推荐<img src="https://percheung.github.io/blogImg/jianguoyun.png" width="55px" alt="" />坚果云

要在坚果云中开启WebDAV并为ChatGPT-Next-Web生成第三方应用授权密码，按照以下步骤进行操作：

1. 打开坚果云官网（https://www.jianguoyun.com/）并登录你的坚果云账号。

2. 点击右上角的账户名，选择【账户信息】，然后选择【安全选项】。

3. 在“安全选项”中找到“第三方应用管理”，选择“添加应用密码”，输入名称（例如：chatGPT），然后选择“生成密码”，最后选择完成。这样就生成了一个应用密码，用于授权第三方应用访问你的坚果云文件。

4. 打开ChatGPT-Next-Web的设置，选择云端数据的配置按钮，点击进入WebDAV配置页面。

5. 在WebDAV配置页面中，填写坚果云提供的WebDAV信息，包括服务器地址、用户名和应用密码。

6. 随后，点击同步按钮即可完成同步。

> 这样，我们就拥有了一个免费的，支持云同步chatGPT平台。
>

   
