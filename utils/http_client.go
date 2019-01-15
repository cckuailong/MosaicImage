package utils

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// http request func
func Http_req(uri string, value url.Values, method string, headers map[string]string) (*http.Response, error){
	var pv io.Reader
	if value != nil{
		post_value := value.Encode()
		pv = strings.NewReader(post_value)
	}else{
		pv=nil
	}
	req, err := http.NewRequest(method, uri, pv)
	if err != nil{
		return nil, err
	}
	if headers != nil{
		for k,v := range(headers){
			req.Header.Add(k, v)
		}
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil{
		return nil, err
	}
	return resp, nil
}
