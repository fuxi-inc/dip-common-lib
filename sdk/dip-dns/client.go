package dip_dns

import "go.uber.org/zap"

/**
*  请求方式类如：
*    response,err := NewClient().
					InitLogger(logger).
					InitDis("http://dis.viv.cn").
					ApiRequestDemo( &idl.ApiDemoRequest{})
*/

type Client struct {
	Logger  *zap.Logger //日志组件
	DnsHost string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InitLogger(logger *zap.Logger) *Client {
	c.Logger = logger
	return c
}

func (c *Client) InitDnsHost(host string) *Client {
	c.DnsHost = host
	return c
}
