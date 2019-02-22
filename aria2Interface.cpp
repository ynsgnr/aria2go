
#define TO_GID(gid_to_convert) aria2::A2Gid gid = (aria2::A2Gid) gid_to_convert;

//#include <aria2/aria2.h> - if shared library is present on system
#include "aria2.h"
#include "aria2Interface.hpp"

#include <string.h>

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
    if(s==NULL){ throw "Unable to create session"; }
    return (void *)s;
}

int Aria2Interface::run_libaria2(){
    //TODO add enum
    return aria2::run(Aria2Interface::session,aria2::RUN_ONCE);
}

char* Aria2Interface::gidToHex_libaria2(void* g){
    TO_GID(g)
    std::string h = aria2::gidToHex(gid);
    char* hex = new char(h.length());
    strcpy(hex,h.c_str());
    return hex;
}

void Aria2Interface::clear_session(){
    if(session!=NULL){
        aria2::sessionFinal(session);
        delete session;
    }
}

void* Aria2Interface::hexToGid_libaria2(char* s){
    if(s==NULL){ throw "Undefined String for Hex To Gid transform"; }
    return (void *) aria2::hexToGid(std::string (s));
}

bool Aria2Interface::isNull_libaria2(void* g){
    return aria2::isNull( (aria2::A2Gid) g);
}

void* Aria2Interface::addUri_libaria2(char* uri,int position=-1){
    //TODO implement options
    if(uri==NULL){ throw "Undefined String for adding uri"; }
    uris.push_back(std::string (uri));
    aria2::A2Gid gid;
    int is_error = aria2::addUri(session,&gid,uris,aria2::KeyVals(),position);
    if (is_error || aria2::isNull(gid)){
        std::string error_message = "Failed to add download uri";
        error_message += std::to_string(is_error);
        throw error_message;
    }
    clear_uris();
    return (void *) gid;    
}

aria2::A2Gid* Aria2Interface::addMetalink_libaria2(char* file_location,int position,int* length){
    std::vector<aria2::A2Gid>* gids;
    int is_error = aria2::addMetalink(session,gids,std::string (file_location),aria2::KeyVals(),position);
    if(is_error || gids==NULL) throw "Unable to add metalink";
    *length = gids->size();
    return gids->data();
}

void* Aria2Interface::arraytest(int* l){
    std::vector<int>* array;
    std::vector<int> array_object {1,12,35,16,43,67};
    array = &array_object;
    *l = array->size();
    return (void*) array->data();
}

void* Aria2Interface::addTorrent_libaria2(char* file_location,int position=-1){
    aria2::A2Gid gid;
    int is_error;
    if(uris.size()>0){
        is_error = aria2::addTorrent(session,&gid,std::string (file_location),uris,aria2::KeyVals(),position);
    }else{
        is_error = aria2::addTorrent(session,&gid,std::string (file_location),aria2::KeyVals(),position);
    }
    if (is_error || aria2::isNull(gid)){
        std::string error_message = "Failed to add download uri";
        error_message += std::to_string(is_error);
        throw error_message;
    }
    clear_uris();
    return (void *) gid;
}

void Aria2Interface::add_uri(char* uri){
    uris.push_back(std::string (uri));
}

void Aria2Interface::clear_uris(){
    uris.clear();
}

void* Aria2Interface::add_all_from_cache(int position=-1){
    //TODO implement options
    aria2::A2Gid gid;
    int is_error = aria2::addUri(session,&gid,uris,aria2::KeyVals(),position);
    if (is_error || aria2::isNull(gid)){
        std::string error_message = "Failed to add download uri";
        error_message += std::to_string(is_error);
        throw error_message;
    }
    clear_uris();
    return (void *) gid;    
}

aria2::A2Gid* Aria2Interface::getActiveDownload_libaria2(int* l){
    std::vector<aria2::A2Gid> gids = aria2::getActiveDownload(session);
    *l = gids.size();
    return gids.data();
}

Aria2Interface::~Aria2Interface(){
    clear_session();
    aria2::libraryDeinit();
}

//TODO check if any ongoing download exists and end session if not
//TODO before starting download check if session is active and if not create and run one