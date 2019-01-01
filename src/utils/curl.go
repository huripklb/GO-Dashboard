package utils

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

// ReqHeader : struct type to store request header
// author : Huripto Sugandi
// date created : 6 Dec 2018
type ReqHeader struct {
	ContentType string
	Content     string
}

// CallRequest : function to get response from external URL/API
// author : Huripto Sugandi
// date created : 6 Dec 2018
func CallRequest(url string, jsonData []byte, method string, header *ReqHeader) ([]byte, bool) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if header != nil {
		req.Header.Set(header.ContentType, header.Content)
	}

	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	if 200 != resp.StatusCode {
		log.Printf("%d Request NOT OK!", resp.StatusCode)
		return nil, false
	}

	return body, true
}
