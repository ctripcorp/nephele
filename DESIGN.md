* [一、总体设计](#一、总体设计)
    * [1.1 基础模型](#基础模型)
    * [1.2 架构模块](#架构模块)
    * [1.3 各模块概要介绍](#各模块概要介绍)
* [二、图片服务设计](#二、图片服务设计)
* [三、性能和稳定性](#三、性能和稳定性)
* [四、监控相关](#四、监控相关)

# 一、总体设计

## 1.1 基础模型
![basic-architecture](https://github.com/ctripcorp/nephele/blob/master/doc/images/nephele_basic.png)

如果需要把图片服务暴露到公网，我们建议首要考虑引入CDN，因为国内CDN不仅能把回源率降到5%以下，这样做能极大减小对服务器的压力，并且也能解决用户最后一公里问题。

## 1.2 架构模块
![overall-architecture](https://github.com/ctripcorp/apollo/blob/master/doc/images/nephele_overall.png)

## 1.3 各模块概要介绍

### 1.3.1 Workbench

* 提供在线图片管理
* 提供在线图片编辑
* 图片测试和监控

### 1.3.2 ImageService

* 提供所有图片相关api
* 支持认证，限流，跨域，原图保护，样式解析等功能。
* 依赖imgcore模块对图片进行处理

### 1.3.3 ImageCore

* 支持裁剪，缩放，旋转，转换，滤镜，水印
* 底层依赖GraphicsMagick，ImageMagick库。

### 1.3.4 存储

* 支持多种存储，本地文件，FastDFS，阿里云OSS，AWS S3

### 1.3.5 监控

* 所有监控异常日志落本地，通过logagent转发到ES集群上。
* 通过Workbench查询聚合数据做定制展示和告警。


# 二、图片服务设计

下面展示了图片服务各模块之间的关系.

![overall-architecture](https://github.com/ctripcorp/apollo/blob/master/doc/images/nephele_module.png)

1. 图片服务包含多个过滤器，分别提供限流，认证，跨域，缓存时效，埋点等功能。
2. 采用异步模式处理图片，发送队列用于控制图片处理并行度，返回队列用于阻塞图片处理响应直到超时时间过期，启动快速失败想响应客户端。
3. 多协程处理图片，协程数相当于CPU核数，这样做能够最大程度有效利用CPU。
4. 核心图片处理模块提供对GM，IM组件的封装，并加入
5. 提供多种存储Driver，支持多个主流媒体存储服务。



# 三、性能和稳定性



# 四、监控相关

Nephele采用本地磁盘写日志的方式记录异常以及监控埋点，这样做的好处是通过格式化日志的方式可以编写不同Logagent实现向不同监控体系传输对应数据，能够对图片服务内部的代码侵入性做到最低，并且使用基于本地持久化的技术也能够最大程度保证日志数据不丢，在携程内部，我们基于CAT做性能调优和监控，开源版本我们将采用更通用的ES作为监控手段。
