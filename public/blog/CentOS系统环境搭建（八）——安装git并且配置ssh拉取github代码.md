点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# CentOS系统环境搭建（八）——安装<img src="https://percheung.github.io/blogImg/git.png" width="40px" alt="" />git并且配置ssh拉取github<img src="https://percheung.github.io/blogImg/github.png" width="40px" alt="" />代码

## 1.git初始配置

安装

```bash
yum install git
```

验证

```bash
git --version
```

配置用户名（记得用你自己的）

```bash
git config --global user.name "PerCheung"
```

配置邮箱（记得用你自己的）

```bash
git config --global user.email C2243736958@qq.com
```

查看配置

```bash
git config --list
```

## 2.设置ssh

生成公钥和私钥，这里的邮箱用你的GitHub账户

```bash
ssh-keygen -t rsa -C "C2243736958@qq.com"
```

然后一直按回车，直到命令执行完。

查看密钥

```bash
cat /root/.ssh/id_rsa.pub
```

然后你拿到了一串东西，如

```bash
ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAXXXXXXXXXXXXXXXXXXXXXXX
XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
XXXXXXsm8KWwcSyQ6qPJbN0BmzDp5nqNFc1k0XXXXX
XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXgToldCvhYY1Kw
SnTfJJJ8FAlqMWc2x5HkAHzGsXD4CdcH
XCOf3EjGYg2ZuHTmThf98d C2243736958@qq.com
```

## 3.设置GitHub

3.1打开https://github.com

3.2点击setting

![image-20231121145640398](https://percheung.github.io/blogImg/202401052303429.png)

3.3点击设置ssh

![img](https://percheung.github.io/blogImg/202401052304724.png)

3.4新建，如图，title随便写，这是给你自己看的，Key将你用`cat /root/.ssh/id_rsa.pub`拿到的东西全部粘贴进去。

![image-20231121150306728](https://percheung.github.io/blogImg/202401052303141.png)

## 4.拉取代码

> 这是我的私人仓库，你可以用别的仓库。复制ssh。

![image-20231121150551089](https://percheung.github.io/blogImg/202401052303301.png)

找个空文件夹，在里面执行下面的指令拉取代码

```bash
git clone git@github.com:PerCheung/tencent.git
```