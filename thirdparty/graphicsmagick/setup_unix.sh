#! /bin/bash

echo `uname -a`
os=`uname -a`

centos(){
sudo yum makecache
sudo yum install -y GraphicsMagick-devel --enablerepo=epel 
}

centos7(){
sudo wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo 
}

centos6(){
sudo wget -O /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-6.repo 
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

[[ $os =~ "el7" ]] &&centos7&&centos
[[ $os =~ "el6" ]] &&centos6&&centos
[[ $os =~ "Ubuntu" ]] &&ubuntu_debian
[[ $os =~ "Debian" ]] &&ubuntu_debian
[[ $os =~ "Mac" ]] &&macOS









