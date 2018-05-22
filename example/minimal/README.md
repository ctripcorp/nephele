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

* [Centos 7](http://file.c-ctrip.com/files/3/nep/zip/82f/2de/e47/671aab3821dd4dab9aec14bcad8c2fd4.zip)

* [Centos 6](http://file.c-ctrip.com/files/3/nep/zip/88d/e2c/185/2158e329294346f4870aec40c630b4f9.zip)

* [Ubuntu](http://file.c-ctrip.com/files/3/nep/zip/300/cd7/235/a915c23e981e42a590633b7c273e5385.zip)

* [Windows](http://file.c-ctrip.com/files/3/nep/zip/7e5/993/749/72c668fa4df142179d1833ae1a4a81a7.zip)


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

## 简单实例

详细使用方法请参考[Nephele API]()。

在~/nephele/image目录下添加一张名为1.jpg的图片。

**查看原图 :8080/1.jpg**

[原图](https://dimg08.c-ctrip.com/images/w20e0s000000hvulz1A96.jpg)

**图片被等比缩放至目标区域区域内 :8080/1.jpg?x-nephele-process=image/resize,w_200,h_200**

[保持长宽比不变，缩放至200X200的区域内的最大图片](https://dimg08.c-ctrip.com/images/w20h0s000000hufjr466E.jpg)

**图片底部被裁剪，最终高为200 :8080/1.jpg?x-nephele-process=image/crop,m_b,h_200**

[裁剪底部直至高为200](https://dimg08.c-ctrip.com/images/w2040s000000hrubd8B2F.jpg)

**图片顺时针旋转90度 :8080/1.jpg?x-nephele-process=image/rotate,v_90**

[顺时针旋转90度](https://dimg08.c-ctrip.com/images/w20q0s000000hued11162.jpg)

**降低图片质量 :8080/1.jpg?x-nephele-process=image/quality,v_5**

[图片质量降低至5](https://dimg08.c-ctrip.com/images/w20c0s000000hvr5p3621.jpg)

**图片格式转换为png :8080/1.jpg?x-nephele-process=image/format,v_png**

[图片格式转为png](https://dimg08.c-ctrip.com/images/w20m0s000000i2hswD841.png)

**锐化图片 :8080/1.jpg?x-nephele-process=image/sharpen,r_44,s_3**

[特定参数锐化](https://dimg08.c-ctrip.com/images/w2030s000000huu7l6749.jpg)

**在~/nephele/image目录下添加用作水印的图片wm1.png**

[用作水印的图片](https://dimg08.c-ctrip.com/images/w20u0s000000hteatB945.png)

**图片打上水印（水印图片名字经base64编码处理） :8080/1.jpg?x-nephele-process=image/watermark,n_d20xLnBuZw==**

[图片被打上水印](https://dimg08.c-ctrip.com/images/w20d0s000000hrxuf69D9.jpg)
