点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 将<img src="https://percheung.github.io/blogImg/Python.png" width="50px" alt="C" />python打包成exe文件

> 以上一篇文章🔗[用python删除重复文件并放入回收站](https://percheung.github.io/blog/%E7%94%A8python%E5%88%A0%E9%99%A4%E9%87%8D%E5%A4%8D%E6%96%87%E4%BB%B6%E5%B9%B6%E6%94%BE%E5%85%A5%E5%9B%9E%E6%94%B6%E7%AB%99)为例，演示了如何用python写一个删除重复文件并放入回收站的功能代码，但是每次都要cmd执行一下`python 删除重复文件.py`也是非常惹人烦的。有没有办法让python像go语言一样，能打包成exe文件，双击执行即可呢？

## 1.安装PyInstaller

要将Python脚本打包成可执行文件（.exe），可以使用第三方工具来实现。以下是一种常用的方法，使用PyInstaller工具将Python脚本打包成独立的可执行文件。

请按照以下步骤操作：

1. 确保您已经安装了Python和pip，并将它们添加到系统的环境变量中。

2. 打开命令提示符或终端，并使用pip安装PyInstaller：

   ```shell
   pip install pyinstaller
   ```

3. 验证是否安装成功。

   ```bash
   pyinstaller -v
   ```

## 2.配置pyinstaller到环境变量

对部分python用户来说，pip安装后发现报错pyinstaller找不到，是因为我们的python环境变量没有配置好，python需要你配置到`${python安装目录}/Scripts`才行。可以参考下面步骤配置一下你的环境变量。

在Windows系统中，可以按照以下步骤进行操作：

- 打开控制面板，并进入"系统和安全" -> "系统" -> "高级系统设置"。
- 在"高级"选项卡下，点击"环境变量"按钮。
- 在"系统变量"部分，找到名为"Path"的变量，并点击"编辑"。
- 在编辑环境变量窗口中，点击"新建"，然后添加Python和PyInstaller所在的目录路径，例如：`C:\PythonXX\Scripts`（其中`XX`表示您的Python版本号）。

## 3.打包

使用命令打包。

```bash
pyinstaller -i Python.ico -F 删除重复文件.py
```

`pyinstaller`是一个用于将Python脚本打包成可执行文件的工具。在给出的命令中，使用了一些选项和参数来配置打包过程。让我解释每个选项的含义：

- `-i`: 这个选项指定了可执行文件的图标。图标文件应该是一个.ico格式的文件。我看有些教程用的png文件也可以，我没有试过，可以尝试一下。对了，`Python.ico`是需要你自己准备的。
- `-F`: 这个选项指定了打包成单个可执行文件。使用了`-F`选项后，`pyinstaller`会将所有依赖项和脚本都打包到一个独立的可执行文件中，使得分发更加方便。
- `删除重复文件.py`的代码我写在上一篇文章🔗[用python删除重复文件并放入回收站](https://percheung.github.io/blog/%E7%94%A8python%E5%88%A0%E9%99%A4%E9%87%8D%E5%A4%8D%E6%96%87%E4%BB%B6%E5%B9%B6%E6%94%BE%E5%85%A5%E5%9B%9E%E6%94%B6%E7%AB%99)里，这里你可以用自己的python文件。

然后就会在同级别目录下产生一个文件夹，文件夹名字叫`dist`，里面会含有一个名字叫`删除重复文件.exe`，之后，我们就只需要将`删除重复文件.exe`拖拽到你想删除重复文件的目录下，双击执行即可。