点击返回[🔗我的博客文章目录](https://percheung.github.io/#/toc)
* 目录
{:toc}
<div onclick="window.scrollTo({top:0,behavior:'smooth'});" style="background-color:white;position:fixed;bottom:20px;right:40px;padding:10px 10px 5px 10px;cursor:pointer;z-index:10;border-radius:13%;box-shadow:0.5px 3px 7px rgba(0,0,0,0.3);"><img src="https://percheung.github.io/blogImg/backTop.png" alt="TOP" style="background-color:white;width:30px;"></div>


# Windows系统下使用docker安装elasticsearch和kibana

能使用这篇文章的前提是你的window电脑已经安装了docker，若没有docker，请参考我的一篇文章[Windows安装docker](https://percheung.github.io/blog/Windows安装docker)，把docker安装一下。若已经有了docker，请往下看。如果你下载的时候感觉拉取镜像很慢，可以参考我这一篇文章[Windows系统的docker设置阿里源](https://percheung.github.io/blog/Windows系统的docker设置阿里源)，把你的docker换成阿里镜像，就会拉取地快一点。

本文将用一个dockers-compose.yaml同时安装elasticsearch和kibana的7版本服务。

现在，打开`cmd`让我们开始吧！

## 1.新建文件夹

```bash
mkdir D:\App\docker\elasticsearch
```

## 2.新建elasticsearch.yml

```bash
mkdir D:\App\docker\elasticsearch\elasticsearch\config
```

```bash
vim D:\App\docker\elasticsearch\elasticsearch\config\elasticsearch.yml
```

内容如下

```yaml
cluster.name: "docker-cluster"
network.host: 0.0.0.0
http.cors.allow-origin: "*"
http.cors.enabled: true
http.cors.allow-headers: Authorization,X-Requested-With,Content-Length,Content-Type
```

## 3.新建kibana.yml

```bash
mkdir D:\App\docker\elasticsearch\kibana\config
```

```bash
vim D:\App\docker\elasticsearch\kibana\config\kibana.yml
```

内容如下

```yaml
#
# ** THIS IS AN AUTO-GENERATED FILE **
#

# Default Kibana configuration for docker target
server.host: "0.0.0.0"
server.shutdownTimeout: "5s"
elasticsearch.hosts: [ "http://elasticsearch:9200" ]
monitoring.ui.container.elasticsearch.enabled: true
i18n.locale: "zh-CN"
```

## 4.创建docker-compose.yaml

```bash
cd D:\App\docker\elasticsearch
```

```bash
vim docker-compose.yaml
```

为防止格式错乱可以用粘贴模式粘贴

```bash
:set paste
```

`docker-compose.yaml`内容如下

```yaml
version: '3.8'
services:
  elasticsearch:
    container_name: elasticsearch
    image: elasticsearch:7.17.1
    ports:
      - 9200:9200
      # 将宿主机的 9200 端口映射到容器的 9200 端口，用于访问 Elasticsearch HTTP API
      - 9300:9300
      # 将宿主机的 9300 端口映射到容器的 9300 端口，用于 Elasticsearch 节点之间的通信
    restart: no
    environment:
      - discovery.type=single-node
      # 设置 Elasticsearch 的发现类型为单节点
      - ES_JAVA_OPTS=-Xms1024m -Xmx1024m
      # 设置 Elasticsearch 的 Java 虚拟机选项，分配 256MB 堆内存
    volumes:
      - ./elasticsearch/logs:/usr/share/elasticsearch/logs
      # 挂载日志目录
      - ./elasticsearch/data:/usr/share/elasticsearch/data
      # 挂载数据目录
      - ./elasticsearch/plugins:/usr/share/elasticsearch/plugins
      # 挂载插件目录
      - ./elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml
      # 挂载配置文件
  kibana:
    container_name: kibana
    restart: no
    image: kibana:7.17.1
    ports:
      - 5601:5601
      # 将宿主机的 5601 端口映射到容器的 5601 端口，用于访问 Kibana Web 界面
    depends_on:
      - elasticsearch
      # 定义依赖关系，表示 Kibana 服务依赖于 Elasticsearch 服务
    volumes:
      - ./kibana/config/kibana.yml:/usr/share/kibana/config/kibana.yml
      # 挂载配置
```

## 5.启动

```bash
docker compose up -d
```

## 6.验证

### 6.1 elasticsearch

访问[http://localhost:9200/](http://localhost:9200/)

获得如下内容，证明es正常。

```json
{
  "name" : "9081ec6ae5f7",
  "cluster_name" : "docker-cluster",
  "cluster_uuid" : "UlpEtChjRaGeWb2l9aWWyA",
  "version" : {
    "number" : "7.17.1",
    "build_flavor" : "default",
    "build_type" : "docker",
    "build_hash" : "e5acb99f822233d62d6444ce45a4543dc1c8059a",
    "build_date" : "2022-02-23T22:20:54.153567231Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

### 6.2 kibana

访问[http://localhost:5601/app/dev_tools#/console](http://localhost:5601/app/dev_tools#/console)

执行

```json
GET _search
{
  "query": {
    "match_all": {}
  }
}
```

获得如下数据，证明kibana成功连接elasticsearch，并且成功启动。

```json
#! Elasticsearch built-in security features are not enabled. Without authentication, your cluster could be accessible to anyone. See https://www.elastic.co/guide/en/elasticsearch/reference/7.17/security-minimal-setup.html to enable security.
#! this request accesses system indices: [.apm-agent-configuration, .apm-custom-link, .kibana_7.17.1_001, .kibana_task_manager_7.17.1_001, .tasks], but in a future major version, direct access to system indices will be prevented by default
{
  "took" : 8,
  "timed_out" : false,
  "_shards" : {
    "total" : 5,
    "successful" : 5,
    "skipped" : 0,
    "failed" : 0
  },}}
......
```



