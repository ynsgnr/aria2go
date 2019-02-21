#ifndef ARIA2_BRIDGE
#define ARIA2_BRIDGE

    #ifdef __cplusplus
    extern "C" {
    #endif
        void* new_aria2go(void);
        void del_aria2go(void *);
        void init_aria2go(void *);
        void* init_aria2go_session (void*);
        int run_aria2go(void*,void*);
        const char * gidToHex_aria2go(void*,void*);
    #ifdef __cplusplus
    }
    #endif

#endif
