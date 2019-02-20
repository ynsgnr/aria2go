
//#include <aria2/aria2.h> - if shared library is present on system
#include "aria2.h"

#include "aria2Interface.hpp"

int Aria2Interface::init_libaria2(){
    aria2::libraryInit();
    return 0;
}

void Aria2Interface::set_session(void * s){
    Aria2Interface::session=(aria2::Session*)s;
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
        throw "Unable to create sesion";
    }
    return (void *)s;
}

int Aria2Interface::run_libaria2(){
    //TODO add enum
    return aria2::run(Aria2Interface::session,aria2::RUN_ONCE);
}

Aria2Interface::~Aria2Interface(){
    aria2::libraryDeinit();
}

//TODO check if any ongoing download exists and end session if not
//TODO before starting download check if session is active and if not create and run one