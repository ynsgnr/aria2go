#include "aria2.h"

class Aria2Interface {
    public:
        ~Aria2Interface();
        int init_libaria2();
        void* init_libaria2_session();
        int run_libaria2();
        void set_session(void*);
        char* gidToHex_libaria2(void*);
        void * hexToGid_libaria2(char* s);
        bool isNull_libaria2(void* g);
        void* addUri_libaria2(char*,int);
        aria2::A2Gid* addMetalink_libaria2(char*,int,int*);
        void* arraytest(int*);
        void* addTorrent_libaria2(char*,int);
        void add_uri(char*);
        void clear_uris();
        void* add_all_from_cache(int);
        aria2::A2Gid* getActiveDownload_libaria2(int*);
    private:
        aria2::Session* session = NULL;
        void clear_session();
        std::vector<std::string> uris;
};