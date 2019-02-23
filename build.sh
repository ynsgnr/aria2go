#!/bin/bash

unameOut="$(uname -s)"

case "${unameOut}" in

    Linux*)     lib_file=libaria2go.so;;

    Darwin*)    lib_file=libaria2go.dylib;;

    CYGWIN*)    lib_file=libaria2go.dll;;

    MINGW*)     lib_file=libaria2go.dll;;

    *)          lib_file=libaria2go.so;;

esac
echo "Compiling ${lib_file}"
g++ -o "${lib_file}" aria2go.cpp -O3 -Wall -Wextra -fPIC -shared -l aria2 -DBUILD_DLL

go test -v

