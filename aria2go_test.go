package aria2

import "testing"

func TestMain(m *testing.M){
	downloader := New()
	downloader.init_aria2go()
}