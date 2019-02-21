#include "aria2.h"

class Aria2Interface {
    public:
        ~Aria2Interface();
        int init_libaria2();
        void* init_libaria2_session();
        int run_libaria2();
        void set_session(void*);
        const char* gidToHex_libaria2(void*);
        void * hexToGid_libaria2(char* s);
        bool isNull_libaria2(void* g);
        void* addUri_libaria2(char*,int);
        void* addMetalink_libaria2(char*,int,int*,int*);
        void* arraytest(int*,int*);
    private:
        aria2::Session* session = NULL;
        void clear_session();
};