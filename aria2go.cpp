
#include "aria2Interface.hpp"
#include "aria2go.h"

// C wrapper for go

//Functions must take in Downloader
//Each Downloader object must be cast to DownloaderLib pointer
//Pointers must be DownloaderLib

void * new_aria2go(void){
    auto r = new Aria2Interface;
    return r;
}

void del_aria2go(void* s){
    Aria2Interface * sacrifice = (Aria2Interface *)s;
    delete sacrifice;
}

void init_aria2go(void* a){
     Aria2Interface * object = (Aria2Interface *)a;
     object->init_libaria2();
}

void* init_aria2go_session (void* a){
    Aria2Interface * object = (Aria2Interface *)a;
    return object->init_libaria2_session();
}

int run_aria2go(void* a,void* s){
    Aria2Interface * object = (Aria2Interface *)a;
    object->set_session(s);
    return object->run_libaria2();
}