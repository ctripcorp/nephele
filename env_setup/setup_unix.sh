#! /bin/bash

echo `uname -a`
os=`uname -a`

centos(){
echo "centos"
sudo yum install -y GraphicsMagick-devel
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

[[ $os =~ "el" ]] &&centos
[[ $os =~ "Ubuntu" ]] &&ubuntu_debian
[[ $os =~ "Debian" ]] &&ubuntu_debian
[[ $os =~ "Mac" ]] &&macOS









