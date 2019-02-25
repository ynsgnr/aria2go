package aria2
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L. -l aria2 -l aria2go
// #include "aria2go.h"
// #include <stdlib.h>
import "C"
import "unsafe"

//ENUMS
type DownloadEvent int
const(
	EVENT_ON_DOWNLOAD_START	DownloadEvent = C.EVENT_ON_DOWNLOAD_START
	EVENT_ON_DOWNLOAD_PAUSE DownloadEvent = C.EVENT_ON_DOWNLOAD_PAUSE
    EVENT_ON_DOWNLOAD_STOP DownloadEvent = C.EVENT_ON_DOWNLOAD_STOP
    EVENT_ON_DOWNLOAD_COMPLETE DownloadEvent = C.EVENT_ON_DOWNLOAD_COMPLETE
    EVENT_ON_DOWNLOAD_ERROR DownloadEvent = C.EVENT_ON_DOWNLOAD_ERROR
	EVENT_ON_BT_DOWNLOAD_COMPLETE DownloadEvent = C.EVENT_ON_BT_DOWNLOAD_COMPLETE
)

type aria2go struct {
	session unsafe.Pointer
}

type Gid struct {
	ptr unsafe.Pointer
}

func New() (aria2go) {
	var ret aria2go
	C.init_aria2go()
	ret.session = C.init_aria2go_session()
	C.run_aria2go(ret.session)
	return ret
}

func (d aria2go)gidToHex(gid Gid) string{
	p := C.gidToHex_aria2go(gid.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (d aria2go)hexToGid(s string) Gid{
	var ret Gid
	ret.ptr = C.hexToGid_aria2go(C.CString(s))
	return ret
}

func (d aria2go)isNull(g Gid) bool{
	if C.isNull_aria2go(g.ptr)==0 { return false }
	return true
}

func (d aria2go)addUriInPosition(uri string,position int) Gid{
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri),C.int(position))
	return ret
}

func (d aria2go)addUri(uri string) Gid{
	var ret Gid
	ret.ptr = C.addUri_aria2go(C.CString(uri),C.int(-1))
	return ret
}

func (d aria2go)addMetalinkInPosition(file_location string,position int) []Gid{
	var gids []Gid
	var gid Gid
	l :=int(C.addMetalink_aria2go(C.CString(file_location),C.int(position)))
	for i := 0; i < l; i++ {
		gid.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, gid)
	}
	return gids
}

func (d aria2go)addMetaLink(file_location string) []Gid{
	return d.addMetalinkInPosition(file_location,-1)
}

func (d aria2go)addUriToCache(uri string){
	C.add_uri(C.CString(uri))
}

func (d aria2go)clearUriCache(){
	C.clear_uris()
}

func (d aria2go)addAllFromCacheWithPosition(p int) Gid{
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(p))
	return ret
}
func (d aria2go)addAllFromCache() Gid{
	var ret Gid
	ret.ptr = C.add_all_from_cache(C.int(-1))
	return ret
}

func (d aria2go)getActiveDownload() []Gid{
	var gids []Gid
	l :=int(C.getActiveDownload_aria2go())
	for i := 0; i < l; i++ {
		var g Gid
		g.ptr = C.get_element_gid(C.int(i))
		gids = append(gids, g)
	}
	return gids
}

func (d aria2go)removeDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr,C.int(0)) 
	//Second variable is bool for forcing
}

func (d aria2go)forceRemoveDownload(g Gid) {
	C.removeDownload_aria2go(g.ptr,C.int(1))
}

func (d aria2go)pauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr,C.int(0)) 
	//Second variable is bool for forcing
}

func (d aria2go)forcePauseDownload(g Gid) {
	C.pauseDownload_aria2go(g.ptr,C.int(1))
}

func (d aria2go)unpauseDownload(g Gid) {
	C.unpauseDownload_aria2go(g.ptr)
}

//Event Callback
type EventCallback func(DownloadEvent,Gid)
var callback EventCallback //Have to be global due to garbage collector

func (d aria2go)setEventCallback( eventCallback EventCallback){
	callback = eventCallback
}

//export runGoCallBack
func runGoCallBack(event C.enum_DownloadEvent, g unsafe.Pointer){
	var gid Gid
	gid.ptr = g
	callback(DownloadEvent(DownloadEvent(event)),gid)
}