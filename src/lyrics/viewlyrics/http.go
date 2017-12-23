package viewlyrics

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/benaan/flyrics/src/model"
)

var client = &http.Client{
	Timeout: 2 * time.Second,
}

func sendRequest(request []byte) (io.ReadCloser, error) {
	req, err := http.NewRequest("POST", "http://search.crintsoft.com/searchlyrics.htm", bytes.NewReader(request))
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "MiniLyrics")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return getResponse(client.Do(req))
}

func createRequest(song *model.Song) []byte {
	header := `<?xml version='1.0' encoding='utf-8' ?><searchV1 artist="%s" title="%s" client="MiniLyrics" OnlyMatched="1" RequestPage='%d' />`
	query := fmt.Sprintf(header, song.Artist, song.Title, 0)
	return encode(query)
}

func getFile(url string) (io.ReadCloser, error) {
	return getResponse(client.Get("http://viewlyrics.com/" + url))
}

func getResponse(response *http.Response, err error) (io.ReadCloser, error) {
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		response.Body.Close()
		return nil, errors.New("Request didn't return statuscode 200")
	}

	return response.Body, nil
}
