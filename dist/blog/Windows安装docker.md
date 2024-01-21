点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# Windows安装docker

> *docker已经成为一种标准*

然而，偏偏docker对Linux和macOS都提供了良好的支持，但对win而言，就不太友好了。但是docker太好用了，但是我又得用windows系统，怎么办呢？

本文安装docker，将使用WSL2，而不是vmeare和Hyper-V，也不需要你另外安装Linux系统（如Ubuntu）。查阅资料均来自官网文档（微软官网和docker官网）。下面让我们开始吧！

## 1.搭建WSL2环境

> 良好的编程习惯就是，做事前要先看看官方是怎么说的。

### 1.1 docker官网对环境要求的描述

#### [System requirements](https://docs.docker.com/desktop/install/windows-install/#system-requirements)

WSL 2 backend Hyper-V backend and Windows containers

---

- WSL version 1.1.3.0 or later.
- Windows 11 64-bit: Home or Pro version 21H2 or higher, or Enterprise or Education version 21H2 or higher.
- Windows 10 64-bit:
  
    - We recommend Home or Pro 22H2 (build 19045) or higher, or Enterprise or Education 22H2 (build 19045) or higher.
    - Minimum required is Home or Pro 21H2 (build 19044) or higher, or Enterprise or Education 21H2 (build 19044) or higher.
- Turn on the WSL 2 feature on Windows. For detailed instructions, refer to the [Microsoft documentation](https://docs.microsoft.com/en-us/windows/wsl/install-win10).
- The following hardware prerequisites are required to successfully run WSL 2 on Windows 10 or Windows 11:
  
    - 64-bit processor with [Second Level Address Translation (SLAT)](https://en.wikipedia.org/wiki/Second_Level_Address_Translation)
    - 4GB system RAM
    - Enable hardware virtualization in BIOS. For more information, see [Virtualization](https://docs.docker.com/desktop/troubleshoot/topics/#virtualization).

---

让我翻译一下：

**系统要求**

WSL 2后端Hyper-V后端和Windows容器

---

- WSL 1.1.3.0版或更高版本。

- Windows 11 64位：家庭版或专业版21H2或更高版本，或企业版或教育版21H2及更高版本。

- Windows 10 64位：

	- 我们推荐Home或Pro 22H2（19045版本）或更高版本，或Enterprise或Education 22H2（19095版本）或更低版本。

	- 最低要求为Home或Pro 21H2（19044版）或更高版本，或Enterprise或Education 21H2（19094版）或更低版本。

- 在Windows上打开WSL 2功能。有关详细说明，请参阅[Microsoft文档](https://docs.microsoft.com/en-us/windows/wsl/install-win10)。

- 在Windows 10或Windows 11上成功运行WSL 2需要以下硬件先决条件：
	- 带[二级地址转换（SLAT）]的64位处理器(https://en.wikipedia.org/wiki/Second_Level_Address_Translation)

	- 4GB系统RAM

	- 在BIOS中启用硬件虚拟化。有关更多信息，请参阅[虚拟化](https://docs.docker.com/desktop/troubleshoot/topics/#virtualization)。

看嘛！你的电脑是不是win10新版本或者win11呢？如果是的话你就可以安装了，然后请记住**在BIOS中启用硬件虚拟化**，这个就不赘述了，准备好了吗？

> so，让我们去https://docs.microsoft.com/en-us/windows/wsl/install-win10看看，先把WSL 2安装上再说！

### 1.2 安装WSL2

> 你会去翻看官方文档吗？反正我已经看了，放到博客里给你们看看要怎么做。

**下面这个指令可不要执行啊！！！**

#### Install WSL command

You can now install everything you need to run WSL with a single command. Open PowerShell or Windows Command Prompt in **administrator** mode by right-clicking and selecting "Run as administrator", enter the wsl --install command, then restart your machine.

PowerShell Copy

```bash
wsl --install
```

看意思，就是说，用`管理员身份`运行的`Windows PowerShell`里执行一下`wsl --install`就可以了。

**but**，nono，你如果看完整官网的文档，就知道这样安装顺便给你还送了一个Ubuntu！微软真的有多不靠谱要多不靠谱。

*服务用户是你的谎言，背刺用户是微软的真相。*

#### 1.2.1 正确的操作

所以正确的方法是，打开控制面板，选择，控制面板\程序，如图。

![image-20240119151048634](https://percheung.github.io/blogImg/image-20240119151048634.png)

然后重启电脑即可。

### 1.3 将WSL 2设置为默认值

现在其实是wsl1，而不是wsl2，所以请用`管理员身份`运行的`Windows PowerShell`里执行一下。

```bash
wsl --set-default-version 2
```

> 至此，WSL 2环境我已经全部准备完毕！

## 2.开始安装docker

> 我比较喜欢官网的东西，我就直接在docker官网下载安装包了，我没有百度云或者什么可以提供。

访问地址：[https://docs.docker.com/desktop/install/windows-install](https://docs.docker.com/desktop/install/windows-install)

![image-20240119151100135](https://percheung.github.io/blogImg/image-20240119151100135.png)

下载完成后，我们获得一个`Docker Desktop Installer.exe`。这个网页下面就是安装教程，如果大家不喜欢看官网的，下面就看我是怎么做的吧。

### 2.1 双击运行Docker Desktop Installer.exe

下一步，下一步...好了，结束了哈哈哈。

## 3.使用docker

> 让我们看看能不能用

cmd执行

```bash
docker -v
```

可以看到docker已经有了。

```bash
docker compose version
```

可以看到docker compose也有了。

老样子，跑个hello world。

```bash
docker run hello-world
```

太美妙了。

![image-20240119160839615](https://percheung.github.io/blogImg/image-20240119160839615.png)

![image-20240119153418799](https://percheung.github.io/blogImg/image-20240119153418799.png)

今后，岂不是可以直接window系统下，用docker跑redis，mysql8，mysql5.7，nginx等等，再也不需要学那些繁琐的window安装软件教程了，只需要学docker compose安装教程了！
