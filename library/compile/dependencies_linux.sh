#!/bin/bash

mkdir dependencies_linux
cd dependencies_linux

test -z "$PREFIX" && PREFIX=/usr/local/$HOST
    curl -L -O https://gmplib.org/download/gmp/gmp-6.1.2.tar.lz && \
    curl -L -O https://github.com/libexpat/libexpat/releases/download/R_2_2_5/expat-2.2.5.tar.bz2 && \
    curl -L -O https://www.sqlite.org/2018/sqlite-autoconf-3230100.tar.gz && \
    curl -L -O http://zlib.net/zlib-1.2.11.tar.gz && \
    curl -L -O https://c-ares.haxx.se/download/c-ares-1.14.0.tar.gz && \
    curl -L -O http://libssh2.org/download/libssh2-1.8.0.tar.gz

    tar xf gmp-6.1.2.tar.lz && \
    cd gmp-6.1.2 && \
    ./configure \
        --disable-shared \
        --enable-static \
        --prefix=/usr/local/$HOST && \
    make install
    cd ..

    tar xf expat-2.2.5.tar.bz2 && \
    cd expat-2.2.5 && \
    ./configure \
        --disable-shared \
        --enable-static \
        --prefix=/usr/local/$HOST && \
    make install
    cd ..

    tar xf sqlite-autoconf-3230100.tar.gz && \
    cd sqlite-autoconf-3230100 && \
    ./configure \
        --disable-shared \
        --enable-static \
        --prefix=/usr/local/$HOST && \
    make install
    cd ..

    tar xf zlib-1.2.11.tar.gz && \
    cd zlib-1.2.11 && \
    ./configure \
        --prefix=/usr/local/$HOST \
        --libdir=/usr/local/$HOST/lib \
        --includedir=/usr/local/$HOST/include \
        --static && \
    make install
    cd ..

    tar xf c-ares-1.14.0.tar.gz && \
    cd c-ares-1.14.0 && \
    ./configure \
        --disable-shared \
        --enable-static \
        --without-random \
        --prefix=/usr/local/$HOST && \
    make install
    cd ..

    tar xf libssh2-1.8.0.tar.gz && \
    cd libssh2-1.8.0 && \
    ./configure \
        --disable-shared \
        --enable-static \
        --prefix=/usr/local/$HOST && \
    make install
    cd ..
