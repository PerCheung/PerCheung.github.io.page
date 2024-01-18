点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({ top: 0, behavior: 'smooth' });" style="background-color: transparent; position: fixed; bottom: 20px; right: 40px; padding: 10px 10px 5px 10px; cursor: pointer; z-index: 10; border-radius: 13%; box-shadow: 0.5px 3px 7px rgba(0, 0, 0, 0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color: transparent; width: 30px;"></div>

# 安装最新版Docker和Docker Compose

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/Docker.png" width="25%" alt="Docker" />
</div>

## 1.添加Docker库

### 1.1 安装必要的证书并允许 apt 包管理器使用以下命令通过 HTTPS 使用存储库

```bash
sudo apt install apt-transport-https ca-certificates curl software-properties-common gnupg lsb-release
```

### 1.2 运行下列命令添加 Docker 的官方 GPG 密钥

```bash
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```

### 1.3 添加 Docker 官方库

```bash
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

### 1.4 更新源列表

```bash
sudo apt update
```

## 2.安装 Docker

### 2.1 安装最新版

```bash
sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

### 2.2 验证docker是否运行

```bash
systemctl status docker
```

### 2.3 设置docker开机自启动

```bash
sudo systemctl start docker
```

```bash
sudo systemctl enable docker
```

## 3.验证

### 3.1 查看docker版本

```bash
sudo docker version
```

### 3.2 测试 Docker

```bash
sudo docker run hello-world
```

### 3.2 查看docker compose版本

> 最新版的docker自带Docker Compose

```bash
sudo docker compose version
```

