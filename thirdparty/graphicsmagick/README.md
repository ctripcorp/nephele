## Nephele GraphicsMagick安装指南


### 安装
* unix系统（centos,ubuntu,debian），执行setup_unix.sh
* windows系统（win7及以上），执行setup_windows.bat

### 检查安装
 在命令窗口输入gm,看到类似如下结果代表安装成功,
```
GraphicsMagick 1.3.21 2015-02-28 Q8 http://www.GraphicsMagick.org/
Copyright (C) 2002-2014 GraphicsMagick Group.
Additional copyrights and licenses apply to this software.
See http://www.GraphicsMagick.org/www/Copyright.html for details.
Usage: gm command [options ...]

Where commands include: 
    animate - animate a sequence of images
      batch - issue multiple commands in interactive or batch mode
  benchmark - benchmark one of the other commands
    compare - compare two images
  composite - composite images together
    conjure - execute a Magick Scripting Language (MSL) XML script
    convert - convert an image or sequence of images
    display - display an image on a workstation running X
       help - obtain usage message for named command
   identify - describe an image or image sequence
     import - capture an application or X server screen
    mogrify - transform an image or sequence of images
    montage - create a composite image (in a grid) from separate images
       time - time one of the other commands
    version - obtain release version
```

 ### 安装失败怎么办？
 如果未能通过脚本成功安装，可以通过源码编译安装http://www.graphicsmagick.org/README.html

 ### 注意：
     脚本执行过程中如果被终止，那么临时文件epel-aliyun.repo可能会留在/etc/yum.repos.d/目录。
