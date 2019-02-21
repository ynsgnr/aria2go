
//#include <aria2/aria2.h> - if shared library is present on system
#include "aria2.h"

#include "aria2Interface.hpp"

int Aria2Interface::init_libaria2(){
    aria2::libraryInit();
    return 0;
}

void Aria2Interface::set_session(void * s){
    clear_session();
    session=(aria2::Session*)s;
}

int downloadEventCallback(aria2::Session* s, aria2::DownloadEvent e,
                          aria2::A2Gid gid, void* userData){
    return 0;
}
        

void* Aria2Interface::init_libaria2_session(){
    //TODO implement session options
    aria2::SessionConfig config;
    config.downloadEventCallback = downloadEventCallback;
    aria2::Session* s = aria2::sessionNew(aria2::KeyVals(),config);
    if(s==NULL){
        throw "Unable to create session";
    }
    return (void *)s;
}

int Aria2Interface::run_libaria2(){
    //TODO add enum
    return aria2::run(Aria2Interface::session,aria2::RUN_ONCE);
}

const char* Aria2Interface::gidToHex_libaria2(void* g){
    aria2::A2Gid gid = (aria2::A2Gid) g;
    return aria2::gidToHex(gid).c_str();
}

void Aria2Interface::clear_session(){
    if(session!=NULL){
        aria2::sessionFinal(session);
    }
}

void * Aria2Interface::hexToGid_libaria2(char* s){
    if(s==NULL){
        throw "Undefined String for Hex To Gid transform";
    }
    return (void *) aria2::hexToGid(std::string (s));
}

bool Aria2Interface::isNull_libaria2(void* g){
    return aria2::isNull( (aria2::A2Gid) g);
}

Aria2Interface::~Aria2Interface(){
    clear_session();
    aria2::libraryDeinit();
}

//TODO check if any ongoing download exists and end session if not
//TODO before starting download check if session is active and if not create and run one