package viewlyrics

import (
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	response := `<?xml version="1.0" encoding='utf-8'?>
<return orgcmd="searchV1" result="OK" badrc="100" ls_dd="1" server_url="http://minilyrics.com/">
   <fileinfo link="link/to/file.lrc" artist="artist1" title="title1" album="album1" uploader="uploadername" timelength="189" downloads="106"/>
   <fileinfo link="link/to/file2.lrc" artist="artist2" title="title2"/>
</return>`
	list, err := parseFileList([]byte(response))

	if err != nil {
		t.Errorf("Didn't expect an error, received %s", err)
	}

	if count := len(list); count != 2 {
		t.Errorf("Expected filelist to contain 2 files, received %d", count)
	}

	expected := File{
		Link:      "link/to/file.lrc",
		Artist:    "artist1",
		Title:     "title1",
		Album:     "album1",
		Downloads: 106,
	}

	if !reflect.DeepEqual(expected, *list[0]) {
		t.Error("Received file didnt match", expected, *list[0])
	}

}
