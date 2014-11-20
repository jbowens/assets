#!/bin/bash

sudo apt-get install git
sudo apt-get install build-essential
cd /usr/local/lib/
git clone git@github.com:sass/sassc.git
git clone git@github.com:sass/libsass.git
# Initialize and update the submodule sass2scss…
cd libsass/
git submodule update –-init
export SASS_LIBSASS_PATH="/usr/local/lib/libsass"
# Make sure it worked…
echo $SASS_LIBSASS_PATH
# Now you can make SassC…
cd /usr/local/lib/sassc/
make
cd /usr/local/bin/
ln -s ../lib/sassc/bin/sassc sassc
