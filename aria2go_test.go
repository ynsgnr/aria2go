package aria2

import "testing"
import "gotest.tools/assert"
import "os"
import "fmt"

var downloader Downloader
var session Session
var gid Gid
var gid_position Gid

func TestMain(m *testing.M){
	downloader = New()
	downloader.init_aria2go()
	os.Exit(m.Run())
}

func TestAll(t *testing.T){
	t.Run("Session init",func(t *testing.T){
		session = downloader.init_aria2go_session()
	})
	t.Run("Run Aria2",func(t *testing.T){
		downloader.run(session)
	})
	t.Run("Add uri",func(t *testing.T){
		//Todo maybe check files md5 with argon2
		gid = downloader.addUri("https://www.w3.org/History/1989/Image1.gif") //oldest file in the internet, unlikely to be deleted
	})
	t.Run("Add uri in position",func(t *testing.T){
		//Todo maybe check files md5 with argon2
		gid_position = downloader.addUriInPosition("https://www.w3.org/History/1989/Image2.gif",0) //oldest file in the internet, unlikely to be deleted
		//Todo add get current position and check it		
	})
	t.Run("is gid null",func(t *testing.T){		
		assert.Equal(t,downloader.isNull(gid),false)
	})
	t.Run("gid to hex",func(t *testing.T){		
		downloader.gidToHex(gid)
		//hex can be anything no way to check
	})
	t.Run("hex to gid",func(t *testing.T){
		hex := downloader.gidToHex(gid)
		fmt.Printf("%s\n",hex)
		gid_converted := downloader.hexToGid(hex)
		assert.Equal(t,gid,gid_converted)
	})
}

func TestArrayDummy(t *testing.T){
	gids := downloader.arraytest()
	assert.DeepEqual(t,gids,[]int{1,12,35,16,43,67})
}