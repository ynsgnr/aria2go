package aria2

import "testing"

var downloader Downloader
var session Session

func TestMain(m *testing.M){
	downloader = New()
	downloader.init_aria2go()
}

func TestSession(m *testing.T){
	session = downloader.init_aria2go_session()
}

func TestRun(m *testing.T){
	downloader.run(session)
}