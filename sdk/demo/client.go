package demo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/demo/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

/**
*  请求方式类如：
*    response,err := NewClient().
					InitLogger(logger).
					InitDis("http://dis.viv.cn").
					ApiRequestDemo( &idl.ApiDemoRequest{})
*/

type Client struct {
	Logger  *zap.Logger //日志组件
	DisHost string
	DaoHost string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InitLogger(logger *zap.Logger) *Client {
	c.Logger = logger
	return c
}

func (c *Client) InitDis(disHost string) *Client {
	c.DisHost = disHost
	return c
}

func (c *Client) InitDao(daoHost string) *Client {
	c.DaoHost = daoHost
	return c
}

func (c *Client) ApiRequestDemo(ctx *gin.Context, request *idl.ApiDemoRequest) (*idl.ApiDemoResponse, error) {
	disUrl := c.DisHost + "/v1/api/userJobExperience"
	method := constants.POST

	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disUrl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	req.Header.Add(constants.HeaderAuthorization, "<Authorization>")
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}

	response := &idl.ApiDemoResponse{}
	err = response.Unmarshal(body)
	return response, err
}
