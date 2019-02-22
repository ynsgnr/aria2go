
#define TO_OBJECT(a) Aria2Interface * object = (Aria2Interface *)a;

#include "aria2Interface.hpp"
#include "aria2go.h"
#include "aria2.h"
#include <string.h>

// C wrapper for go

//Functions must take in Downloader
//Each Downloader object must be cast to DownloaderLib pointer
//Pointers must be DownloaderLib

void* new_aria2go(void){
    auto r = new Aria2Interface;
    return r;
}

void del_aria2go(void* a){
    TO_OBJECT(a)
    if(current_array!=NULL) delete current_array;
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

char* gidToHex_aria2go(void* a,void* gid){
    TO_OBJECT(a)
    return object->gidToHex_libaria2(gid);
}

void* hexToGid_aria2go(void* a,char * s){
    TO_OBJECT(a)
    return object->hexToGid_libaria2(s);
}

int isNull_aria2go(void* a, void* g){
    TO_OBJECT(a)
    return object->isNull_libaria2(g);
}

void* addUri_aria2go(void* a, char* uri, int position=-1){
    TO_OBJECT(a)
    return object->addUri_libaria2(uri,position);
}

int addMetalink_aria2go(void* a,char* file_location,int position=-1){
    TO_OBJECT(a)
    int* l = new int();
    int* s = new int();
    if(current_array!=NULL) delete current_array; //TODO fix here
    current_array = object->addMetalink_libaria2(file_location,position,l,s);
    current_index_size = *s/sizeof(void*);
    current_array_length = *l;
    delete l;
    delete s;
    return current_array_length;
}

void* get_element_gid(int index){
    if(index>=current_array_length) throw "Out Of Index";
    aria2::A2Gid* array = (aria2::A2Gid*)current_array;
    return array+index;
}

int get_element_int_value(int index){
    if(index>=current_array_length) throw "Out Of Index";
    int* array = (int*)current_array;
    return *(array+index);
}

int arraytest(void* a){
    TO_OBJECT(a)
    int* l = new int();
    int* s = new int();
    current_array = object->arraytest(l,s);
    current_index_size = (*s)/sizeof(void*);
    current_array_length = *l;
    delete l;
    delete s;
    return current_array_length;
}