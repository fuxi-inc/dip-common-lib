package demo

import (
    "github.com/fuxi-inc/dip-common-lib/sdk/demo/idl"
    "log"
)

type Client struct {
    Logger  *log.Logger //日志组件
    DisHost string
    DaoHost string
}

func NewClient(logger *log.Logger) (*Client, error) {
    return &Client{Logger: logger}, nil
}

func (c *Client) InitDis(disHost string) *Client {
    c.DisHost = disHost
    return c
}

func (c *Client) InitDao(daoHost string) *Client {
    c.DaoHost = daoHost
    return c
}

func (c *Client) ApiRequestDemo(req *idl.ApiDemoRequest) (*idl.ApiDemoResponse, error) {

    return &idl.ApiDemoResponse{}, nil
}
