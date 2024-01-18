点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({ top: 0, behavior: 'smooth' });" style="background-color: transparent; position: fixed; bottom: 20px; right: 40px; padding: 10px 10px 5px 10px; cursor: pointer; z-index: 10; border-radius: 13%; box-shadow: 0.5px 3px 7px rgba(0, 0, 0, 0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color: transparent; width: 30px;"></div>

# 更新<img src="https://percheung.github.io/blogImg/Ubuntu.png" width="150px" alt="" />并同步网络时间

## 1.更新Ubuntu

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/Ubuntu.png" width="20%" alt="ubuntu" />
</div>


### 1.1 查看ubuntu版本和详细信息

要查看Ubuntu的版本和详细信息，您可以使用以下命令之一：

1. 使用lsb\_release命令：

```bash
lsb_release -a
```

该命令会显示您的Ubuntu版本、发行版名称和发行版代码名称。

2. 使用cat命令查看`/etc/lsb-release`文件：

```bash
cat /etc/lsb-release
```

这个命令会显示与lsb\_release命令相同的信息。

3. 使用uname命令查看内核版本：

```bash
uname -r
```

这个命令会显示当前正在运行的内核版本。

4. 使用lsb\_release命令查看版本号：

```bash
lsb_release -r
```

该命令将只显示Ubuntu的版本号。

任何一个命令都可以提供您所需的Ubuntu版本和详细信息。

### 1.2 创建root用户

```bash
sudo passwd root
```

### 1.3 更新Ubuntu

在Ubuntu中，使用的包管理器是APT（Advanced Package Tool），您可以使用以下命令来更新您的系统：

```bash
sudo apt update
```

上述命令将更新软件包列表，以获取最新的可用软件包信息。

接下来，您可以使用以下命令来升级已安装的软件包：

```bash
sudo apt upgrade
```

该命令将升级您系统中已安装的所有软件包到它们的最新版本。

如果您希望自动确认软件包更新而不需要手动确认，请使用以下命令：

```bash
sudo apt update && sudo apt upgrade -y
```

请注意，系统更新可能需要一些时间，并且在更新过程中可能会提示您输入密码以获取管理员权限（使用`sudo`命令）。确保您具有适当的权限来执行这些操作。

另外，如果您还想更新Ubuntu系统的发行版版本，可以使用以下命令：

```bash
sudo do-release-upgrade
```

该命令将引导您完成发行版升级过程，例如从Ubuntu 20.04 LTS升级到Ubuntu 22.04 LTS。在执行此命令之前，请务必备份重要的数据和配置文件，并确保您的系统已经备份和更新了所有重要的软件包。

## 2.同步网络时间

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/time.png" width="10%" alt="time" />
</div>

安装ntpdate

```bash
sudo apt-get install ntpdate
```

停止ntp

```bash
sudo service ntp stop
```

设置系统时间与网络时间同步

```bash
sudo ntpdate cn.pool.ntp.org
```

设置时区亚洲上海

```bash
sudo cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```

启动 NTP 服务

```bash
sudo service ntp start
```

开机自启动

```bash
sudo systemctl enable ntp
```

验证时间同步

```bash
ntpq -p
```

将时间更新到硬件

```bash
sudo hwclock --systohc
```

查看时间

```bash
date
```