点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>

# 安装Java的完整jdk并且配置JAVA_HOME

<div style="text-align: center;">
  <img src="https://percheung.github.io/blogImg/java.png" width="20%" alt="java" />
</div>

## 1.安装

```bash
sudo apt install openjdk-8-jdk
```

## 2.查找Java路径

```bash
update-alternatives --display java
```

`/usr/lib/jvm/java-8-openjdk-amd64`就是jdk的实际位置

## 3.测试Java是否安装成功

```bash
java -version
```

## 4.配置JAVA_HOME

编辑环境变量

```bash
sudo vim /etc/profile
```

末尾写上

```bash
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64
export JRE_HOME=$JAVA_HOME/jre
export CLASSPATH=$JAVA_HOME/lib:$JRE_HOME/lib:$CLASSPATH
export PATH=$JAVA_HOME/bin:$JRE_HOME/bin:$PATH
```

刷新

```bash
source /etc/profile
```

验证

```bash
echo $JAVA_HOME
```

