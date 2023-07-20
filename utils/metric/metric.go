package metric

import (
	"encoding/json"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/fuxi-inc/dip-common-lib/utils/security"
	"github.com/zeromicro/go-zero/core/netx"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"unicode/utf8"
)

type MValue struct {
	Value int `json:"value"`
}

type Metric struct {
	Dimensions map[string]string `json:"dimensions"`
	GroupId    int               `json:"groupId"`
	MetricName string            `json:"metricName"`
	Period     int               `json:"period"`
	Time       int64             `json:"time"`
	Type       int               `json:"type"`
	Values     *MValue           `json:"values"`
}

type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type Client struct {
	Logger       *zap.Logger //日志组件
	Host         string      `json:"host"`
	AccessKey    string      `json:"access_key"`
	AccessSecret string      `json:"access_secret"`
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InitLogger(logger *zap.Logger) *Client {
	c.Logger = logger
	return c
}

func (c *Client) InitHost(host string) *Client {
	c.Host = host
	return c
}

func (c *Client) HostStr() string {
	parser, err := url.Parse(c.Host)
	if err != nil {
		return ""
	}
	return parser.Host
}

func (c *Client) InitAccess(key, secret string) *Client {
	c.AccessKey = key
	c.AccessSecret = secret
	return c
}

func (c *Client) GenSignString(contentMd5, contentDate string) string {
	signString := constants.POST + "\n" + contentMd5 + "\n" + constants.MIMEApplicationJSON + "\n" + contentDate + "\n" + "x-cms-api-version:1.0" + "\n" + "x-cms-ip:" + netx.InternalIp() + "\n" + "x-cms-signature:hmac-sha1" + "\n" + "/metric/custom/upload"
	return security.HmacSHA1(signString, c.AccessSecret)
}

func (c *Client) Push(m []Metric) error {
	requestBody := converter.ToString(m)

	pushUrl := c.Host + "/metric/custom/upload"
	method := constants.POST
	payload := strings.NewReader(requestBody)

	client := &http.Client{}
	req, err := http.NewRequest(method, pushUrl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Push request,error:%s", err.Error()))
		return err
	}

	contentLength := converter.IntToString(utf8.RuneCountInString(requestBody))
	contentMd5 := security.MD5(requestBody)
	contentDate := time.Now().Format(time.RFC1123)
	req.Header.Add(constants.HeaderAuthorization, fmt.Sprintf("%s:%s", c.AccessKey, c.GenSignString(contentMd5, contentDate)))
	req.Header.Add(constants.HeaderContentLength, contentLength)
	req.Header.Add(constants.HeaderContentMd5, contentMd5)
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)
	req.Header.Add(constants.HeaderDate, contentDate)
	req.Header.Add(constants.HeaderHost, c.HostStr())
	req.Header.Add(constants.HeaderUserAgent, "dip-common-lib")
	req.Header.Add("x-cms-api-version", "1.0")
	req.Header.Add("x-cms-signature", "hmac-sha1")
	req.Header.Add("x-cms-ip", netx.InternalIp())

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Push client.Do,error:%s", err.Error()))
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Push ioutil.ReadAll,error:%s", err.Error()))
		return err
	}

	response := &Response{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Push response.Unmarshal,error:%s", err.Error()))
		return err
	}
	if response.Code != "200" {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", converter.ToString(response.Msg)))
		return fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response.Msg))
	}

	return nil
}
