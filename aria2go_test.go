package aria2go

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"gotest.tools/assert"
)

var downloader Aria2go
var gid Gid
var gid_position Gid
var gid_to_pause Gid

var start_count = 0
var pause_coount = 0
var stop_count = 0
var complete_count = 0
var BT_complete_count = 0
var error_count = 0

//Test files
//Oldest files in internet unlikely to be deleted
const file_path_1 string = "Image1.gif"
const file_path_1_1 string = "Image1.1.gif"
const file_md5_1 string = "92d36637578442ae99c4171b88101610"
const file_link_1 string = "https://www.w3.org/History/1989/Image1.gif"
const file_path_2 string = "Image2.gif"
const file_md5_2 string = "c3d386b7effd6a520a96dc5c4eee0189"
const file_link_2 string = "https://www.w3.org/History/1989/Image2.gif"
const file_path_3 string = "Image3.gif"
const file_md5_3 string = "4211beb988c8f74a9dbb546efaa52bcc"
const file_link_3 string = "https://www.w3.org/History/1989/Image3.gif"
const file_path_4 string = "proposal-magnify.gif"
const file_md5_4 string = "c3d386b7effd6a520a96dc5c4eee0189"
const file_link_4 string = "https://www.w3.org/History/1989/proposal-magnify.gif"

func TestMain(m *testing.M) {
	downloader = New()
	downloader.SetEventCallback(func(event DownloadEvent, g Gid) {
		fmt.Printf("Callback Called:\n")
		s := g.GetStatus()
		switch event {
		case EVENT_ON_DOWNLOAD_START:
			fmt.Printf("Download Start: %s\n", s)
			start_count++
		case EVENT_ON_DOWNLOAD_PAUSE:
			fmt.Printf("Download Pause: %s\n", s)
			pause_coount++
		case EVENT_ON_DOWNLOAD_STOP:
			fmt.Printf("Download Stop: %s\n", s)
			stop_count++
		case EVENT_ON_DOWNLOAD_COMPLETE:
			fmt.Printf("Download Complete: %s\n", s)
			complete_count++
		case EVENT_ON_BT_DOWNLOAD_COMPLETE:
			fmt.Printf("Download BT Complete: %s\n", s)
			BT_complete_count++
		default:
			fmt.Printf("Download Error: %s\n", s)
			error_count++
		}
	})
	ret := m.Run()
	//Delete downloaded files
	os.Remove(file_path_1)
	os.Remove(file_path_1_1)
	os.Remove(file_path_2)
	os.Remove(file_path_3)
	os.Remove(file_path_4)
	os.Exit(ret)
	downloader.Finalize()
}

func TestAll(t *testing.T) {
	t.Run("Check if function calls protected against nullptr", func(t *testing.T) {
		//Since session not initilized, expected case is nothing to be happen
		//If test crashes on this part check pointers and add protection
		downloader.AddUri(file_link_1)
		g := downloader.AddUriInPosition(file_link_2, 0)
		downloader.IsNull(g)
		h := downloader.GidToHex(g)
		downloader.HexToGid(h)
		downloader.GetActiveDownload()
		downloader.RemoveDownload(g)
		downloader.ForceRemoveDownload(g)
		downloader.PauseDownload(g)
		downloader.UnpauseDownload(g)
		downloader.ForcePauseDownload(g)
		downloader.UnpauseDownload(g)
	})
	t.Run("Init Session", func(t *testing.T) {
		downloader.Init_aria2go_session(true)
	})
	t.Run("Init Session Second time", func(t *testing.T) {
		downloader.Init_aria2go_session(false)
	})
	t.Run("Run", func(t *testing.T) {
		downloader.KeepRunning()
	})
	t.Run("Add uri", func(t *testing.T) {
		gid = downloader.AddUri(file_link_1)
		time.Sleep(2 * time.Second)
	})
	t.Run("Add uri in position", func(t *testing.T) {
		gid_position = downloader.AddUriInPosition(file_link_2, 0)
		time.Sleep(2 * time.Second)
		//Todo add get current position and check it
	})
	t.Run("is gid null", func(t *testing.T) {
		assert.Equal(t, downloader.IsNull(gid), false)
	})
	t.Run("gid to hex", func(t *testing.T) {
		downloader.GidToHex(gid)
		//hex can be anything no way to check
	})
	t.Run("hex to gid", func(t *testing.T) {
		hex := downloader.GidToHex(gid)
		gid_converted := downloader.HexToGid(hex)
		assert.Equal(t, gid, gid_converted)
	})
	t.Run("Check Status in position", func(t *testing.T) {
		sp := gid_position.GetStatus()
		assert.Equal(t, sp, DOWNLOAD_COMPLETE)
	})
	t.Run("Check Status", func(t *testing.T) {
		sp := gid.GetStatus()
		assert.Equal(t, sp, DOWNLOAD_COMPLETE)
	})
	t.Run("Check Event Counts", func(t *testing.T) {
		assert.Equal(t, start_count, 2)
		assert.Equal(t, complete_count, 2)
	})
	t.Run("Check Downloaded Files MD5", func(t *testing.T) {
		md5, err := hash_file_md5(file_path_1)
		if err != nil {
			t.Error("File not found: " + file_path_1)
		}
		assert.Equal(t, md5, file_md5_1)
		md5, err = hash_file_md5(file_path_2)
		if err != nil {
			t.Error("File not found: " + file_path_2)
		}
		assert.Equal(t, md5, file_md5_2)
	})
	t.Run("Uri Cache", func(t *testing.T) {
		downloader.AddUriToCache("Test")
	})
	t.Run("Claer Uri Cache", func(t *testing.T) {
		downloader.ClearUriCache()
	})
	t.Run("Add All From Cache", func(t *testing.T) {
		downloader.AddUriToCache(file_link_3)
		downloader.AddUriToCache(file_link_4) //This file will not be downloaded because it does not point to the same file
		gid = downloader.AddAllFromCache()
		time.Sleep(2 * time.Second)
	})
	t.Run("Check Downloaded Files From Cache MD5", func(t *testing.T) {
		md5, err := hash_file_md5(file_path_3)
		if err != nil {
			t.Error("File not found: " + file_path_3)
		}
		assert.Equal(t, md5, file_md5_3)
		md5, err = hash_file_md5(file_path_4)
		if err == nil {
			t.Error("Links that does not point to same file downloaded: " + file_path_4)
		}
	})
	t.Run("Get Active Download", func(t *testing.T) {
		downloader.GetActiveDownload()
	})
	t.Run("Remove Download", func(t *testing.T) {
		downloader.RemoveDownload(gid)
		//TODO check events
	})
	t.Run("Remove Download Force", func(t *testing.T) {
		downloader.ForceRemoveDownload(gid)
	})
	t.Run("Pause Download", func(t *testing.T) {
		gid_to_pause = downloader.AddUri(file_link_1)
		downloader.PauseDownload(gid_to_pause)
	})
	t.Run("Unpause Download", func(t *testing.T) {
		downloader.UnpauseDownload(gid_to_pause)
	})
	t.Run("Force Pause Download", func(t *testing.T) {
		gid_to_pause = downloader.AddUri(file_link_1)
		downloader.ForcePauseDownload(gid_to_pause)
	})
	t.Run("Unpause Force Paused Download", func(t *testing.T) {
		downloader.UnpauseDownload(gid_to_pause)
		time.Sleep(2 * time.Second)
		md5, err := hash_file_md5(file_path_1_1)
		if err != nil {
			t.Error("File not found: " + file_path_3)
		}
		assert.Equal(t, file_md5_1, md5)
	})
	t.Run("Gid Functions", func(t *testing.T) {
		gid = downloader.AddUri(file_link_4)
		//TODO find a compatible file to test these functions, and add assert
		gid.GetStatus()
		gid.GetTotalLength()
		gid.GetBitfield()
		gid.GetDownloadSpeed()
		gid.GetUploadSpeed()
		gid.GetInfoHash()
		gid.GetNumPieces()
		gid.GetConnections()
		gid.GetErrorCode()
		gid.GetNumFiles()
		time.Sleep(2 * time.Second)
	})
	t.Run("File Data Path", func(t *testing.T) {
		files := gid.GetFiles()
		s := strings.Split(files[0].Path, "/")
		assert.Equal(t, file_path_4, s[len(s)-1])
	})
	t.Run("Global Stats", func(t *testing.T) {
		gs := downloader.GetGlobalStat()
		assert.Equal(t, 6, gs.NumStopped)
	})
	t.Run("BTMI basic", func(t *testing.T) {
		//TODO add a torrent file and do other tests
		btmi := gid.GetBtMetaInfo()
		assert.Equal(t, btmi.Valid, false)
	})
}

func hash_file_md5(filePath string) (string, error) {
	//This function from https://mrwaggel.be/post/generate-md5-hash-of-a-file-in-golang/
	//Credits to Mr.Waggel
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}
