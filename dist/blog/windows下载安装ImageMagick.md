点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# windows下载安装ImageMagick

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/ImageMagick.png" width="30%" alt="ImageMagick" />
</div>

## 1.下载

下载页面地址是[https://imagemagick.org/script/download.php#windows](https://imagemagick.org/script/download.php#windows)

你会进入到页面的**Windows Binary Release**部分，然后接着看。

![image-20240121145902632](https://percheung.github.io/blogImg/image-20240121145902632.png)

这里我选择用该安装包，理由是：

1. x64架构（这里根据自己电脑实际情况分析）。
2. 每像素 16 位组件（16 bits-per-pixel component）。
3. 启用[高动态范围成像](https://imagemagick.org/script/high-dynamic-range.php)（with high dynamic-range imaging enabled）。
4. 便携式，只需复制到您的主机并运行（无需安装程序，无需 Windows 注册表项）。

## 2.安装

*这里要注意一个小坑，这个压缩包居然没有父文件夹，解压后会所有的exe文件直接全蹦出来，建议解压选择解压到某某目录下。*

下载压缩包后，放在本地解压即可。

将解压后的文件放到自定义目录。

例如我放到C:\App\ImageMagick下。那么将C:\App\ImageMagick配置到path环境变量就算安装完成。

## 3.验证

打开一个cmd面板输入下面的指令即可。

> 用的是官网给的方式

```bash
magick logo: logo.gif
magick identify logo.gif
magick logo.gif win:
```

### 3.1 依赖缺失问题

官网提到：If you have any problems, you likely need vcomp140.dll. To install it, download [Visual C++ Redistributable Package](https://support.microsoft.com/en-us/help/2977003/the-latest-supported-visual-c-downloads).

意思是你的windows系统可能会由于缺少`vcomp140.dll`文件导致报错失败，这时候解决方案很简单，去微软的网站[Visual C++ Redistributable Package](https://support.microsoft.com/en-us/help/2977003/the-latest-supported-visual-c-downloads)，直接找最新版，最上面那个，选择你的电脑对应架构的安装包，下载安装即可。

![image-20240121150814967](https://percheung.github.io/blogImg/image-20240121150814967.png)

这个的安装下载就非常简单了，不赘述，无脑下一步即可。

### 3.2 依赖缺失解决

成功后会再次执行官网给的命令，则执行命令的目录下会产生一张ImageMagick的logo图片，这样就算安装完成，关于ImageMagick的具体使用请看我的下一篇博客。

[ImageMagick使用手册](https://percheung.github.io/blog/ImageMagick使用手册)