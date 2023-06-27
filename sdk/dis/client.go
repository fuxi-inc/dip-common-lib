package dis

import (
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
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

// 数据对象属性注册
func (c *Client) ApiDOCreate(ctx *gin.Context, request *idl.ApiDOCreateRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 数据对象属性更新
func (c *Client) ApiDOUpdate(ctx *gin.Context, request *idl.ApiDOUpdateRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 数据对象属性删除
func (c *Client) ApiDODelete(ctx *gin.Context, request *idl.ApiDODeleteRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 授权发起
func (c *Client) ApiAuthInit(ctx *gin.Context, request *idl.ApiAuthInitRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 授权确认
func (c *Client) ApiAuthConf(ctx *gin.Context, request *idl.ApiAuthConfRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 注册数据查询
func (c *Client) ApiRegDataQuery(ctx *gin.Context, request *idl.ApiRegDataRequest) (*idl.ApiRegDataQueryResponse, error) {
	return nil, nil
}

// 数据对象属性查询
func (c *Client) ApiDOQuery(ctx *gin.Context, request *idl.ApiDOQueryRequest) (*idl.ApiDOQueryResponse, error) {
	return nil, nil
}

// 数据对象权属查询
func (c *Client) ApiDOAuthQuery(ctx *gin.Context, request *idl.ApiDOQueryRequest) (*idl.ApiDOQueryResponse, error) {
	return nil, nil
}
