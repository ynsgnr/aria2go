# What Is This?
This go module adds ability to use [aria2](https://github.com/aria2/aria2). Aria2 is a powerful downloader tool that can use split download streams to speed up process and can download variaty of links and files including bittorent and metafile.

# How to use
To use this library, it needs to be initilized with `New()` function. Recomended way to use aria2go is to use `(aria2go) keepRunning()` function to keep running session within a go routine. When aria2go no longer needed or program finishes `(aria2go) finalize()` function must be called to prevent memory leaks. `(aria2go) keepRunning()` must be called before performing any other process.

If you want to custumize aria2go run methods you can use  `(aria2go) init_aria2go_session(keepRunning bool)` and other run functions to initilize and run a session. If you are not using the recomended way you must initilize a session first then use run functions as your will. If you run a function without initilizing a session first It will return immediately with an empty or nil value. For mor information you can check [libaria2 functions documentation](https://aria2.github.io/manual/en/html/libaria2.html#functions).

# Structs
 - ## aria2go
    This is an empty struct with only functions. Main functions included here

 - ## Gid
    This is the object represents a download process. When a new uri or file added this struct returns to manage download. It includes several functions and a pointer that can only be used in C.

 - ## FileData
    This struct includes information about downloaded files. You can get this struct using `getFiles()` function included in `Gid` structure. `getFiles()` function returns an array of `FileData` which includes all the files that download process has. Struct also includes these data:

    `index`:  1-based index of the file in the download. This is the same order with the files in multi-file torrent.

    `path`:  The local file path to this file when downloaded.

    `length` : The file size in bytes. This is not the current size of the local file.

    `completedLength`:  The completed length of this file in bytes. Please note that it is possible that sum of `completedLength` for each file in a download process is less than the return value of `getCompletedLength()` function of the download process. This is because the `completedLength` only calculates completed pieces. On the other hand, `getCompletedLength()` takes into account of partially completed piece.

 - ## GlobalStat
    This struct can be obtained by using `getGlobalStat()` function on `aria2go` object. And It has information about blobal state of all download processes.

	`downloadSpeed` : Overall download speed (byte/sec).

	`uploadSpeed`   : Overall upload speed(byte/sec).

	`numActive`     : The number of active downloads.

	`numWaiting`    : The number of waiting downloads.

	`numStopped`    : The number of stopped downloads.

 - ## BtMetaInfoData
    This struct can be  obtained by using `getBtMetaInfo()` funciton on a `Gid` object. It is only valid if the `Gid` belongs to is created with a torrent file and has information about torrent file. Othervise `valid` variable will be false

    `comment`   : Comment for the torrent. It is string value.

    `creationData`  : The creation time of the torrent. It's time object

    `multiMod`  : Boolean value showing torrent file has multiple files or not. If both `singleMod` and `multiMod` are false file mode is not available.

    `singleMod` : Boolean value showing torrent file has single file or not. If both `singleMod` and `multiMod` are false file mode is not available.

    `name`   : Name of the torrent. It is string value.

    `valid` : Boolean value showing `BtMetaInfoData` is valid or not. It is  `true`only when download process added with a torrent file

# Functions

- `(aria2go) init_aria2go_session(keepRunning bool)`: Takes `keepRunning` value to determine if  `run()` function returns when all download processes finished or not. For memory reasons, there can only be one session initilized at one time. So this function removes previous sessions if any initilized. Download processes of previous session is canceled.

- `(aria2go) run()`: Runs aria2go session. This function returns when no downloads are left to be processed.

- `(aria2go) runOnce()`: Runs aria2go session once. This function returns after one event polling. In the current implementation, event polling timeouts in 1 second. This function also returns on each timeout. On return, when no downloads are left to be processed, this function returns 0. Otherwise, returns 1, indicating that the caller must call this function one or more time to complete downloads.

- `(aria2go) keepRunning()`: This is the recomended function to use when running aria2go. It initilizes its own session, and runs in a routine.
This function return immediately, however session keeps running inside a routine. This function will initilize a new session, and removes previous session if any initilized. This function needs to be called before adding any uri, or any process

- `(aria2go) runUntillFinished()`: This function returns when all downloads are finished if session initilized with `keepRunning` as `false`. Otherwise it will not going to return. It is not advised to use this function in a go routine, instead use `keepRunning()` function which runs in a go routine in its own.

- `(aria2go) gidToHex(Gid) string`: Returns textual representation of the gid.

- `(aria2go) hexToGid(string) Gid`: Returns GID converted from the textual representation hex.

- `(aria2go) isNull(Gid) bool`: Returns true if the gid is invalid.

- `(aria2go) addUriInPosition(uri string, position int) Gid`: Adds new HTTP(S)/FTP/BitTorrent Magnet URI. The uris includes URI to be downloaded. For BitTorrent Magnet URI, the uris must have only one element and it should be BitTorrent Magnet URI. URIs in the uris must point to the same file. If you mix other URIs which point to another file, aria2 does not complain but download may fail. If the position is not negative integer, the new download is inserted at position in the waiting queue. If the position is negative or the position is larger than the size of the queue, it is appended at the end of the queue. 

- `(aria2go) addUri(uri string) Gid`:  Adds new HTTP(S)/FTP/BitTorrent Magnet URI. The uris includes URI to be downloaded.

- `(aria2go) addMetalinkInPosition(file_location string, position int) []Gid`:  Adds Metalink file download from file_location path. If the position is not negative integer, the new download is inserted at position in the waiting queue. If the position is negative or the position is larger than the size of the queue, it is appended at the end of the queue. This function returns 0 if it succeeds, or negative error code.

- `(aria2go) addMetaLink(file_location string) []Gid`:  Adds Metalink file download from file_location path.

- `(aria2go) addUriToCache(uri string)`: Adds uri to cache which kept in C side. This function exist to pass multiple uris to `addUri(uri string)` functions. When these functions run, uri cahe is cleared and given uri added with cached uris. Uris must point to the same file otherwise they are ignored and uri cache is cleared eitherway.

- `(aria2go) addAllFromCacheWithPosition(position int) Gid`:  Adds every uri from uri cache to be downloaded in a position. Uris can be added with `(aria2go) addUriToCache(uri string)` function. Each uri must point to the same file.

- `(aria2go) addAllFromCache() Gid`:  Adds every uri from uri cache to be downloaded. Uris can be added with `(aria2go) addUriToCache(uri string)` function. Each uri must point to the same file.

- `(aria2go) getActiveDownload() []Gid`:  Returns the array of active download GID.

- `(aria2go) removeDownload(gid Gid)`:  Removes the download denoted by the gid. If the specified download is in progress, it is stopped at first. The status of removed download becomes `DOWNLOAD_REMOVED`.

- `(aria2go) forceRemoveDownload(gid Gid)`: Force removes the download denoted by the gid. Removal will take place without any action which takes time such as contacting BitTorrent tracker.

- `(aria2go) pauseDownload(gid Gid)`:   Pauses the download denoted by the gid. The status of paused download becomes DOWNLOAD_PAUSED. If the download is active, the download is placed on the first position of waiting queue. As long as the status is `DOWNLOAD_PAUSED`, the download will not start. To change status to `DOWNLOAD_WAITING`, use `unpauseDownload()` function. Please note that, to make pause work, the application must set SessionConfig::keepRunning to true. Otherwise, the behavior is undefined.

- `(aria2go) forcePauseDownload(gid Gid)`: Pauses the download denoted by the gid.  Pause will take place without any action which takes time such as contacting BitTorrent tracker. 

- `(aria2go) unpauseDownload(gid Gid)`: Changes the status of the download denoted by the gid from `DOWNLOAD_PAUSED` to `DOWNLOAD_WAITING`. This makes the download eligible to restart.

- `(aria2go) setEventCallback(eventCallback EventCallback)`: Sets the event callback to be triggered whenever a download processes event triggered. Event callback is function that takes DownloadEvent and Gid as input and returns nothing. DownloadEvent is an ENUM representing event happened. And Event callback function must be in this format : `func(DownloadEvent, Gid)`. Check enums section for download events.

- `(aria2go) finalize()` : Finilizes initilized session and deinits library by freeing up taken memory. Must be run before go program comes to an end.

- `(g Gid) getStatus() DownloadStatus` : Returns status of this download as DownloadStatus enum. Check enums section for more info.

- `(g Gid) getTotalLength() int64` : Returns the total length of this download in bytes.

- `(g Gid) getCompletedLength() int64` : Returns the completed length of this download in bytes.

- `(g Gid) getUploadLength() int64` : Returns the uploaded length of this download in bytes.

- `(g Gid) getBitfield() string` : Returns the download progress in byte-string. The highest bit corresponds to piece index 0. The set bits indicate the piece is available and unset bits indicate the piece is missing. The spare bits at the end are set to zero. When download has not started yet, returns empty string.

- `(g Gid) getDownloadSpeed() int` : Returns download speed of this download measured in bytes/sec.

- `(g Gid) getUploadSpeed() int` : Returns upload speed of this download measured in bytes/sec.

- `(g Gid) getInfoHash() string` : Returns 20 bytes InfoHash if BitTorrent transfer is involved. Otherwise the empty string is returned.

- `(g Gid) getNumPieces() int` : Returns the number of pieces.

- `(g Gid) getConnections() int` : Returns the number of peers/servers the client has connected to.

- `(g Gid) getErrorCode() int` : Returns the last error code occurred in this download. The error codes are defined in [EXIT STATUS section of aria2c](https://aria2.github.io/manual/en/html/aria2c.html#exit-status). This value has its meaning only for stopped/completed downloads.

- `(g Gid) getDir() string` : Returns the array of files this download contains.

- `(g Gid) getNumFiles() int` : Returns the number of files. The return value is equivalent to

- `(g Gid) getFiles() []FileData` : Returns the array of files this download contains.

- `(g Gid) getBtMetaInfo() BtMetaInfoData` : Returns the information retrieved from “.torrent” file. This function is only meaningful only when BitTorrent transfer is involved in the download and the download is not stopped/completed.

- `(d aria2go) getGlobalStat() GlobalStat` : Returns global statistics such as overall download and upload speed.

# Enums
 DownloadEvent : Used in callback function. Can be one of:
  - EVENT_ON_DOWNLOAD_START
  - EVENT_ON_DOWNLOAD_PAUSE
  - EVENT_ON_DOWNLOAD_STOP
  - EVENT_ON_DOWNLOAD_COMPLETE
  - EVENT_ON_DOWNLOAD_ERROR
  - EVENT_ON_BT_DOWNLOAD_COMPLETE (BT used as BitTorrent.)

 DownloadStatus : Used in download process' status. Can be one of:
  - DOWNLOAD_ACTIVE
  - DOWNLOAD_WAITING
  - DOWNLOAD_PAUSED
  - DOWNLOAD_COMPLETE
  - DOWNLOAD_ERROR
  - DOWNLOAD_REMOVED


# Why
As a side project and I wanted to learn cross plotform development, compiling C, C++ in different enviroments and how comfigure and make works, while doing that I also wanted to create something useful. I have learned a lot about memory, C, go and compiling while writing this module and tring to compile included libraries for three different platforms.

# Libraries and Dependencies
Libraries complied for Mac, Linux and Windows by me. You can compile libraries and exchange files at your will. You can check aria2 C++ library documents for dependencies.

# Building
Thanks to recent simplifications a build process is not necessary. You can use it as any go module and golang takes care of the rest. However If you want to feel confortable you can build dependencies in your system and install them. Each dependency has a different process. You can check aria2 C++ library documents for dependencies.

