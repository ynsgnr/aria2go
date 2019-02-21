package aria2

import "testing"
//import "gotest.tools/assert"

var downloader Downloader
var session Session

func TestMain(t *testing.M){
	downloader = New()
	downloader.init_aria2go()
}

func TestSession(t *testing.T){
	session = downloader.init_aria2go_session()
}

func TestRun(t *testing.T){
	downloader.run(session)
}
