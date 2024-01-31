点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# Windows安装Anaconda教程

> 本文是我实践后写的，无脑跟随安装即可

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/anaconda.png" width="25%" alt="window" />
  <p></p>
  <p>
  <strong style="color:#42b029">在我看来，Anaconda的图标如同一只灵蛇咬住了自己的尾巴，优美而神秘</strong>
  </p>
</div>

**全称：Anaconda**

**发音：安尼康达**

**含义：蟒蛇**

<p style="color:#42b029">误解了，anaconda是森楠的意思，Python才是蟒蛇的意思。</p>

> **Anaconda**是一个[开源](https://zh.wikipedia.org/wiki/开放源代码)[Python](https://zh.wikipedia.org/wiki/Python)和[R语言](https://zh.wikipedia.org/wiki/R语言)的发行版本，用于[计算科学](https://zh.wikipedia.org/wiki/计算科学)（[数据科学](https://zh.wikipedia.org/wiki/数据科学)、[机器学习](https://zh.wikipedia.org/wiki/机器学习)、[大数据处理](https://zh.wikipedia.org/wiki/大数据)和[预测分析](https://zh.wikipedia.org/wiki/预测分析)），Anaconda致力于简化[软件包管理系统](https://zh.wikipedia.org/wiki/软件包管理系统)和部署。Anaconda透过[Conda](https://zh.wikipedia.org/wiki/Conda)进行[软件包管理](https://zh.wikipedia.org/wiki/软件包管理系统)，并拥有许多适用于[Windows](https://zh.wikipedia.org/wiki/Microsoft_Windows)、[Linux](https://zh.wikipedia.org/wiki/Linux)和[MacOS](https://zh.wikipedia.org/wiki/MacOS)的[数据科学](https://zh.wikipedia.org/wiki/数据科学)[软件包](https://zh.wikipedia.org/wiki/软件包)。

## 1.下载

**Anaconda**下载地址是🔗[https://www.anaconda.com/download](https://www.anaconda.com/download)

进入页面后直接点击`Download`即可。

然后，你会获得`Anaconda3-Windows.exe`。

## 2.安装

双击exe即可，next，为所有人安装。

![image-20240122202855920](https://percheung.github.io/blogImg/image-20240122202855920.png)

我就直接用默认位置了，**请记住你的位置**，比如我的，`C:\ProgramData\anaconda3`

![image-20240122203124133](https://percheung.github.io/blogImg/image-20240122203124133.png)

一路全部勾选即可。我讲一下三个勾选是做什么的，以免你不放心。

1. 第一个按钮，在开始菜单创建快捷路径（有没有快捷路径，无关紧要）。
2. 第二个按钮，配置环境变量（我试过了，是不生效的，看了很多教程，有说一定要勾选的，也有说一定别勾选的，实践了一下，勾不勾完全没有影响到电脑配置，所以这里看心情就好，按理来说，配置环境变量是有意义的，应该勾选上才对）。
3. 第三个按钮，安装完成后清除垃圾文件以及垃圾缓存（这个安装包还挺贴心的）。

![image-20240122203321224](https://percheung.github.io/blogImg/image-20240122203321224.png)

## 3.配置环境变量

一路next，一路全部勾选，无脑安装即可，不会出错的，安装完成后，我发现很多小白说我的python呢？问题的答案是环境变量没有配置好。

> 这里参考一下环境变量是怎么编辑的，看懂了下面我就开始了。

在Windows系统中，可以按照以下步骤进行操作：

- 打开控制面板，并进入"系统和安全" -> "系统" -> "高级系统设置"。（右击此电脑进属性就能找到，这里chatGPT说的有点啰嗦了）
- 在"高级"选项卡下，点击"环境变量"按钮。
- 在"系统变量"部分，找到名为"Path"的变量，并点击"编辑"。
- 在编辑环境变量窗口中，点击"新建"。

**下面的步骤是必须的**

### 3.1 ANACONDA_HOME

废话一下，我看很多教程说，环境变量直接在PATH里写下面这套东西。

```
X:\Anaconda 
X:\Anaconda\Scripts 
X:\Anaconda\Library\mingw-w64\bin
X:\Anaconda\Library\usr\bin 
X:\Anaconda\Library\bin
```

这样写当然没什么问题，只不过太不优美了。观察一下你的环境变量。

![image-20240122205048250](https://percheung.github.io/blogImg/image-20240122205048250.png)

标准的写法，都会先在这里写上软件的HOME。

#### 3.1.1 点击新建

![image-20240122205231694](https://percheung.github.io/blogImg/image-20240122205231694.png)

#### 3.1.2 写上键和值

![image-20240122205516430](https://percheung.github.io/blogImg/image-20240122205516430.png)

键名直接用我的，复制一下即可，蟒蛇的家。

```
ANACONDA_HOME
```

至于变量值还记得应该写什么吗？对了，正是步骤2的时候强调过的，请记住你的安装路径，填写你自己的安装路径即可。

然后记得点击确定，下面我们开始写Path。

### 3.2 Path

在系统变量里，找到Path，点击Path，然后点击编辑。Path很重要，千万别编辑里瞎写哈。

![image-20240122205707877](https://percheung.github.io/blogImg/image-20240122205707877.png)

你可以用这个写法，复制粘贴五遍。（先别急着操作😂，先往下再看一点）

```
%ANACONDA_HOME%
```

```
%ANACONDA_HOME%\Scripts
```

```
%ANACONDA_HOME%\Library\mingw-w64\bin
```

```
%ANACONDA_HOME%\Library\usr\bin
```

```
%ANACONDA_HOME%\Library\bin
```

![image-20240122210009961](https://percheung.github.io/blogImg/image-20240122210009961.png)

而我选择这样做，点击`编辑文本`。

![image-20240122210332670](https://percheung.github.io/blogImg/image-20240122210332670.png)

然后直接将我下面给的内容粘贴到尾部，一路点击确定即可。

```
%ANACONDA_HOME%;%ANACONDA_HOME%\Scripts;%ANACONDA_HOME%\Library\mingw-w64\bin;%ANACONDA_HOME%\Library\usr\bin;%ANACONDA_HOME%\Library\bin;
```

![image-20240122210637796](https://percheung.github.io/blogImg/image-20240122210637796.png)

然后环境变量就配置好了，ANACONDA_HOME的意义正是为Path之上定义一个变量，Path里就可以用这个统一变量往下写了，而不是同样的事情写五遍，而且后期更换软件`Anaconda`位置的时候，我们只需要改一下ANACONDA_HOME的值就可以了，而不需要你跑到Path里改五遍，这就是编程思维。

## 4.验证

做了这么多了，看看结果怎么样。

```bash
python --version
```

```bash
conda --version
```

![image-20240122211212969](https://percheung.github.io/blogImg/image-20240122211212969.png)

小蛇和蟒蛇我们都有了，这里已经大功告成！

## 5.扩展知识

> 不感兴趣可以不看了，这部分与安装已经没有关系了

![image-20240122211454957](https://percheung.github.io/blogImg/image-20240122211454957.png)

我们可以看到，蟒蛇安装的软件一共就这6个，我用了一下，说说看法。

1. Anaconda Navigator：是可视化环境工具箱，一般没用，真正的高手都用命令，可视化的东西对高手来说效率很低。（chatGPT对我的纠正： Anaconda Navigator是一个可视化的环境和包管理工具，它允许用户轻松地管理不同的Python环境和安装的软件包。虽然一些有经验的用户可能更喜欢使用命令行来管理环境，但对于初学者或者需要可视化管理的用户来说，Anaconda Navigator提供了一个很好的界面）。
2. Jupyter Notebook：局域网协同办公的云编辑器，看ui已经过时了（chatGPT对我的纠正：Jupyter Notebook是一个交互式的笔记本工具，可以用于数据分析、可视化和机器学习等任务。它的优势在于能够将代码、文本和图像整合在一个文档中，并且支持多种语言。虽然它可能不是最适合大规模协同办公的工具，但在数据分析和探索性编程方面仍然非常有用。）。
3. Spyder：本地python IDE，用起来感觉被pycharm吊打看起来，废了（chatGPT对我的纠正：Spyder是一个科学计算和数据分析的Python IDE，它提供了类似于MATLAB的界面和功能。虽然PyCharm是另一个流行的Python IDE，但Spyder在数据分析和科学计算方面有其独特的优势，例如集成了IPython控制台和数据变量查看器等功能）。
4. Reset Spyder Settings：一个重置Spyder指令脚本（chatGPT对我的纠正：Reset Spyder Settings是用于重置Spyder IDE设置的指令，可以帮助用户恢复默认设置或解决一些配置问题）。
5. Anaconda Prompt：Anaconda发行版中附带的一个命令行终端（或命令提示符）工具。它提供了一个特定的环境，用于管理和运行Anaconda中的Python环境和软件包（我还是喜欢一个cmd搞定所有）。
6. Anaconda PowerShell Prompt：Anaconda发行版提供的另一个命令行工具，它基于Windows的PowerShell环境。与Anaconda Prompt类似，Anaconda PowerShell Prompt专门用于管理和运行Anaconda中的Python环境和软件包。（不可能比Windows PowerShell好用吧？）
