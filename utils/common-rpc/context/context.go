package context

import (
	"context"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/common"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/config"
	"net/http"
)

type Context struct {
	CallInfo           *common.CallInfo
	HttpHeader         http.Header
	conf               *config.ServiceConfig
	rpcCluster         string
	connectTimeoutMsec int
	sendTimeoutMsec    int
	recvTimeoutMsec    int
	//用于上报没有设置trace的metric
	NoTraceFlag bool

	//用于hashringlb负载.
	LbHashKey string
	//规则熔断属性
	fusingProperties map[string]string
}

func NewContext() *Context {
	dctx := &Context{NoTraceFlag: false}
	dctx.HttpHeader = make(map[string][]string)
	dctx.CallInfo = common.NewCallInfo()
	return dctx
}

func (ctx *Context) SetHintInfo(hintcode string, hintcontent string) {
	ctx.CallInfo.HintCode = hintcode
	ctx.CallInfo.HintContent = hintcontent
}

func (ctx *Context) SetTraceInfo(traceid string, spanid string) {
	ctx.CallInfo.TraceId = traceid
	ctx.CallInfo.SpanId = spanid
}

func (ctx *Context) SetConfig(conf *config.ServiceConfig) {
	ctx.conf = conf
}

func (ctx *Context) GetConfig() *config.ServiceConfig {
	return ctx.conf
}

func (ctx *Context) SetRpcCluster(cluster string) {
	ctx.rpcCluster = cluster
}

func (ctx *Context) GetRpcCluster() string {
	return ctx.rpcCluster
}

func (ctx *Context) SetConnectTimeoutMsec(connectTimeoutMsec int) {
	ctx.connectTimeoutMsec = connectTimeoutMsec
}

func (ctx *Context) GetConnectTimeoutMsec() int {
	return ctx.connectTimeoutMsec
}

func (ctx *Context) SetSendTimeoutMsec(sendTimeoutMsec int) {
	ctx.sendTimeoutMsec = sendTimeoutMsec
}

func (ctx *Context) GetSendTimeoutMsec() int {
	return ctx.sendTimeoutMsec
}

func (ctx *Context) SetRecvTimeoutMsec(recvTimeoutMsec int) {
	ctx.recvTimeoutMsec = recvTimeoutMsec
}

func (ctx *Context) GetRecvTimeoutMsec() int {
	return ctx.recvTimeoutMsec
}

func (ctx *Context) SetLbHashKey(key string) {
	ctx.LbHashKey = key
}

func (ctx *Context) GetLbHashKey() string {
	return ctx.LbHashKey
}

func (ctx *Context) SetFusingProperties(fusingProperties map[string]string) {
	ctx.fusingProperties = fusingProperties
}

func (ctx *Context) GetFusingProperties() map[string]string {
	return ctx.fusingProperties
}

func GetRpcCluster(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if cluster, ok := ctx.Value(contextRpcClusterKey{}).(string); ok {
		return cluster
	}
	return ""
}

func GetConnectTimeout(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	if connectTimeout, ok := ctx.Value(contextConnectTimeoutMsecKey{}).(int); ok {
		return connectTimeout
	}
	return 0
}

func GetSendTimeout(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	if sendTimeout, ok := ctx.Value(contextSendTimeoutMsecKey{}).(int); ok {
		return sendTimeout
	}
	return 0
}

func GetRecvTimeout(ctx context.Context) int {
	if ctx == nil {
		return 0
	}
	if recvTimeout, ok := ctx.Value(contextRecvTimeoutMsecKey{}).(int); ok {
		return recvTimeout
	}
	return 0
}

func GetHttpHeader(ctx context.Context) http.Header {
	if userHeaders, ok := ctx.Value(contextHttpHeadersKey{}).(http.Header); ok {
		return userHeaders
	}
	return http.Header{}
}

func GetLbHashKey(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if lbHashKey, ok := ctx.Value(contextLbHashKey{}).(string); ok {
		return lbHashKey
	}
	return ""
}

func GetFusingProperties(ctx context.Context) map[string]string {
	if ctx == nil {
		return make(map[string]string, 0)
	}
	if fusingProperties, ok := ctx.Value(contextFusingPropertiesKey{}).(map[string]string); ok {
		return fusingProperties
	}
	return make(map[string]string, 0)
}
