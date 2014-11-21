#!/bin/bash

npm install -g typescript
npm install -g uglify-js

SASS_LOC=`pwd`
echo $SASS_LOC
cd ..
git clone https://github.com/hcatlin/sassc.git
git clone https://github.com/hcatlin/libsass.git
git clone https://github.com/suapapa/go_sass.git
cd libsass/
git submodule update
export SASS_LIBSASS_PATH="`pwd`"
cd ../sassc
make
cd ..
cd go_sass/
sudo ./install_libsass.sh
cd $SASS_LOC

go get github.com/suapapa/go_sass
go get github.com/stretchr/testify
