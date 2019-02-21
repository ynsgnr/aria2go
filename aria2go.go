package aria2
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L. -l aria2 -l aria2go
// #include "aria2go.h"
// #include <stdlib.h>
import "C"
import "unsafe"

type Downloader struct {
	ptr unsafe.Pointer
}

type Session struct {
	ptr unsafe.Pointer
}

type Gid struct {
	ptr unsafe.Pointer
}

func New() Downloader {
	var ret Downloader;
	ret.ptr = C.new_aria2go();
	return ret;
}

func (d Downloader)del_aria2go(){
	C.del_aria2go(d.ptr);
}

func (d Downloader)init_aria2go(){
	C.init_aria2go(d.ptr);
}

func (d Downloader)init_aria2go_session()Session{
	var ret Session;
	ret.ptr = C.init_aria2go_session(d.ptr)
	return ret
}

func (d Downloader)run(s Session){
	C.run_aria2go(d.ptr,s.ptr)
}

func (d Downloader)gidToHex(gid Gid) string{
	p := C.gidToHex_aria2go(d.ptr,gid.ptr)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

func (d Downloader)hexToGid(s string) Gid{
	var ret Gid
	ret.ptr = C.hexToGid_aria2go(d.ptr,C.CString(s))
	return ret
}

func (d Downloader)isNull(g Gid) bool{
	if C.isNull_aria2go(d.ptr,g.ptr)==0 { return false }
	return true
}

func (d Downloader)addUriInPosition(uri string,position int) Gid{
	var ret Gid
	ret.ptr = C.addUri_aria2go(d.ptr,C.CString(uri),C.int(position))
	return ret
}

func (d Downloader)addUri(uri string) Gid{
	var ret Gid
	ret.ptr = C.addUri_aria2go(d.ptr,C.CString(uri),C.int(-1))
	return ret
}