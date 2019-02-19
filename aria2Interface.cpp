
//#include <aria2/aria2.h> - if shared library is present on system
#include "aria2.h"

#include "aria2Interface.hpp"

#include <iostream>

int Aria2Interface::init_libaria2(){
    aria2::libraryInit();
    return 0;
}