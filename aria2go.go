package aria2

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L. -l aria2
// #include "aria2go.h"
// #include <stdlib.h>
import "C"
import "unsafe"

//ENUMS
type DownloadEvent int

const (
	EVENT_ON_DOWNLOAD_START       DownloadEvent = C.EVENT_ON_DOWNLOAD_START
	EVENT_ON_DOWNLOAD_PAUSE       DownloadEvent = C.EVENT_ON_DOWNLOAD_PAUSE
	EVENT_ON_DOWNLOAD_STOP        DownloadEvent = C.EVENT_ON_DOWNLOAD_STOP
	EVENT_ON_DOWNLOAD_COMPLETE    DownloadEvent = C.EVENT_ON_DOWNLOAD_COMPLETE
	EVENT_ON_DOWNLOAD_ERROR       DownloadEvent = C.EVENT_ON_DOWNLOAD_ERROR
	EVENT_ON_BT_DOWNLOAD_COMPLETE DownloadEvent = C.EVENT_ON_BT_DOWNLOAD_COMPLETE
)

//ENUMS
type DownloadStatus int

const (
	DOWNLOAD_ACTIVE   DownloadStatus = C.DOWNLOAD_ACTIVE
	DOWNLOAD_WAITING  DownloadStatus = C.DOWNLOAD_WAITING
	DOWNLOAD_PAUSED   DownloadStatus = C.DOWNLOAD_PAUSED
	DOWNLOAD_COMPLETE DownloadStatus = C.DOWNLOAD_COMPLETE
	DOWNLOAD_ERROR    DownloadStatus = C.DOWNLOAD_ERROR
	DOWNLOAD_REMOVED  DownloadStatus = C.DOWNLOAD_REMOVED
)

func (ds DownloadStatus) String() string {
	names := [...]string{
		"active",
		"waiting",
		"paused",
		"completed",
		"has an error",
		"removed"}
	return names[ds]
}

type aria2go struct {
}

type Gid struct {
	ptr unsafe.Pointer
}

type FileData struct {
	index           int
	path            string
	completedLength int
	selected        bool
	//uris []string
}

type GlobalStat struct {
	downloadSpeed int
	uploadSpeed   int
	numActive     int
	numWaiting    int
	numStopped    int
}

func New() aria2go {
	var ret aria2go
	C.init_aria2go()
	return ret
}

//TODO separate session function bc we can not add uris until initilized

func (d aria2go) init_aria2go_session(keepRunning bool) {
	C.finalize_aria2go()
	if keepRunning {
		C.init_aria2go_session(C.int(1))
	} else {
		C.init_aria2go_session(C.int(0))
	}
}

func (d aria2go) run() int {
	//C.finalize_aria2go()
	//d.session = C.init_aria2go_session(C.int(0))
	return int(C.run_aria2go(C.int(0)))
}

func (d aria2go) runUntillFinished() {
	//Returns when all downloads finisged
	r := 1
	for r == 1 {
		r = int(C.run_aria2go(C.int(1)))
	}
}

func (d aria2go) keepRunning() {
	C.finalize_aria2go()
	C.init_aria2go_session(C.int(1))
	go func() {
		C.keepruning_aria2go()
	}()
}

func (d aria2go) runOnce() int {
	//C.finalize_aria2go()
	//d.session = C.init_aria2go_session(C.int(0))
	return int(C.run_aria2go(C.int(1)))
}

func (d aria2go) gidToHex(gid Gid) string {
	p := C.gidToHex_aria2go(gid.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (d aria2go) hexToGid(s string) Gid {
	var ret Gid
	ret.ptr = C.hexToGid_aria2go(C.CString(s))
	return ret
}

func (d aria2go) isNull(g Gid) bool {
	if C.isNull_aria2go(g.ptr) == 0 {
		return false
	}
	return true
}

func (d aria2go) addUriInPosition(uri string, position int) Gid {
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri), C.int(position))
	return ret
}

func (d aria2go) addUri(uri string) Gid {
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri), C.int(-1))
	return ret
}

func (d aria2go) addMetalinkInPosition(file_location string, position int) []Gid {
	var gids []Gid
	var gid Gid
	l := int(C.addMetalink_aria2go(C.CString(file_location), C.int(position)))
	for i := 0; i < l; i++ {
		gid.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, gid)
	}
	return gids
}

func (d aria2go) addMetaLink(file_location string) []Gid {
	return d.addMetalinkInPosition(file_location, -1)
}

func (d aria2go) addUriToCache(uri string) {
	//You can use this one even if you didnt start session yet
	C.add_uri(C.CString(uri))
}

func (d aria2go) clearUriCache() {
	C.clear_uris()
}

func (d aria2go) addAllFromCacheWithPosition(p int) Gid {
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(p))
	return ret
}
func (d aria2go) addAllFromCache() Gid {
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(-1))
	return ret
}

func (d aria2go) getActiveDownload() []Gid {
	var gids []Gid
	l := int(C.getActiveDownload_aria2go())
	for i := 0; i < l; i++ {
		var g Gid
		g.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, g)
	}
	return gids
}

func (d aria2go) removeDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr, C.int(0))
	//Second variable is bool for forcing
}

func (d aria2go) forceRemoveDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr, C.int(1))
}

func (d aria2go) pauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr, C.int(0))
	//Second variable is bool for forcing
}

func (d aria2go) forcePauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr, C.int(1))
}

func (d aria2go) unpauseDownload(g Gid) {
	C.unpauseDownload_aria2go(g.ptr)
}

//Event Callback
type EventCallback func(DownloadEvent, Gid)

var callback EventCallback //Have to be global due to garbage collector

func (d aria2go) setEventCallback(eventCallback EventCallback) {
	callback = eventCallback
}

//export runGoCallBack
func runGoCallBack(event C.enum_DownloadEvent, g unsafe.Pointer) {
	var gid Gid
	gid.ptr = g
	callback(DownloadEvent(event), gid)
}

func (d aria2go) finalize() {
	C.finalize_aria2go()
	C.deinit_aria2go()
}

func (g Gid) getStatus() DownloadStatus {
	return DownloadStatus(C.getStatus_gid(g.ptr))
}

func (g Gid) getTotalLength() int {
	return int(C.getTotalLength_gid(g.ptr))
}

func (g Gid) getBitfield() string {
	p := C.getBitfield_gid(g.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (g Gid) getDownloadSpeed() int {
	return int(C.getDownloadSpeed_gid(g.ptr))
}

func (g Gid) getUploadSpeed() int {
	return int(C.getUploadSpeed_gid(g.ptr))
}

func (g Gid) getInfoHash() string {
	p := C.getInfoHash_gid(g.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (g Gid) getNumPieces() int {
	return int(C.getNumPieces_gid(g.ptr))
}

func (g Gid) getConnections() int {
	return int(C.getConnections_gid(g.ptr))
}

func (g Gid) getErrorCode() int {
	return int(C.getErrorCode_gid(g.ptr))
}

func (g Gid) getNumFiles() int {
	return int(C.getNumFiles_gid(g.ptr))
}

func (g Gid) getFiles() []FileData {
	var files []FileData
	var ptr unsafe.Pointer
	l := int(C.getFiles_gid(g.ptr))
	for i := 0; i < l; i++ {
		var f FileData
		ptr = C.get_element_fileData(C.int(i))
		f.index = int(C.get_index_fileData(ptr))
		p := C.get_path_fileData(ptr)
		f.path = C.GoString(p)
		C.free(unsafe.Pointer(p))
		f.completedLength = int(C.get_completedLength_fileData(ptr))
		if int(C.get_selected_fileData(ptr)) == 0 {
			f.selected = true
		} else {
			f.selected = false
		}
		//Uris not added due to high cost with loops
		files = append(files, f)
	}
	return files
}

func (d aria2go) getGlobalStat() GlobalStat {
	var globalStat GlobalStat
	gs := C.getGlobalStat_aria2go()
	globalStat.downloadSpeed = int(C.get_downloadSpeed_globalStat(gs))
	globalStat.uploadSpeed = int(C.get_uploadSpeed_globalStat(gs))
	globalStat.numActive = int(C.get_numActive_globalStat(gs))
	globalStat.numWaiting = int(C.get_numWaiting_globalStat(gs))
	globalStat.numStopped = int(C.get_numStopped_globalStat(gs))
	C.free(gs)
	return globalStat
}
