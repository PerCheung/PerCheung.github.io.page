点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# ffmpeg使用手册

> 本教程持续更新，我学到多少就更新多少

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/FFmpeg.png" width="30%" alt="ffmpeg" />
</div>

## ffmpeg是什么

FFmpeg是一个开放源代码的自由软件，可以执行音频和视频多种格式的录影、转换、串流功能，包含了libavcodec——这是一个用于多个项目中音频和视频的解码器函式库，以及libavformat——一个音频与视频格式转换函式库。

具体请看[https://zh.wikipedia.org/wiki/FFmpeg](https://zh.wikipedia.org/wiki/FFmpeg)

## 指令总结

### 1.查看ffmpeg版本

```bash
ffmpeg -version
```

### 2.mkv转mp4

```bash
ffmpeg -i input.mkv -vcodec copy -acodec copy out.mp4
```

### 3.裁剪 .mkv 视频

```bash
ffmpeg -i input.mkv -ss [start] -t [duration] -c copy output.mkv
```

- `input.mkv` 是你要裁剪的原始视频文件。
- `[start]` 是你要开始裁剪的时间，格式为 `HH:MM:SS`。
- `[duration]` 是你要裁剪的持续时间，格式也是 `HH:MM:SS`。
- `output.mkv` 是裁剪后的新视频文件。

### 4.不调节帧率，尽可能保证原视频质量的情况下将原始视频压缩

要使用FFmpeg将视频文件的大小减小，可以进行以下操作：

```bash
ffmpeg -i input.mp4 -crf 23 -preset medium output.mp4
```

这个命令将会将输入视频文件`input.mp4`的占用体积减小，并生成输出视频文件`output.mp4`。

#### 4.1 crf

在FFmpeg中，`-crf`参数是用于控制视频编码的质量与压缩比之间的权衡的参数。CRF（Constant Rate Factor）是一种基于目标质量的恒定比特率控制方法，它允许你设置输出视频的质量水平。CRF值越小，视频的质量越高，但文件大小也会增加。

在一般情况下，CRF值的范围是0到51。不同的范围值对应不同的质量和压缩级别。以下是一些常用的CRF值和对应的特征：

- 0：无损模式，产生非常大的文件，质量最好。
- 18-28：一般推荐范围，较高的质量和文件大小平衡。
- 23：默认值，适用于大多数情况，提供良好的质量和合理的文件大小。
- 29-51：较低质量的范围，适用于低比特率压缩或要求较小文件大小的情况。

你可以根据自己的需求和偏好选择合适的CRF值。较低的CRF值会提供更好的质量，但会增加文件大小。较高的CRF值会减小文件大小，但可能导致一些质量损失。

请注意，不同的视频编解码器可能对CRF值的范围有所不同，具体取值范围和效果可能会有所差异。建议在实际应用中进行测试和调整，以找到最适合你需求的CRF值。

#### 4.2 preset

在FFmpeg中，`-preset`参数用于设置视频编码器的预设（preset）。预设是一组预定义的编码参数，可以影响编码速度和输出质量之间的权衡。不同的预设提供了不同的速度和质量选项。

以下是一些常用的预设选项：

- `ultrafast`：最快的编码速度，但质量较低。
- `superfast`：非常快的编码速度，稍微降低了一些质量。
- `veryfast`：非常快的编码速度，质量稍低。
- `faster`：快速的编码速度，质量略有降低。
- `fast`：相对较快的编码速度，质量较好。
- `medium`：平衡了速度和质量，是默认预设。
- `slow`：较慢的编码速度，质量更好。
- `slower`：更慢的编码速度，质量更高。
- `veryslow`：非常慢的编码速度，提供最好的质量。

这些预设选项提供了不同的速度和质量平衡。如果你希望编码速度更快，可以选择速度较快的预设，但可能会牺牲一些质量。如果你更注重输出质量，可以选择速度较慢的预设，但编码速度会变慢。

你可以根据你的需求选择适合的预设选项。要指定预设，请使用`-preset`参数后跟预设名称，例如`-preset medium`。

请注意，不同的视频编解码器可能支持不同的预设选项，具体的预设选项和效果可能会有所不同。建议在实际应用中进行测试和调整，以找到最适合你需求的预设选项。

### 5.调节视频帧率

```bash
ffmpeg -i input.mp4 -r 24 output.mp4
```

`-i` 选项指定了输入文件的名称为 `input.mp4`，`-r` 选项指定了输出文件的帧速率为 24 帧每秒（fps），并且 `output.mp4` 是输出文件的名称。

### 6.调节帧率，尽可能保证原视频质量的情况下将原始视频压缩

> 将5和6的总结起来

```bash
ffmpeg -i input.mp4 -r 24 -crf 23 -preset slower output.mp4
```

效果是将我原先10兆大小的视频文件，压缩到了5兆。

### 7.裁剪视频

```bash
ffmpeg -ss 00:00:00 -t 00:13:55 -i input.mp4 -vcodec copy -acodec copy out.mp4
```

这是一个使用 FFmpeg 的命令行指令，用于剪辑视频文件。让我为你解释每个参数的含义：

- `-ss 00:00:00`：这个参数表示要从输入视频文件的起始位置开始剪辑。在这个例子中，剪辑的起始时间被设置为 00:00:00，也就是视频的开头。

- `-t 00:13:55`：这个参数表示要剪辑的视频时长。在这个例子中，剪辑的时长被设置为 00:13:55，也就是 13 分钟 55 秒。

- `-i input.mp4`：这个参数指定输入的视频文件。在这个例子中，输入文件的名称是 `input.mp4`。请确保该文件存在并位于当前工作目录中，或者提供完整的文件路径。

- `-vcodec copy`：这个参数表示视频流的编码方式，`copy` 表示无需重新编码，直接复制原始视频流。这样可以保持视频质量而不进行重新压缩。

- `-acodec copy`：这个参数表示音频流的编码方式，`copy` 表示无需重新编码，直接复制原始音频流。这样可以保持音频质量而不进行重新压缩。

- `out.mp4`：这个参数指定输出的视频文件名。在这个例子中，输出文件将被命名为 `out.mp4`。请根据需要提供所需的输出文件名和路径。

综合起来，这个命令的作用是从输入视频文件的起始位置开始剪辑指定时长的视频，并将剪辑后的视频保存为输出文件，同时保持视频和音频的原始编码方式和质量。
