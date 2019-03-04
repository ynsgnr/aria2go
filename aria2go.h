#ifndef ARIA2_BRIDGE
#define ARIA2_BRIDGE

    //#include "aria2.h" //Included for enums

    #ifdef __cplusplus
    extern "C" {
    #endif
        void init_aria2go();
        void init_aria2go_session(int);
        int run_aria2go(int);
        void keepruning_aria2go();
        char* gidToHex_aria2go(void*);
        void* hexToGid_aria2go(char*);
        int isNull_aria2go(void*);
        void* addUri_aria2go(char*,int);
        int addMetalink_aria2go(char*,int);
        void* get_element_gid(int);
        void add_uri(char*);
        void clear_uris();
        void* add_all_from_cache(int);
        void* addTorrent_aria2go(char*,int);
        int getActiveDownload_aria2go();
        int removeDownload_aria2go(void*, int);
        int pauseDownload_aria2go(void*, int);
        int unpauseDownload_aria2go(void*);

         // callback functions:
        enum DownloadEvent {
            EVENT_ON_DOWNLOAD_START = 0,
            EVENT_ON_DOWNLOAD_PAUSE = 1,
            EVENT_ON_DOWNLOAD_STOP = 2,
            EVENT_ON_DOWNLOAD_COMPLETE = 3,
            EVENT_ON_DOWNLOAD_ERROR = 4,
            EVENT_ON_BT_DOWNLOAD_COMPLETE = 5
        };
        extern void runGoCallBack(enum DownloadEvent,void*);

        //Download handle functions
        enum DownloadStatus{
            DOWNLOAD_ACTIVE = 0,
            DOWNLOAD_WAITING = 1,
            DOWNLOAD_PAUSED = 2,
            DOWNLOAD_COMPLETE = 3,
            DOWNLOAD_ERROR = 4,
            DOWNLOAD_REMOVED = 5
        };
        void* getDownloadHandle_aria2go(void*);
        enum DownloadStatus getStatus_gid(void*);
    
        int getTotalLength_gid(void*);
        int getCompletedLength_gid(void*);
        int getUploadLength_gid(void*);
        char* getBitfield_gid(void*);
        int getDownloadSpeed_gid(void*);
        int getUploadSpeed_gid(void*);
        char* getInfoHash_gid(void*);
        int getPieceLength_gid(void*);
        int getNumPieces_gid(void*);
        int getConnections_gid(void*);
        int getErrorCode_gid(void*);
        int getNumFiles_gid(void*);
        int getFiles_gid(void*);
        void* get_element_fileData(int);

        int finalize_aria2go();
        int deinit_aria2go();
    #ifdef __cplusplus
    }
    #endif

#endif
