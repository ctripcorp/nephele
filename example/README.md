# The Example Nephele Program

[Class 1](https://github.com/ctripcorp/nephele/blob/master/docs/cn/REQUIREMENT_CLASS.md)

## 介绍

基于[Nephele Pluggable Framework]()(NPF)，Nephele团队针对不同场景的需求的强弱，开发了多款**现成**的图片服务产品。我们把它们称为[The Example Nephele Program](https://github.com/ctripcorp/nephele/blob/master/example/README.md)。

* 用户或开发者只需要通过简单的配置部署它们的预编译文件，无需进行二次开发，便可以享受到完整的图片处理功能。

* 同时，它们又是NPF的示例代码，是使用NPF开发最好的参考。

现在已有的两款分别是[Minimal](https://github.com/ctripcorp/nephele/blob/master/example/minimal/README.md)和[Pluggable](https://github.com/ctripcorp/nephele/blob/master/example/pluggable/README.md)。

## 选择

如果你需要的功能非常简单，比如为自己的博客搭建一个动态图片处理服务以弱化版本更新引入的排版问题。你可能会希望这个服务最好能和你的博客网站部署在同一台虚机上并占用最少的服务器资源。那么Mininal对你而言就很不错。

如果你关心更多东西，比如你在纠结利用数据库中已有的数据来左右图片处理的流程，或者你需要考虑图片服务和一些主流存储系统的兼容性，又或者你必须把图片服务的工作状态写入到一个现有的日志系统中。那么直接使用Pluggable肯定是比只搭建Minimal并在其周遭再封装一层要来得方便的。
