package aria2
// #cgo CXXFLAGS: -std=c++11
// #cgo LDFLAGS: -L. -l aria2 -l aria2go
// #include "aria2go.h"
import "C"
import "unsafe"

type Downloader struct {
	ptr unsafe.Pointer
}

type Session struct {
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