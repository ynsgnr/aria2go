
#define TO_GID(gid_to_convert) aria2::A2Gid gid = (aria2::A2Gid) gid_to_convert;
#define ERROR_MESSAGE(message,code) {std::string error_message = message; error_message += std::to_string(code); throw error_message;}

#include "aria2go.h"
#include "aria2.h"
#include <string.h>

#include <iostream>

// C wrapper for go

aria2::A2Gid* current_gid_array = NULL;
int current_gid_array_length;
aria2::Session* session = NULL;
std::vector<std::string> uris;

enum DownloadEvent {
    EVENT_ON_DOWNLOAD_START = aria2::EVENT_ON_DOWNLOAD_START,
    EVENT_ON_DOWNLOAD_PAUSE = aria2::EVENT_ON_DOWNLOAD_PAUSE,
    EVENT_ON_DOWNLOAD_STOP = aria2::EVENT_ON_DOWNLOAD_STOP,
    EVENT_ON_DOWNLOAD_COMPLETE = aria2::EVENT_ON_DOWNLOAD_COMPLETE,
    EVENT_ON_DOWNLOAD_ERROR = aria2::EVENT_ON_DOWNLOAD_ERROR,
    EVENT_ON_BT_DOWNLOAD_COMPLETE = aria2::EVENT_ON_BT_DOWNLOAD_COMPLETE
};


int downloadEventCallback(aria2::Session* s, aria2::DownloadEvent e,
                          aria2::A2Gid gid, void* userData){
    return 0;
}

void init_aria2go(){
    aria2::libraryInit();
}

void* init_aria2go_session (){
    aria2::SessionConfig config;
    config.downloadEventCallback = downloadEventCallback;
    aria2::Session* s = aria2::sessionNew(aria2::KeyVals(),config);
    if(s==NULL){ return nullptr; }
    return (void *)s;
}

int run_aria2go(void* s){
    if(s == nullptr) return -1;
    session = (aria2::Session*)s;
    return aria2::run(session,aria2::RUN_ONCE);
}

char* gidToHex_aria2go(void* g){
    TO_GID(g)
    std::string h = aria2::gidToHex(gid);
    char* hex = new char(h.length());
    strcpy(hex,h.c_str());
    return hex;
}

void* hexToGid_aria2go(char * s){
    if(s==nullptr){ return nullptr; /* Undefined String for Hex To Gid transform */ }
    return (void *) aria2::hexToGid(std::string (s));
}

int isNull_aria2go( void* g){
    return aria2::isNull( (aria2::A2Gid) g);
}

void* addUri_aria2go( char* uri, int position=-1){
    //TODO implement options
    if(uri==nullptr){ return nullptr; /*throw "Undefined String for adding uri";*/}
    uris.push_back(std::string (uri));
    aria2::A2Gid gid;
    int error_code = aria2::addUri(session,&gid,uris,aria2::KeyVals(),position);
    if (error_code || aria2::isNull(gid)) ERROR_MESSAGE("Failed to add download uri",error_code)
    clear_uris();
    return (void *) gid;    
}

int addMetalink_aria2go(char* file_location,int position=-1){
    if(current_gid_array!=NULL) delete current_gid_array;
    std::vector<aria2::A2Gid>* gids;
    int error_code = aria2::addMetalink(session,gids,std::string (file_location),aria2::KeyVals(),position);
    if(error_code || gids==NULL) ERROR_MESSAGE("Unable to add metalink",error_code)
    current_gid_array_length = gids->size();
    current_gid_array = gids->data();
    return current_gid_array_length;
}

void* get_element_gid(int index){
    if(index>=current_gid_array_length) throw "Out Of Index";
    return (void*)current_gid_array[index];
}

void add_uri(char* uri){
    uris.push_back(std::string (uri));   
}

void clear_uris(){
    uris.clear();
}

void* add_all_from_cache(int position=-1){
    //TODO implement options
    aria2::A2Gid gid;
    int error_code = aria2::addUri(session,&gid,uris,aria2::KeyVals(),position);
    if (error_code || aria2::isNull(gid)) ERROR_MESSAGE("Failed to add download uri",error_code)
    clear_uris();
    return (void *) gid;    
}

void* addTorrent_aria2go(char* file,int position=-1){
    aria2::A2Gid gid;
    int error_code;
    if(uris.size()>0){
        error_code = aria2::addTorrent(session,&gid,std::string (file),uris,aria2::KeyVals(),position);
    }else{
        error_code = aria2::addTorrent(session,&gid,std::string (file),aria2::KeyVals(),position);
    }
    if (error_code || aria2::isNull(gid)) ERROR_MESSAGE("Failed to add download uri",error_code)
    clear_uris();
    return (void *) gid;
}

int getActiveDownload_aria2go(){
    if(current_gid_array!=NULL) delete current_gid_array;
    std::vector<aria2::A2Gid> gids = aria2::getActiveDownload(session);
    current_gid_array_length = gids.size();
    current_gid_array = gids.data();
    return current_gid_array_length;
}

int removeDownload_aria2go(void* g, int force){
    if(g==nullptr) return -1;
    TO_GID(g)
    if(aria2::isNull(gid)) return -1;
    int error_code = aria2::removeDownload(session,gid,force);
    return error_code;
}

int pauseDownload_aria2go(void* g, int force){
    if(g==nullptr) return -1;
    TO_GID(g)
    if(aria2::isNull(gid)) return -1;
    int error_code = aria2::pauseDownload(session,gid,force);
    return error_code;
}

int unpauseDownload_aria2go(void* g){
    if(g==nullptr) return -1;
    TO_GID(g)
    if(aria2::isNull(gid)) return -1;
    int error_code = aria2::unpauseDownload(session,gid);
    return error_code;
}

void callCallback(){
    runCallBack();
}