/**
** @创建时间 : 2022/3/21 09:38
** @作者 : fzy
 */
package api

import (
	"demod/lib/logger"
	"encoding/json"
)

type Topic struct {
	client *Client
}

func (t *Topic) Client(client *Client) *Topic {
	return &Topic{
		client: client,
	}
}

type QueryAllData struct {
	Topics  []string `json:"topics"`
	Message string   `json:"message"`
}

// 获取所有的topic
func (t *Topic) QueryAll() ([]string, error) {
	body, err := t.client.Get("api/topics")

	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	var queryAllData QueryAllData

	_ = json.Unmarshal(body, &queryAllData)

	return queryAllData.Topics, nil
}

// 删除 topic
func (t *Topic) Delete(topic string) error {
	_, err := t.client.Delete("api/topics/" + topic)

	if err != nil {
		logger.Sugar.Error(err)
		return err
	}

	return nil
}
