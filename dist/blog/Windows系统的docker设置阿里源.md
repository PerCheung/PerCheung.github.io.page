点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# Windows系统的docker设置阿里源

由于我们生活在中国大陆，所以外网的访问总是那么慢又困难，用docker拉取几兆的小镜象还能忍受，拉取几百兆的大型镜像的时候，我只能说：！！！！！

这时候，阿里就是救赎。

## 1.获得镜像加速器

访问阿里镜像加速器网站[https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors](https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors)

![image-20240119190838134](https://percheung.github.io/blogImg/image-20240119190838134.png)

## 2.配置docker

复制一下，然后打开你的docker。我们的配置json不一定一样，不用管。

![image-20240119191331480](https://percheung.github.io/blogImg/image-20240119191331480.png)

在json里添加一句

```json
"registry-mirrors": ["https://XXXXX.mirror.aliyuncs.com"]
```

然后点击提交，我这个版本就会重启你的docker。你的版本如果没有，记得手动重启一下你的docker。

打完收工！

