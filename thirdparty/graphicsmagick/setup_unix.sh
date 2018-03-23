#! /bin/bash

echo `uname -a`
os=`uname -a`

centos7(){
sudo wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo 
sudo yum makecache 
sudo yum install -y GraphicsMagick-devel --enablerepo=epel  
}


centos6(){
sudo wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-6.repo
sudo yum makecache
sudo yum install -y GraphicsMagick-devel --enablerepo=epel 
}

ubuntu_debian(){
sudo apt-get -y install gcc
sudo apt-get -y install pkg-config

sudo apt-get -y --purge remove graphicsmagick*
sudo apt-get -y install libgraphicsmagick-dev  #in fact libgraphicsmagick1-dev will be installed
sudo apt-get -y install graphicsmagick*
}

macOS(){
brew install -y /GraphicsMagick*/
}

[[ $os =~ "el7" ]] &&centos7
[[ $os =~ "el6" ]] &&centos6
[[ $os =~ "Ubuntu" ]] &&ubuntu_debian
[[ $os =~ "Debian" ]] &&ubuntu_debian
[[ $os =~ "Mac" ]] &&macOS









