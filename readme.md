# What
This go module adds ability to use [aria2](https://github.com/aria2/aria2). Aria2 is a powerful downloader tool that can use split download streams to speed up process and can download variaty of links including bittorent.

# Why
As a side project I wanted to learn cross plotform development, compiling C, C++ in different enviroments and how comfigure and make works. Also I wanted to create something usefull in the process. This is a side project for a cross platform flutter app called Pera, a cross platform downloader and torrent client.

# Libraries and Dependencies
Libraries complied for Mac, Linux and Windows by me. You can compile libraries and exchange files at your will. Current makefile is not tested fully yet and not guaranteed to work.

# Building
Independent of your operating system you can use
```
g++ -o libaria2go aria2Interface.cpp aria2go.cpp -O3 -Wall -Wextra -fPIC -shared -l aria2 -DBUILD_DLL
```
to build if you make any changes to code. This needs to run before go commands to for go to be able to link files. This will produce already build file of libaria2go, extension varias according to operating system: .dll fow Windows .so for Linux and .dylib for Mac. Makefile is not completely tested yet, so its not advised to use it

