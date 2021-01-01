#!/bin/bash
pushd eb
./configure --disable-shared --disable-ebnet --disable-nls
pushd eb
make
popd
popd
