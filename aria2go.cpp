
#define TO_OBJECT(a) 
#define TO_GID(gid_to_convert) aria2::A2Gid gid = (aria2::A2Gid) gid_to_convert;

#include "aria2go.h"
#include "aria2.h"
#include <string.h>

// C wrapper for go

//Functions must take in Downloader
//Each Downloader object must be cast to DownloaderLib pointer
//Pointers must be DownloaderLib

aria2::A2Gid* current_gid_array;
int current_gid_array_length;
aria2::Session* session = NULL;
std::vector<std::string> uris;

void* new_aria2go(void){
    return NULL;
}

void del_aria2go(void* a){
}

int downloadEventCallback(aria2::Session* s, aria2::DownloadEvent e,
                          aria2::A2Gid gid, void* userData){
    return 0;
}

void init_aria2go(void* a){
    aria2::libraryInit();
}

void* init_aria2go_session (void* a){
    aria2::SessionConfig config;
    config.downloadEventCallback = downloadEventCallback;
    aria2::Session* s = aria2::sessionNew(aria2::KeyVals(),config);
    if(s==NULL){ throw "Unable to create session"; }
    return (void *)s;
}

int run_aria2go(void* a,void* s){
    session = (aria2::Session*)s;
    return aria2::run(session,aria2::RUN_ONCE);
}

char* gidToHex_aria2go(void* a,void* g){
    TO_GID(g)
    std::string h = aria2::gidToHex(gid);
    char* hex = new char(h.length());
    strcpy(hex,h.c_str());
    return hex;
}

void* hexToGid_aria2go(void* a,char * s){
    if(s==NULL){ throw "Undefined String for Hex To Gid transform"; }
    return (void *) aria2::hexToGid(std::string (s));
}

int isNull_aria2go(void* a, void* g){
    return aria2::isNull( (aria2::A2Gid) g);
}

void* addUri_aria2go(void* a, char* uri, int position=-1){
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

int addMetalink_aria2go(void* a,char* file_location,int position=-1){
    TO_OBJECT(a)
    int* l = new int();
    if(current_gid_array!=NULL) delete current_gid_array; //TODO fix here
    std::vector<aria2::A2Gid>* gids;
    int is_error = aria2::addMetalink(session,gids,std::string (file_location),aria2::KeyVals(),position);
    if(is_error || gids==NULL) throw "Unable to add metalink";
    *l = gids->size();
    current_gid_array = gids->data();
    current_gid_array_length = *l;
    delete l;
    return current_gid_array_length;
}

void* get_element_gid(int index){
    if(index>=current_gid_array_length) throw "Out Of Index";
    return (void*)current_gid_array[index];
}

void add_uri(void* a,char* uri){
    uris.push_back(std::string (uri));   
}

void clear_uris(){
    uris.clear();
}

void* add_all_from_cache(void* a,int position=-1){
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

void* addTorrent_aria2go(void* a,char* file,int position=-1){
    aria2::A2Gid gid;
    int is_error;
    if(uris.size()>0){
        is_error = aria2::addTorrent(session,&gid,std::string (file),uris,aria2::KeyVals(),position);
    }else{
        is_error = aria2::addTorrent(session,&gid,std::string (file),aria2::KeyVals(),position);
    }
    if (is_error || aria2::isNull(gid)){
        std::string error_message = "Failed to add download uri";
        error_message += std::to_string(is_error);
        throw error_message;
    }
    clear_uris();
    return (void *) gid;
}

int getActiveDownload_aria2go(void* a){
    int* l = new int();    
    if(current_gid_array!=NULL) delete current_gid_array;
    std::vector<aria2::A2Gid> gids = aria2::getActiveDownload(session);
    *l = gids.size();
    current_gid_array = gids.data();
    current_gid_array_length = *l;
    delete l;
    return current_gid_array_length;
}
