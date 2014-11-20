#!/bin/bash

sudo apt-get install git
sudo apt-get install build-essential
git clone git@github.com:sass/sassc.git
git clone git@github.com:sass/libsass.git
# Initialize and update the submodule sass2scss…
cd libsass/
git submodule update –-init
export SASS_LIBSASS_PATH="`pwd`/libsass"
# Make sure it worked…
echo $SASS_LIBSASS_PATH
