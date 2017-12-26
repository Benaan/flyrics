package viewlyrics

import "encoding/xml"

type File struct {
	Link      string `xml:"link,attr"`
	Artist    string `xml:"artist,attr"`
	Title     string `xml:"title,attr"`
	Album     string `xml:"album,attr"`
	Downloads int    `xml:"downloads,attr"`
	Rating    string `xml:"rating,attr"`
}

type FileList struct {
	Files []*File `xml:"fileinfo"`
}

func parseFileList(input []byte) ([]*File, error) {
	var list FileList
	err := xml.Unmarshal(input, &list)
	return list.Files, err
}
