# The Example Nephele Program: Minimal

[Class 2](https://github.com/ctripcorp/nephele/blob/master/docs/cn/REQUIREMENT_CLASS.md)

## 目录

* [介绍](#介绍)

* [安装](#安装)

* [启动](#启动)

* [试用](#试用)

## 介绍

这里简短罗列一下Minimal的现状和未来的规划。
* 因为资源的原因，现阶段的Minimal还未引入上传功能，将在短期内引入。
* Minimal的日志是直接落到本地的，近期计划开发相应的日志同步工具，将日志同步到[ElasticSearch](https://www.elastic.co/products/elasticsearch)和[CAT](https://github.com/dianping/cat)。
* Minimal使用本地磁盘作为存储，如果要使用OSS或者S3之类的分布式存储的，应该转用Pluggable。
* 基本的图片处理功能都已经有了。

## 安装

**下载预编译文件**

(暂时还下载不到，近期吧)

* [Centos 7]()

* [Centos 6]()

* [Ubuntu]()

* [Windows]()


**下载源代码**

```bash
    git clone https://github.com/ctripcorp/nephele.git $GOPATH/src/github.com/ctripcorp/nephele
```

**下载依赖**

```bash
    cd nephele
    govendor sync
```

你可能需要前往[github.com/golang](https://github.com/golang)手动下载墙外包，并将它们移至正确路径。

**编译**

```bash
    cd example/minimal
    go build
```

## 启动

**运行**

[须知]执行下面的命令后，在用户根目录下会生成一个名为nephele的文件夹。

```bash
    ./minimal --open
```

如果你看到类似下面这样的打印信息，那么你已经成功启动了Nephele Minimal：

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code: gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /image/*imagepath         --> github.com/ctripcorp/nephele/service/handler.(*HandlerFactory).Build.func1 (3 handlers)
[GIN-debug] POST   /image                    --> github.com/ctripcorp/nephele/service/handler.(*HandlerFactory).Build.func1 (3 handlers)
[GIN-debug] GET    /healthcheck              --> github.com/ctripcorp/nephele/service/handler.(*HandlerFactory).Build.func1 (3 handlers)
```

同时，来到~/nephele/log下，你会看到Nephele Minimal生成的日志文件。

## 试用

(施工中，明天吧。)
