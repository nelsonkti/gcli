package helper

import (
	"github.com/valyala/fasthttp"
)

func Post(url string, param interface{}) (error, string) {
	return Request("POST", url, param)
}

func Get(url string, param interface{}) (error, string) {
	return Request("GET", url, param)
}

func Request(method, url string, param interface{}) (error, string) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	req.SetRequestURI(url)
	req.Header.SetMethod(method)
	req.Header.SetContentType("application/json")

	req.SetBody(JsonMarshalByte(param))
	err := fasthttp.Do(req, resp)

	if err != nil {
		return err, ""
	}

	b := resp.Body()

	return nil, string(b)
}
