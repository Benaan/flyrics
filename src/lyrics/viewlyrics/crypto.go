package viewlyrics

import (
	"crypto/md5"
	"io"
	"io/ioutil"
)

func encode(query string) []byte {
	hash := md5.New()
	hash.Write([]byte(query))
	hash.Write([]byte("Mlv1clt4.0"))
	request := []byte{2, 0, 4, 0, 0, 0}
	request = append(request, hash.Sum(nil)...)
	request = append(request, []byte(query)...)
	return request
}

func decode(response io.Reader) ([]byte, error) {
	body, err := ioutil.ReadAll(response)
	if err != nil {
		return nil, err
	}
	magicKey := body[1]
	result := []byte{}
	for _, char := range body[22:] {
		result = append(result, char^magicKey)
	}
	return result, nil
}
