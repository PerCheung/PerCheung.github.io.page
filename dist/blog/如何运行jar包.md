点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
# 如何运行<img src="https://percheung.github.io/blogImg/jar.png" width="50px" alt="" />jar包

>运行jar包是每个Java程序员都要面临的问题，简单到`java -jar`，然而正由于过于简单，这里面很多技巧，却被大家忽略。

## 参数设置举例说明

`-server`：指定Java虚拟机以服务器模式运行，以获得更好的性能。
`-Dfile.encoding=UTF-8`：设置文件编码为UTF-8，确保正确处理文本数据。
`-XX:ActiveProcessorCount=8`：指定并行处理器的数目为8，以影响Java虚拟机的线程和CPU利用情况。
`-Xms2g`：指定Java虚拟机的最小堆内存为2g。
`-Xmx2g`：指定Java虚拟机的最大堆内存为2G。
`-XX:+UseParallelGC`：启用并行垃圾回收器，以提高垃圾回收的效率。
`-XX:+UseCompressedOops`：启用压缩指针，以减少内存消耗。
`-XX:ParallelGCThreads=8`：指定并行垃圾回收器的线程数为8。
`-XX:+DisableExplicitGC`：禁用显式的垃圾回收调用
`-XX:+HeapDumpOnOutOfMemoryError`：当发生堆溢出异常时，自动生成堆转储文件（Heap Dump）。
`-Djava.net.preferIPv4Stack=true`：该属性为 true 可确保 JVM 在具备 IPv6 支持的情况下仍然优先使用 IPv4。

## 设置技巧

1、 `-Xms`和`-Xmx`，在java8之后的jvm虚拟机里，应当设置为相等值，这样反而能达到最高效地分配堆空间，一般来说，参数写你服务器内存四分之一，比如你服务器是32g内存，这里就写8g。

以centos服务器为例，获取总内存指令
```bash
free -h | awk '/^Mem:/ {print $2}'
```
2、 `-XX:+UseParallelGC`启动并行垃圾回收后，注意要设置`-XX:ParallelGCThreads=8`，这个参数8，要根据你服务器cpu的核数来写，比如你是8核服务器，就写8。同理，`-XX:ActiveProcessorCount=8`也是如此，这将使得jar包是8核运行。

以centos服务器为例，获取cpu核数指令

```bash
nproc
```
3、`-XX:+HeapDumpOnOutOfMemoryError`是一个很重要的参数，特别是你部署在服务器上，发生堆内存泄露时，如果没有这个参数，就不会生成类似`java_pid74935.hprof`文件，你又如何分析堆内存溢出问题呢？关于堆内存溢出分析，我会在专栏[🔗java深入解剖](https://blog.csdn.net/weixin_43982359/category_11252566.html)里写的。感兴趣可以进去看看。

## 最终脚本

综上所述，运行jar包我们写成一个sh脚本。

```bash
#!/bin/bash

# 分配堆内存大小
heap_memory=2
# 获取线程数量
thread_count=$(nproc)

# 启动 Java 应用程序
#`-server`：指定Java虚拟机以服务器模式运行，以获得更好的性能。
#`-Dfile.encoding=UTF-8`：设置文件编码为UTF-8，确保正确处理文本数据。
#`-XX:ActiveProcessorCount=8`：指定并行处理器的数目为8，以影响Java虚拟机的线程和CPU利用情况。
#`-Xms2g`：指定Java虚拟机的最小堆内存为2g。
#`-Xmx4g`：指定Java虚拟机的最大堆内存为4g。
#`-XX:+UseParallelGC`：启用并行垃圾回收器，以提高垃圾回收的效率。
#`-XX:+UseCompressedOops`：启用压缩指针，以减少内存消耗。
#`-XX:ParallelGCThreads=8`：指定并行垃圾回收器的线程数为8。
#`-XX:+DisableExplicitGC`：禁用显式的垃圾回收调用
#`-XX:+HeapDumpOnOutOfMemoryError`：当发生堆溢出异常时，自动生成堆转储文件（Heap Dump）。
#`-Djava.net.preferIPv4Stack=true`：该属性为 true 可确保 JVM 在具备 IPv6 支持的情况下仍然优先使用 IPv4。
nohup java -server -Dfile.encoding=UTF-8 -XX:ActiveProcessorCount=$thread_count -Xms${heap_memory}G -Xmx${heap_memory}G -XX:+DisableExplicitGC -XX:+UseParallelGC -XX:+UseCompressedOops -XX:ParallelGCThreads=$thread_count -XX:+HeapDumpOnOutOfMemoryError -Djava.net.preferIPv4Stack=true -jar /server/newgonow/newgonow_server/newgonow_jar/newgeo-0.1.jar &
```
这个脚本，将把运行日志打到`nohup.out`，若运行时发生内存泄漏，也会产生`java_pid74935.hprof`，之后，我们就可以用内存泄漏工具进行分析了。而且由于它动态的获取服务器的线核数和内存，使得它具有强大的并行运行能力和垃圾回收能力，这保证它可以一直高效地运转。而且你可以放到任何服务器上，它都会尽量找到它的极限能力。