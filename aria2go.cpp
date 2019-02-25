
#define TO_GID(gid_to_convert) aria2::A2Gid gid = (aria2::A2Gid) gid_to_convert;
#define TO_HANDLE_POINTER(handle_to_convert) aria2::DownloadHandle* handle = (aria2::DownloadHandle*) handle_to_convert;
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

int downloadEventCallback(aria2::Session* s, aria2::DownloadEvent e,
                          aria2::A2Gid gid, void* userData){
    if(session == NULL || s != session) return 0;
    DownloadEvent event;
    switch (e) {
        case aria2::EVENT_ON_DOWNLOAD_START:
            event = EVENT_ON_DOWNLOAD_START;
            break;
        case aria2::EVENT_ON_DOWNLOAD_PAUSE:
            event = EVENT_ON_DOWNLOAD_PAUSE;
            break;
        case aria2::EVENT_ON_DOWNLOAD_STOP:
            event = EVENT_ON_DOWNLOAD_STOP;
            break;
        case aria2::EVENT_ON_DOWNLOAD_COMPLETE:
            event = EVENT_ON_DOWNLOAD_COMPLETE;
            break;
        case aria2::EVENT_ON_BT_DOWNLOAD_COMPLETE:
            event = EVENT_ON_BT_DOWNLOAD_COMPLETE;
            break;
        default:
            event = EVENT_ON_DOWNLOAD_ERROR;
    }
    runGoCallBack(event,(void*)gid);
    return 0;
}

void init_aria2go(){
    aria2::libraryInit();
}

void* init_aria2go_session (int keep_running){
    aria2::SessionConfig config;
    config.downloadEventCallback = downloadEventCallback;
    config.keepRunning = keep_running;
    aria2::Session* s = aria2::sessionNew(aria2::KeyVals(),config);
    if(s==NULL){ return nullptr; }
    return (void *)s;
}

int run_aria2go(void* s, int run_mode){
    if(s == nullptr) return -1;
    session = (aria2::Session*)s;
    if(run_mode){
        return aria2::run(session,aria2::RUN_ONCE);
    }else{
        return aria2::run(session,aria2::RUN_DEFAULT);
    }
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

int finalize_aria2go(){
    if(session){
        aria2::shutdown(session);
        int r = aria2::sessionFinal(session);
        return r;
    }
    return -1;
}

int deinit_aria2go(){
    return aria2::libraryDeinit();
}

void* getDownloadHandle_aria2go(void* g){
    TO_GID(g)
    aria2::DownloadHandle* h = aria2::getDownloadHandle(session,gid);
    return (void*) h;
}

enum DownloadStatus getStatus_downloadHandle(void* h){
    TO_HANDLE_POINTER(h)
    aria2::DownloadStatus s = handle->getStatus();
    switch (s) {
        case aria2::DOWNLOAD_ACTIVE:
            return DOWNLOAD_ACTIVE;
        case aria2::DOWNLOAD_WAITING:
            return DOWNLOAD_WAITING;
        case aria2::DOWNLOAD_PAUSED:
            return DOWNLOAD_PAUSED;
        case aria2::DOWNLOAD_COMPLETE:
            return DOWNLOAD_COMPLETE;
        case aria2::DOWNLOAD_REMOVED:
            return DOWNLOAD_REMOVED;
        default:
            return DOWNLOAD_ERROR;
    }
}