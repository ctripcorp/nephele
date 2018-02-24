# Nephele

## Background


随着图片量的快速增长，需要专门的图片系统来管理...
随着业务的不断发展，图片预先处理已不能满足业务需求，需要访问时实时动态缩放裁剪等处理...
随着对图片质量要求越来越高，各种滤镜等高级需求增长，添加了图片编辑器，供用户自行编辑。

Nephele因此而诞生。
=======

Nephele是携程框架部门研发的开源图片服务，使用Go开发，意在为业务和开发提供一套完整的在线图片解决方案，其架构简单，可以稳定运行在绝大多数主流环境下，不仅提供功能完善的图片上传，访问，处理api，可视化监控和在线图片编辑功能也能够方便非开发人员协同管理。

## Introduction

Nephele是携程框架部门研发的开源图片服务，使用Go开发，意在为业务和开发提供一套完整的在线图片解决方案，其架构简单，可以稳定运行在绝大多数主流环境下，不仅提供功能完善的图片上传，访问，处理api，可视化监控和在线图片编辑功能也能够方便非开发人员协同管理。

## Functionality

  * **多存储支持**
  	* 本地文件，FastDFS，OSS，S3
  * **图片编码及优化**
  	* JPEG，PNG，GIF，WEBP。
  	* 动态最优质量评估。
  * **多种图片处理**
  	* 裁剪，缩放，旋转，转换，滤镜，水印，原图保护，样式。
  * **图片监控**
  	* 资源消耗，访问量，延时，失败率，回源率。
  * **图片管理面板**
  	* 图片管理，图片编辑。  
  * **快速部署** 		
  	  	
  	
  	
  	
  	