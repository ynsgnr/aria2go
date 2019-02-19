#ifndef ARIA2_BRIDGE
#define ARIA2_BRIDGE

    #ifdef __cplusplus
    extern "C" {
    #endif
        void* new_aria2go(void);
        void del_aria2go(void *);
        void init_aria2go(void *);
    #ifdef __cplusplus
    }
    #endif

#endif
