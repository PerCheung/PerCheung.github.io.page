点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)

<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# CentOS系统环境搭建（四）——安装chatGPT

*Welcome to the AI era!*

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/openai.png" width="20%" alt="Typora" />
</div>

**使用docker compose安装**

在`/usr/local`文件夹下创建`chatgpt`

```bash
mkdir chatgpt
```

创建`docker-compose.yaml`

```bash
vim docker-compose.yaml
```

docker-compose.yaml内容如下

```yaml
version: '3.8'
services:
  chatgpt:
    image: yidadaa/chatgpt-next-web
    restart: always
    network_mode: host
    environment:
      - OPENAI_API_KEY="你的api kay"
      - CODE="你的管理员密码"
      - BASE_URL="ai的访问路径"
    container_name: chatgpt
    hostname: localhost
```

随后使用命令创建容器并运行

```bash
docker compose up -d
```
然后访问一下`http://ip:3000`，看看

𝓘 𝓵𝓸𝓿𝓮 𝔂𝓸𝓾 𝓽𝓱𝓻𝓮𝓮 𝓽𝓱𝓸𝓾𝓼𝓪𝓷𝓭 𝓽𝓲𝓶𝓮𝓼.

爱（AI）你三千遍。