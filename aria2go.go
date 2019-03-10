package aria2go

// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L. -l aria2
// #include "aria2go.h"
// #include <stdlib.h>
import "C"
import (
	"time"
	"unsafe"
)

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

type Aria2go struct {
}

type Gid struct {
	ptr unsafe.Pointer
}

type FileData struct {
	Index           int
	Path            string
	Length          int64
	CompletedLength int64
	Selected        bool
	//uris []string
}

type GlobalStat struct {
	DownloadSpeed int
	UploadSpeed   int
	NumActive     int
	NumWaiting    int
	NumStopped    int
}

type BtMetaInfoData struct {
	//announce list not implemented
	Comment      string
	CreationDate time.Time
	MultiMod     bool
	SingleMod    bool
	Name         string
	Valid        bool
}

func New() Aria2go {
	var ret Aria2go
	C.init_aria2go()
	return ret
}

func (d Aria2go) Init_aria2go_session(keepRunning bool) {
	C.finalize_aria2go()
	if keepRunning {
		C.init_aria2go_session(C.int(1))
	} else {
		C.init_aria2go_session(C.int(0))
	}
}

func (d Aria2go) Run() int {
	return int(C.run_aria2go(C.int(0)))
}

func (d Aria2go) RunUntillFinished() {
	//Returns when all downloads finisged
	C.keepruning_aria2go()
}

func (d Aria2go) KeepRunning() {
	C.finalize_aria2go()
	C.init_aria2go_session(C.int(1))
	go func() {
		C.keepruning_aria2go()
	}()
}

func (d Aria2go) RunOnce() int {
	return int(C.run_aria2go(C.int(1)))
}

func (d Aria2go) GidToHex(gid Gid) string {
	p := C.gidToHex_aria2go(gid.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (d Aria2go) HexToGid(s string) Gid {
	var ret Gid
	ret.ptr = C.hexToGid_aria2go(C.CString(s))
	return ret
}

func (d Aria2go) IsNull(g Gid) bool {
	if C.isNull_aria2go(g.ptr) == 0 {
		return false
	}
	return true
}

func (d Aria2go) AddUriInPosition(uri string, position int) Gid {
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri), C.int(position))
	return ret
}

func (d Aria2go) AddUri(uri string) Gid {
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri), C.int(-1))
	return ret
}

func (d Aria2go) AddMetalinkInPosition(file_location string, position int) []Gid {
	var gids []Gid
	var gid Gid
	l := int(C.addMetalink_aria2go(C.CString(file_location), C.int(position)))
	for i := 0; i < l; i++ {
		gid.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, gid)
	}
	return gids
}

func (d Aria2go) AddMetaLink(file_location string) []Gid {
	return d.AddMetalinkInPosition(file_location, -1)
}

func (d Aria2go) AddUriToCache(uri string) {
	//You can use this one even if you didnt start session yet
	C.add_uri(C.CString(uri))
}

func (d Aria2go) ClearUriCache() {
	C.clear_uris()
}

func (d Aria2go) AddAllFromCacheWithPosition(p int) Gid {
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(p))
	return ret
}
func (d Aria2go) AddAllFromCache() Gid {
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(-1))
	return ret
}

func (d Aria2go) GetActiveDownload() []Gid {
	var gids []Gid
	l := int(C.getActiveDownload_aria2go())
	for i := 0; i < l; i++ {
		var g Gid
		g.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, g)
	}
	return gids
}

func (d Aria2go) RemoveDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr, C.int(0))
	//Second variable is bool for forcing
}

func (d Aria2go) ForceRemoveDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr, C.int(1))
}

func (d Aria2go) PauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr, C.int(0))
	//Second variable is bool for forcing
}

func (d Aria2go) ForcePauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr, C.int(1))
}

func (d Aria2go) UnpauseDownload(g Gid) {
	C.unpauseDownload_aria2go(g.ptr)
}

//Event Callback
type EventCallback func(DownloadEvent, Gid)

var callback EventCallback //Have to be global due to garbage collector

func (d Aria2go) SetEventCallback(eventCallback EventCallback) {
	callback = eventCallback
}

//export runGoCallBack
func runGoCallBack(event C.enum_DownloadEvent, g unsafe.Pointer) {
	if callback == nil {
		return
	}
	var gid Gid
	gid.ptr = g
	callback(DownloadEvent(event), gid)
}

func (d Aria2go) Finalize() {
	C.finalize_aria2go()
	C.deinit_aria2go()
}

func (g Gid) GetStatus() DownloadStatus {
	return DownloadStatus(C.getStatus_gid(g.ptr))
}

func (g Gid) GetTotalLength() int64 {
	return int64(C.getTotalLength_gid(g.ptr))
}

func (g Gid) GetCompletedLength() int64 {
	return int64(C.getCompletedLength_gid(g.ptr))
}

func (g Gid) GetUploadLength() int64 {
	return int64(C.getUploadLength_gid(g.ptr))
}

func (g Gid) GetBitfield() string {
	p := C.getBitfield_gid(g.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (g Gid) GetDownloadSpeed() int {
	return int(C.getDownloadSpeed_gid(g.ptr))
}

func (g Gid) GetUploadSpeed() int {
	return int(C.getUploadSpeed_gid(g.ptr))
}

func (g Gid) GetInfoHash() string {
	p := C.getInfoHash_gid(g.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (g Gid) GetNumPieces() int {
	return int(C.getNumPieces_gid(g.ptr))
}

func (g Gid) GetConnections() int {
	return int(C.getConnections_gid(g.ptr))
}

func (g Gid) GetErrorCode() int {
	return int(C.getErrorCode_gid(g.ptr))
}

func (g Gid) GetNumFiles() int {
	return int(C.getNumFiles_gid(g.ptr))
}

func (g Gid) GetDir() string {
	p := C.getDir_gid(g.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (g Gid) GetBtMetaInfo() BtMetaInfoData {
	var btMetaInfo BtMetaInfoData
	btmi := C.getBtMetaInfo_gid(g.ptr)
	if btmi == nil {
		btMetaInfo.Valid = false
	} else {
		//Comment
		c := C.get_comment_BtMetaInfo(btmi)
		btMetaInfo.Comment = C.GoString(c)
		C.free(unsafe.Pointer(c))
		//Creation Time
		btMetaInfo.CreationDate = time.Unix(int64(C.get_creationDate_BtMetaInfo(btmi)), 0)
		//Mode
		mode := int(C.get_mode_BtMetaInfo(btmi))
		btMetaInfo.SingleMod = false
		btMetaInfo.MultiMod = false
		if mode == 0 {
			btMetaInfo.SingleMod = true
		} else if mode == 1 {
			btMetaInfo.MultiMod = true
		}
		//Name
		n := C.get_name_BtMetaInfo(btmi)
		btMetaInfo.Name = C.GoString(n)
		C.free(unsafe.Pointer(n))

		C.free(btmi)
	}
	return btMetaInfo
}

func (g Gid) GetFiles() []FileData {
	var files []FileData
	var ptr unsafe.Pointer
	l := int(C.getFiles_gid(g.ptr))
	for i := 0; i < l; i++ {
		var f FileData
		ptr = C.get_element_fileData(C.int(i))
		f.Index = int(C.get_index_fileData(ptr))
		p := C.get_path_fileData(ptr)
		f.Path = C.GoString(p)
		C.free(unsafe.Pointer(p))
		f.Length = int64(C.get_length_fileData(ptr))
		f.CompletedLength = int64(C.get_completedLength_fileData(ptr))
		if int(C.get_selected_fileData(ptr)) == 0 {
			f.Selected = true
		} else {
			f.Selected = false
		}
		//Uris not added due to high cost with loops
		files = append(files, f)
	}
	return files
}

func (d Aria2go) GetGlobalStat() GlobalStat {
	var globalStat GlobalStat
	gs := C.getGlobalStat_aria2go()
	globalStat.DownloadSpeed = int(C.get_downloadSpeed_globalStat(gs))
	globalStat.UploadSpeed = int(C.get_uploadSpeed_globalStat(gs))
	globalStat.NumActive = int(C.get_numActive_globalStat(gs))
	globalStat.NumWaiting = int(C.get_numWaiting_globalStat(gs))
	globalStat.NumStopped = int(C.get_numStopped_globalStat(gs))
	C.free(gs)
	return globalStat
}
