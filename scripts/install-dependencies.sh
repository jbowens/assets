#!/bin/bash

npm install -g typescript
npm install -g uglify-js

SASS_LOC = `pwd`
echo $SASS_LOC
git clone https://github.com/hcatlin/sassc.git
git clone https://github.com/hcatlin/libsass.git
cd libsass/
git submodule update –-init
echo 'SASS_LIBSASS_PATH="`pwd`"' >> /etc/environment
source /etc/environment
cd ../sassc
make
cd ..
export PATH='$PATH:`pwd`/sassc/bin'

