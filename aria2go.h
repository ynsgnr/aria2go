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
        const char* gidToHex_aria2go(void*,void*);
        void* hexToGid_aria2go(void*,char*);
        int isNull_aria2go(void*,void*);
        void* addUri_aria2go(void*,char*,int);
        int addMetalink_aria2go(void*,char*,int);
        void* get_element(int);
        int get_element_int_value(int);
        void* current_array;
        int current_index_size;
        int current_array_length;
        int arraytest(void*);
    #ifdef __cplusplus
    }
    #endif

#endif
