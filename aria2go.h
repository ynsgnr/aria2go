#ifndef ARIA2_BRIDGE
#define ARIA2_BRIDGE

    #ifdef __cplusplus
    extern "C" {
    #endif
        void* new_aria2go(void);
        void del_aria2go(void*);
        void init_aria2go(void*);
        void* init_aria2go_session (void*);
        int run_aria2go(void*,void*);
        char* gidToHex_aria2go(void*,void*);
        void* hexToGid_aria2go(void*,char*);
        int isNull_aria2go(void*,void*);
        void* addUri_aria2go(void*,char*,int);
        int addMetalink_aria2go(void*,char*,int);
        void* get_element_gid(int);
        int get_element_int_value(int);
        int arraytest(void*);
        void add_uri(void*,char*);
        void clear_uris();
        void* add_all_from_cache(void*,int);
        void* addTorrent_aria2go(void*,char*,int);
        int getActiveDownload_aria2go(void*);
    #ifdef __cplusplus
    }
    #endif

#endif
