点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
# ffmpeg使用

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