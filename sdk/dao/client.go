package dao

import "go.uber.org/zap"

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
