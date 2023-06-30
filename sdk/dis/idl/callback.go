package idl

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

type CallbackData struct {
	FromDoi    string            `json:"from_doi,omitempty"`    //消息发起者
	ToDoi      string            `json:"to_doi,omitempty"`      //消息通知者
	DataDoi    string            `json:"data_doi,omitempty"`    //数据对象
	NotifyType IDL.NotifyType    `json:"notify_type,omitempty"` //消息类型
	Params     interface{}       `json:"params,omitempty"`      //消息参数
	Fields     map[string]string `json:"fields,omitempty"`      //扩展字段
}

func (r *CallbackData) ToString() string {
	return converter.ToString(r)
}

func (r *CallbackData) Send(url string) (*ApiDisResponse, error) {
	method := "POST"
	payload := strings.NewReader(r.ToString())

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	resp := &ApiDisResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
