点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# ImageMagick使用手册

> 本教程持续更新，我学到多少就更新多少

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/wizard.png" width="50%" alt="ImageMagick" />
</div>

## ImageMagick是什么

ImageMagick是一个用于查看、编辑位图文件以及进行图像格式转换的开放源代码软件套装。它可以读取、编辑超过100种图帧式。ImageMagick以ImageMagick许可证 （页面存档备份，存于互联网档案馆）（一个类似BSD的许可证）发布。

具体请看[https://zh.wikipedia.org/wiki/ImageMagick](https://zh.wikipedia.org/wiki/ImageMagick)

## 指令总结

### 1.查看版本

```bash
magick -version
```

### 2.官网验证指令解读

```bash
magick logo: logo.gif
magick identify logo.gif
magick logo.gif win:
```

这三个命令都是与ImageMagick相关的命令，用于图像处理和操作。以下是对每个命令的解释。

1. `magick logo: logo.gif`

这个命令是用来创建一个名为`logo.gif`的GIF图像文件，其中图像内容为空。`logo:`是一个特殊的标识符，它表示创建一个指定尺寸的空白图像。这里的命令将创建一个名为`logo.gif`的空白图像文件。

2. `magick identify logo.gif`

这个命令用于获取图像文件的详细信息，例如图像的尺寸、格式、颜色等。`identify`是ImageMagick的命令，它可以识别图像文件并提供有关图像的元数据。在这里，命令将显示关于`logo.gif`图像文件的详细信息。

3. `magick logo.gif win:`

这个命令是用来将图像文件从一种格式转换为另一种格式。`win:`是一种特殊的标识符，表示将图像文件转换为Windows图标文件（ICO格式）。在这里，命令将把`logo.gif`图像文件转换为Windows图标文件格式。

### 3.svg转png

使用ImageMagick的`magick`命令可以将SVG转换为PNG，并将背景设置为透明。

```bash
magick -background none -format png FFmpeg.svg FFmpeg.png
```

这个命令会将`FFmpeg.svg`转换为PNG文件`FFmpeg.png`，并将背景设置为透明。

*回想当初我为了能实现`svg转png`，还专门写过一个微型的spring boot项目，还写了一篇博客[svg转png](https://percheung.github.io/blog/svg转png)，没想到对这个工具来说只是一行代码的事情，被按在地上狠狠摩擦。*

### 4.png转ico

```bash
magick convert macaw.png -resize 256x256 macaw.ico
```

这个命令将把`macaw.png`转换为256x256像素的ICO文件`macaw.ico`。可以用`-resize`根据需要调整输出ICO文件的尺寸。

### 5.将指定文件夹中的所有PNG图像转换为ICO图标

```bash
magick mogrify -path "C:\Users\Peter\Pictures\icon" -format ico -resize 256x256 "C:\Users\Peter\Pictures\图标\*.png"
```

这个命令使用ImageMagick的`mogrify`工具对指定文件夹中的PNG图像进行批量处理，并将它们转换为ICO格式的图标文件。下面是每个参数的详细解释：

- `magick mogrify`：这是用于批量处理图像的ImageMagick命令。它可以对指定的图像进行一系列的转换和操作。

- `-path "C:\Users\Peter\Pictures\icon"`：这个参数指定了输出图标文件的路径。在这个例子中，图标文件将被输出到`C:\Users\Peter\Pictures\icon`文件夹中。你可以根据需要修改此路径。

- `-format ico`：这个参数指定输出图标文件的格式为ICO。ICO是一种常用的图标文件格式，通常用于Windows操作系统。

- `-resize 256x256`：这个参数指定了调整图像尺寸的大小为256x256像素。它将调整输入图像的大小以适应图标的尺寸要求。你可以根据需要调整此值。

- `"C:\Users\Peter\Pictures\图标\*.png"`：这个参数指定了待处理的PNG图像文件的路径和文件名模式。在这个例子中，它表示`C:\Users\Peter\Pictures\图标`文件夹下的所有PNG图像文件。`*.png`是一个通配符，表示匹配该文件夹下的所有PNG文件。

通过这个命令，ImageMagick将在指定的文件夹中查找所有的PNG图像文件，将它们调整为指定的尺寸，并将其转换为ICO格式的图标文件，最后输出到指定的路径中。

*该命令让我真切感受到该工具的强大*

### 6.裁剪图片上半部分

只保留图片上面的50%

```bash
magick input.jpg -crop 100%x50%+0+0 output.jpg
```

`-crop 100%x50%+0+0`参数告诉ImageMagick裁剪输入图片，保留宽度为100%、高度为50%的部分，并且从左上角开始裁剪。执行这个命令后，你将得到一个名为`output.jpg`的图片文件，其中只包含输入图片上半部分的内容。

### 7.裁剪图片下半部分

只保留图片下面的50%

```bash
magick input.jpg -gravity south -crop 100%x50% output.jpg
```

这个命令将会截取`input.jpg`图片的下半部分，并将结果保存为`output.jpg`。

- `input.jpg`：你要截取的原始图片的文件名。
- `-gravity south`：这个选项将确定截取的起点位置为图片的底部。
- `-crop 100%x50%`：这个选项将截取图片的宽度为100%（保持原始宽度），高度为50%（即下半部分）。
- `output.jpg`：你要保存截取结果的文件名。