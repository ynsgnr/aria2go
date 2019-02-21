package aria2

import "testing"
import "gotest.tools/assert"

var downloader Downloader
var session Session
var gid Gid
var gid_position Gid

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

func TestAddUri(t *testing.T){
	//Todo maybe check files md5 with argon2
	gid = downloader.addUri("https://www.w3.org/History/1989/Image1.gif") //oldest file in the internet, unlikely to be deleted
}

func TestAddUriInPosition(t *testing.T){
	//Todo maybe check files md5 with argon2
	gid_position = downloader.addUriInPosition("https://www.w3.org/History/1989/Image2.gif",0) //oldest file in the internet, unlikely to be deleted
	//Todo add get current position and check it
}

func TestisNull(t *testing.T){
	assert.Equal(t,downloader.isNull(gid),true)
}

func TestGidToHex(t *testing.T){
	downloader.gidToHex(gid)
	//hex can be anything no way to check
}

func TestHexToGid(t *testing.T){
	hex := downloader.gidToHex(gid)
	gid_converted := downloader.hexToGid(hex)
	assert.Equal(t,gid,gid_converted)
}

func TestArrayDummy(t *testing.T){
	gids := downloader.arraytest()
	assert.Equal(t,gids,[]int{1,12,35,16,43,67})
}