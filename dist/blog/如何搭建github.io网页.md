###### 点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
# 如何搭建github.io网页

*在我的理解里，github.io，不是博客，这是服务器。我的博客网址为*🔗[https://percheung.github.io](https://percheung.github.io/)

> 我就默认你已经有GitHub账户了。

## 1.github.io仓库的初始化

### 1.1 创建仓库

进入[https://github.com/new](https://github.com/new)，如图。新建仓库{你的GitHub账户名}.github.io，记得勾选public。

![image-20231226165946916](https://percheung.github.io/blogImg/202312261659154.png)

### 1.2 设置主题

新建文件`_config.yml`，内容写

```yaml
theme: jekyll-theme-minimal
```

![image-20231228171732435](https://percheung.github.io/blogImg/202401020949474.png)

现在官方具备如下主题。

```yaml
# 主题设置
theme: jekyll-theme-architect
theme: jekyll-theme-cayman
theme: jekyll-theme-dinky
theme: jekyll-theme-hacker
theme: jekyll-theme-leap-day
theme: jekyll-theme-merlot
theme: jekyll-theme-midnight
theme: minima
theme: jekyll-theme-minimal
theme: jekyll-theme-modernist
theme: jekyll-theme-slate
theme: jekyll-theme-tactile
theme: jekyll-theme-time-machine
```

最新的主题有什么具体可以参考[GitHub主题页面](https://pages.github.com/themes/)。

设置好后，你项目里的markdown文件就会随着主题的变化而变化。这里也可以什么主题都不选，那么会有一个默认的主题。

这是我的`_config.yml`文件内容。

```yaml
title: My Blog
description: 鸿雁长飞光不度，鱼龙浅跃水成文
```

