#include "aria2.h"

class Aria2Interface {
    public:
        ~Aria2Interface();
        int init_libaria2();
        void* init_libaria2_session();
        int run_libaria2();
        void set_session(void*);
        const char* gidToHex_libaria2(void*);
    private:
        aria2::Session* session = NULL;
        void clear_session();
};