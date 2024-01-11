点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({ top: 0, behavior: 'smooth' });" style="background-color: transparent; position: fixed; bottom: 20px; right: 40px; padding: 10px 10px 5px 10px; cursor: pointer; z-index: 10; border-radius: 13%; box-shadow: 0.5px 3px 7px rgba(0, 0, 0, 0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color: transparent; width: 30px;"></div>

# ffmpeg使用

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/video.png" width="30%" alt="ffmpeg" />
</div>

## mkv转mp4

```bash
ffmpeg -i input.mkv -vcodec copy -acodec copy out.mp4
```

## 裁剪 .mkv 视频

```bash
ffmpeg -i input.mkv -ss [start] -t [duration] -c copy output.mkv
```

- `input.mkv` 是你要裁剪的原始视频文件。
- `[start]` 是你要开始裁剪的时间，格式为 `HH:MM:SS`。
- `[duration]` 是你要裁剪的持续时间，格式也是 `HH:MM:SS`。
- `output.mkv` 是裁剪后的新视频文件。