/**
** @创建时间 : 2022/3/21 09:39
** @作者 : fzy
 */
package api

import (
	"demod/lib/logger"
	"demod/util/xnsq/service/registry"
	"io/ioutil"
	"net/http"
	"strings"
)

var NsqApiClient *Client

type Client struct {
	addr  string
	Port  string
	opt   registry.Options
	topic Topic
}

func NewClient(opt registry.Options) *Client {

	NsqApiClient := &Client{
		opt: opt,
	}

	NsqApiClient.buildAdminAddr()

	return NsqApiClient
}

func (c *Client) Topic() *Topic {
	return &Topic{client: c}
}

func (c *Client) buildAdminAddr() {
	if !strings.Contains(c.opt.NSQAdminAddress, "://") {
		https := "https://"
		if c.opt.Env == "local" || c.opt.Env == "develop" {
			https = "http://"
		}
		c.addr = https + c.opt.NSQAdminAddress
	}
}

func (c *Client) Get(router string) ([]byte, error) {
	url := c.addr + "/" + router

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	logger.Sugar.Error(err)
	return body, err
}

func (c *Client) Delete(router string) ([]byte, error) {
	url := c.addr + "/" + router

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return body, err
}
