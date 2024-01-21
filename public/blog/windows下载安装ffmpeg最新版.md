点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# windows下载安装ffmpeg最新版
<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/FFmpeg.png" width="30%" alt="ffmpeg" />
</div>
## 1.下载

下载页面地址是[https://ffmpeg.org/download.html](https://ffmpeg.org/download.html)

![image-20240121144302906](https://percheung.github.io/blogImg/image-20240121144302906.png)

![image-20240121143829969](https://percheung.github.io/blogImg/image-20240121143829969.png)

## 2.安装

下载7z压缩包后，放在本地解压即可。

将解压后的文件放到自定义目录。

例如我放到C:\App\ffmpeg下。那么将C:\App\ffmpeg\bin配置到path环境变量就算安装完成。

## 3.验证

```bash
ffmpeg -version
```

如果输出的有结果，就算下载安装完成。关于ffmpeg的具体使用请看我的下一篇博客。

[ffmpeg使用手册](https://percheung.github.io/blog/ffmpeg使用手册)

