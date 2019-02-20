#include "aria2.h"

class Aria2Interface {
    public:
        ~Aria2Interface();
        int init_libaria2();
        void* init_libaria2_session();
        int run_libaria2();
        void set_session(void*);
    private:
        aria2::Session* session;
};