
#define TO_OBJECT(a) Aria2Interface * object = (Aria2Interface *)a;

#include "aria2Interface.hpp"
#include "aria2go.h"
#include "aria2.h"

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
    char* h = object->gidToHex_libaria2(gid);
    return h;
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

void* current_array;
int current_array_length;
aria2::A2Gid* current_gid_array;
int current_gid_array_length;

int addMetalink_aria2go(void* a,char* file_location,int position=-1){
    TO_OBJECT(a)
    int* l = new int();
    if(current_gid_array!=NULL) delete current_gid_array; //TODO fix here
    current_gid_array = object->addMetalink_libaria2(file_location,position,l);
    current_gid_array_length = *l;
    delete l;
    return current_gid_array_length;
}

void* get_element_gid(int index){
    if(index>=current_gid_array_length) throw "Out Of Index";
    return (void*)current_gid_array[index];
}

int get_element_int_value(int index){
    if(index>=current_array_length) throw "Out Of Index";
    int* array = (int*)current_array;
    return *(array+index);
}

int arraytest(void* a){
    TO_OBJECT(a)
    int* l = new int();
    current_array = object->arraytest(l);
    current_array_length = *l;
    delete l;
    return current_array_length;
}

void add_uri(void* a,char* uri){
    TO_OBJECT(a)
    object->add_uri(uri);
}

void clear_uris(void* a){
    TO_OBJECT(a)
    object->clear_uris();
}

void* add_all_from_cache(void* a,int position=-1){
    TO_OBJECT(a)
    return object->add_all_from_cache(position);
}

void* addTorrent_aria2go(void* a,char* file,int position=-1){
    TO_OBJECT(a)
    return object->addTorrent_libaria2(file,position);
}

int getActiveDownload_aria2go(void* a){
    TO_OBJECT(a)
    int* l = new int();    
    if(current_gid_array!=NULL) delete current_gid_array;
    current_gid_array = object->getActiveDownload_libaria2(l);
    current_gid_array_length = *l;
    delete l;
    return current_array_length;
}
