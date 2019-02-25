package aria2

import "testing"
import "gotest.tools/assert"
import "os"

import "fmt"

var downloader aria2go
var gid Gid
var gid_position Gid
var gid_to_pause Gid

func TestMain(m *testing.M){
	downloader = New()
	os.Exit(m.Run())
}

func TestAll(t *testing.T){
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
		gid_converted := downloader.hexToGid(hex)
		assert.Equal(t,gid,gid_converted)
	})
	t.Run("Uri Cache",func(t *testing.T){
		downloader.addUriToCache("Test")
	})
	t.Run("Claer Uri Cache",func(t *testing.T){
		downloader.clearUriCache()
	})
	t.Run("Add All From Cache",func(t *testing.T){
		downloader.addUriToCache("https://www.w3.org/History/1989/Image3.gif")
		downloader.addUriToCache("https://www.w3.org/History/1989/proposal-magnify.gif")
		downloader.addAllFromCache()
	})
	t.Run("Get Active Download",func(t *testing.T){
		downloader.getActiveDownload()
	})
	t.Run("Remove Download",func(t *testing.T){
		downloader.removeDownload(gid)
		//TODO check events
	})
	t.Run("Remove Download Force",func(t *testing.T){
		downloader.forceRemoveDownload(gid)
	})
	t.Run("Pause Download",func(t *testing.T){
		gid_to_pause = downloader.addUri("https://www.w3.org/History/1989/Image1.gif")
		downloader.pauseDownload(gid_to_pause)
	})
	t.Run("Unpause Download",func(t *testing.T){
		downloader.unpauseDownload(gid_to_pause)
	})
	t.Run("Force Pause Download",func(t *testing.T){
		gid_to_pause = downloader.addUri("https://www.w3.org/History/1989/Image1.gif")
		downloader.forcePauseDownload(gid_to_pause)
	})
	t.Run("Unpause Force Paused Download",func(t *testing.T){
		downloader.unpauseDownload(gid_to_pause)
	})
}

func TestCallback(t *testing.T){
	downloader.setEventCallback(func(){
		fmt.Printf("Test Calback")
	})
	downloader.callCallback()
}