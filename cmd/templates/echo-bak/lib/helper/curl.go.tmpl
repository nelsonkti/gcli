package helper

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//返回体
type HttpResponse struct {
	Code int32
	Data string
	Msg  string
}
type errorT struct {
	errorMsg string
}

func (e *errorT) Error() string {
	return e.errorMsg
}

func newError(errorMsg string) *errorT {
	e := errorT{errorMsg: errorMsg}
	return &e
}

/**
*method  请求方法 (get,post.....)
*urlAddress (url)
*parameter(请求参数  map)
 */
func Http(method, urlAddress string, parameter map[string]string, header map[string]string) *HttpResponse {
	Url, _ := url.Parse(urlAddress)
	params := url.Values{}

	//body
	for key, value := range parameter {
		params.Add(key, value)
	}

	var req *http.Request
	method = strings.ToUpper(method)

	if method == "GET" {
		Url.RawQuery = params.Encode()
		req, _ = http.NewRequest(method, Url.String(), nil)
	} else {
		parms := ioutil.NopCloser(strings.NewReader(params.Encode()))
		req, _ = http.NewRequest(method, Url.String(), parms)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//header
	for k, v := range header {
		req.Header.Add(k, v)
	}

	var client *http.Client = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3) //开始请求连接几秒
				if err != nil {
					return nil, newError("connection timed out")
				}
				return c, nil
			},
			MaxIdleConnsPerHost:   10,               //保持连接几秒
			ResponseHeaderTimeout: time.Second * 10, //连接返回几秒
		},
	}

	httpResponse := &HttpResponse{Code: 0, Msg: "ok"}
	res, err := client.Do(req)

	if err != nil {
		httpResponse.Code = 1
		httpResponse.Msg = err.Error()
	} else {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		httpResponse.Code = 0
		httpResponse.Data = string(body)
	}
	return httpResponse
}
