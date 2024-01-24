package dip_dns

import (
	"strings"
	"time"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/miekg/dns"
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

// DNS版本
func (c *Client) GetPublicKey(identifier string) string {
	if identifier == "" {
		return ""
	}
	identifier = dns.Fqdn(identifier)

	c.Logger.Info("[GetPublicKey] receive query user key", zap.String("identifier", identifier))

	dnsClient := new(dns.Client)
	dnsClient.Timeout = 15000 * time.Millisecond

	qtype := dns.TypeCERT
	req := new(dns.Msg)
	req.Zero = true
	req.SetQuestion(identifier, qtype)
	req.SetEdns0(4096, false)
	msg, _, err := dnsClient.Exchange(req, c.DnsHost)
	if err != nil {
		c.Logger.Error("[GetPublicKey] dns.Exchange error", zap.String("error", err.Error()))
		return ""
	}
	if msg == nil {
		// 失败：DNS解析无结果
		c.Logger.Error("[GetPublicKey] failed to handle the request", zap.String("req", converter.ToString(req)))
		return ""
	}

	if len(msg.Answer) == 0 {
		c.Logger.Error("[GetPublicKey] failed to find the public-key", zap.String("identity_identifier", identifier))
		return ""
	}
	a := msg.Answer[0]

	tmp := strings.TrimPrefix(a.String(), a.Header().String())
	slice := strings.Split(tmp, " ")
	if len(slice) != 4 {
		// 失败：无法从结果RR中获取用户公钥
		c.Logger.Error("[GetPublicKey] failed to split the public-key from the answer RR", zap.String("answer", tmp))
		return ""
	}

	c.Logger.Info("[GetPublicKey] receive query user key response", zap.String("identifier", identifier), zap.String("response", slice[3]))
	return slice[3]
}

func (c *Client) GetDataOwner(identifier string) string {
	if identifier == "" {
		return ""
	}
	identifier = dns.Fqdn(identifier)
	c.Logger.Info("[GetDataOwner] receive query data", zap.String("identifier", identifier))

	qtype := dns.TypeRP

	req := new(dns.Msg)
	req.SetQuestion(identifier, qtype)
	req.SetEdns0(4096, false)

	msg, err := dns.Exchange(req, c.DnsHost)
	if err != nil {
		c.Logger.Error("[GetDataOwner] dns.Exchange error", zap.String("error", err.Error()))
		return ""
	}

	if len(msg.Answer) == 0 {
		c.Logger.Error("[GetDataOwner] failed to find the ownerID", zap.String("identity_identifier", identifier))
		return ""
	}
	// a := msg.Answer[0]
	var owners string
	for index, a := range msg.Answer {
		if index == 0 {
			owners = ""
		} else {
			owners += ","
		}

		tmp := strings.TrimPrefix(a.String(), a.Header().String())
		slice := strings.Split(tmp, " ")
		if len(slice) != 2 {
			c.Logger.Error("[GetDataOwner] failed to split the ownerID from the answer RR", zap.String("answer", tmp))
			return ""
		}

		tmp = strings.Trim(slice[0], "\"")

		tmp2 := strings.Split(tmp, "data")
		if len(tmp2) > 2 {
			c.Logger.Error("[GetDataOwner] failed to split the ownerID from the whole name", zap.String("answer", tmp))
			return ""
		}
		c.Logger.Info("[GetDataOwner] receive query user key response", zap.String("identifier", identifier), zap.String("response", tmp2[0]))
		owners += tmp2[0]
	}

	return owners
}
