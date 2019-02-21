
#define TO_OBJECT(a) Aria2Interface * object = (Aria2Interface *)a;

#include "aria2Interface.hpp"
#include "aria2go.h"
#include <string.h>

// C wrapper for go

//Functions must take in Downloader
//Each Downloader object must be cast to DownloaderLib pointer
//Pointers must be DownloaderLib

void * new_aria2go(void){
    auto r = new Aria2Interface;
    return r;
}

void del_aria2go(void* a){
    TO_OBJECT(a)
    delete object;
}

void init_aria2go(void* a){
    TO_OBJECT(a)
    object->init_libaria2();
}

void* init_aria2go_session (void* a){
    TO_OBJECT(a)
    return object->init_libaria2_session();
}

int run_aria2go(void* a,void* s){
    TO_OBJECT(a)
    object->set_session(s);
    return object->run_libaria2();
}

const char * gidToHex_aria2go(void* a,void* gid){
    TO_OBJECT(a)
    return object->gidToHex_libaria2(gid);
}

void * hexToGid_aria2go(void* a,char * s){
    TO_OBJECT(a)
    return object->hexToGid_libaria2(s);
}

int isNull_aria2go(void* a, void* g){
    TO_OBJECT(a)
    return object->isNull_libaria2(g);
}