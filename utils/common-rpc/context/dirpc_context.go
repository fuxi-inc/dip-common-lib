package context

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/config"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/exception"
	"github.com/gin-gonic/gin"
)

type contextHttpHeadersKey struct{}
type contextHttpQueryStringKey struct{}
type contextRpcClusterKey struct{}
type contextConnectTimeoutMsecKey struct{}
type contextSendTimeoutMsecKey struct{}
type contextRecvTimeoutMsecKey struct{}
type contextTimeoutMsecKey struct{}
type contextLbHashKey struct{}
type contextFusingPropertiesKey struct{}
type contextCACert struct{}
type contextKeepAlive struct{}
type contextScheme struct{}
type contextRetryNum struct{}
type MapForRestFulUrl struct{}
type UriForCallee struct{}

/*
func SetHttpClientRestfulParam(parent gin.Context, restfulParma map[string]string) gin.Context {
	return context.WithValue(parent, MapForRestFulUrl{}, restfulParma)
}

func SetHttpHeader(parent gin.Context, headers http.Header) gin.Context {
	return context.WithValue(parent, contextHttpHeadersKey{}, headers)
}

func SetHttpUrlValues(parent gin.Context, values url.Values) gin.Context {
	return context.WithValue(parent, contextHttpQueryStringKey{}, values)
}

func SetRpcCluster(parent gin.Context, cluster string) gin.Context {
	return context.WithValue(parent, contextRpcClusterKey{}, cluster)
}

func SetConnectTimeoutMsec(parent gin.Context, connectTimeoutMsec int) gin.Context {
	return context.WithValue(parent, contextConnectTimeoutMsecKey{}, connectTimeoutMsec)
}

func SetSendTimeoutMsec(parent gin.Context, sendTimeoutMsec int) gin.Context {
	return context.WithValue(parent, contextSendTimeoutMsecKey{}, sendTimeoutMsec)
}

func SetRecvTimeoutMsec(parent gin.Context, recvTimeoutMsec int) gin.Context {
	return context.WithValue(parent, contextRecvTimeoutMsecKey{}, recvTimeoutMsec)
}

func SetTimeoutMsec(parent gin.Context, timeoutMsec int) gin.Context {
	return context.WithValue(parent, contextTimeoutMsecKey{}, timeoutMsec)
}

func SetLbHashKey(parent gin.Context, lbHashKey string) gin.Context {
	return context.WithValue(parent, contextLbHashKey{}, lbHashKey)
}

func SetFusingProperties(parent gin.Context, properties map[string]string) gin.Context {
	return context.WithValue(parent, contextFusingPropertiesKey{}, properties)
}

func SetCACert(parent gin.Context, CACert string) gin.Context {
	return context.WithValue(parent, contextCACert{}, CACert)
}

func SetKeepAlive(parent gin.Context, KeepAlive bool) gin.Context {
	return context.WithValue(parent, contextKeepAlive{}, KeepAlive)
}

func SetScheme(parent gin.Context, Scheme string) gin.Context {
	return context.WithValue(parent, contextScheme{}, Scheme)
}

func SetRetryNum(parent gin.Context, RetryNum int) gin.Context {
	return context.WithValue(parent, contextRetryNum{}, RetryNum)
}
*/

type CallInfo struct {
	// 其它
	RetryFlag int // 当前是第几次重试
	// 延迟信息
	Latency   time.Duration
	Footprint *Footprint
	// 错误信息
	ErrNo  int
	ErrMsg string
	// Host
	Host string
	Ip   string
	// config

	KeepAlive       bool
	Timeout         int
	ConnectTimeout  int
	Retry           int
	Chaos           int
	FusingType      string // 熔断类型 取值：auto、manual
	MeshDegradeType string

	/***** http *******/
	// 入参、出参
	Url  string
	Path string
	Body string
	Resp string
	// Header
	Headers http.Header

	//mesh真实访问ip
	MeshRealIp string
}

func NewCallInfo() *CallInfo {
	return &CallInfo{
		Footprint: NewFootprint(),
	}
}

type DirpcContext struct {
	gin.Context

	CallInfo *CallInfo
	Conf     *config.ServiceConfig
}

func NewDirpcContext(ctx gin.Context) *DirpcContext {
	return &DirpcContext{
		Context:  ctx,
		CallInfo: NewCallInfo(),
	}
}

func (p *DirpcContext) SetErr(err error) *DirpcContext {
	if p != nil && p.CallInfo != nil {
		if err == nil {
			p.CallInfo.ErrNo = exception.OK
			p.CallInfo.ErrMsg = "ok"
		} else if e, ok := err.(exception.TDirpcException); ok {
			p.CallInfo.ErrNo = e.TypeId()
			p.CallInfo.ErrMsg = err.Error()
		} else {
			p.CallInfo.ErrNo = exception.DIRPC_UNKNOW_EXCEPTION
			p.CallInfo.ErrMsg = err.Error()
		}
	}

	return p
}

func (p *DirpcContext) SetLatency(latency time.Duration) *DirpcContext {
	if p != nil && p.CallInfo != nil {
		p.CallInfo.Latency = latency
	}

	return p
}

func (p *DirpcContext) GetLogEntry() string {
	buf := make([]byte, 0, 1024)
	log := bytes.NewBuffer(buf)

	log.WriteString("||latency=")
	log.WriteString(strconv.FormatInt(int64(p.CallInfo.Latency/time.Millisecond), 10))

	// Host
	if p.CallInfo.Host != "" {
		log.WriteString("||host=")
		log.WriteString(p.CallInfo.Host)
	}

	// errinfo
	log.WriteString("||errno=")
	log.WriteString(strconv.Itoa(p.CallInfo.ErrNo))

	log.WriteString("||errmsg=")
	log.WriteString(p.CallInfo.ErrMsg)

	// footprint
	if footprint := p.CallInfo.Footprint; footprint != nil {
		if footstr := footprint.String(); footstr != "" {
			log.WriteString("||footprint=")
			log.WriteString(footstr)
		}
	}
	//降级类型
	if p.CallInfo.MeshDegradeType != "" {
		log.WriteString("||meshDegradeType=")
		log.WriteString(p.CallInfo.MeshDegradeType)
	}
	if p.CallInfo.MeshRealIp != "" {
		log.WriteString("||realIp=")
		log.WriteString(p.CallInfo.MeshRealIp)
	}

	// config
	log.WriteString("||keepAlive=")
	log.WriteString(strconv.FormatBool(p.CallInfo.KeepAlive))
	log.WriteString("||protocol=http")
	log.WriteString("||timeout=")
	log.WriteString(strconv.Itoa(p.CallInfo.Timeout))
	log.WriteString("||connectTimeout=")
	log.WriteString(strconv.Itoa(p.CallInfo.ConnectTimeout))
	log.WriteString("||retry=")
	log.WriteString(strconv.Itoa(p.CallInfo.Retry))
	log.WriteString("||retryFlag=")
	log.WriteString(strconv.Itoa(p.CallInfo.RetryFlag))
	log.WriteString("||chaos=")
	log.WriteString(strconv.Itoa(p.CallInfo.Chaos))

	// req & resp
	if p.CallInfo.Url != "" {
		log.WriteString("||url=")
		log.WriteString(p.CallInfo.Url)
	}
	if p.CallInfo.Body != "" {
		maxHttpBodyLen := p.Conf.GetLogMaxHttpBodyLen()
		bodyLen := len(p.CallInfo.Body)
		if maxHttpBodyLen > 0 && bodyLen > maxHttpBodyLen {
			log.WriteString("||body=")
			log.WriteString(fmt.Sprintf("body len is over logMaxHttpBodyLen. bodyLen=%d,logMaxHttpBodyLen=%d", bodyLen, maxHttpBodyLen))
		} else {
			retStr := strings.Replace(p.CallInfo.Body, "\r\n", " ", -1)
			log.WriteString("||body=")
			log.WriteString(strings.Replace(retStr, "\n", " ", -1))
		}
	}
	if p.CallInfo.Resp != "" {
		maxHttpRespLen := p.Conf.GetLogMaxHttpRespLen()
		respLen := len(p.CallInfo.Resp)
		if maxHttpRespLen > 0 && respLen > maxHttpRespLen {
			log.WriteString("||response=")
			log.WriteString(fmt.Sprintf("resp len is over logMaxHttpRespLen. respLen=%d,logMaxHttpRespLen=%d", respLen, maxHttpRespLen))
		} else {
			retStr := strings.Replace(p.CallInfo.Resp, "\r\n", " ", -1)
			log.WriteString("||response=")
			log.WriteString(strings.Replace(retStr, "\n", " ", -1))
		}
	}

	return log.String()
}

func (p *DirpcContext) GetCustomHttpHeader() http.Header {
	if userHeaders, ok := p.Context.Value(contextHttpHeadersKey{}).(http.Header); ok {
		return userHeaders
	}
	return http.Header{}
}

func (p *DirpcContext) GetUrlValues() url.Values {
	if values, ok := p.Context.Value(contextHttpQueryStringKey{}).(url.Values); ok {
		return values
	}
	return nil
}

func (p *DirpcContext) GetRpcCluster() string {
	if cluster, ok := p.Context.Value(contextRpcClusterKey{}).(string); ok {
		return cluster
	}
	return ""
}

func (p *DirpcContext) GetConnectTimeoutMsec() int {
	if connectTimeout, ok := p.Context.Value(contextConnectTimeoutMsecKey{}).(int); ok {
		return connectTimeout
	}
	return 0
}

func (p *DirpcContext) GetSendTimeoutMsec() int {
	if sendTimeout, ok := p.Context.Value(contextSendTimeoutMsecKey{}).(int); ok {
		return sendTimeout
	}
	return 0
}

func (p *DirpcContext) GetRecvTimeoutMsec() int {
	if recvTimeout, ok := p.Context.Value(contextRecvTimeoutMsecKey{}).(int); ok {
		return recvTimeout
	}
	return 0
}

func (p *DirpcContext) GetTimeoutMsec() int {
	if timeout, ok := p.Context.Value(contextTimeoutMsecKey{}).(int); ok {
		return timeout
	}
	return 0
}

func (p *DirpcContext) GetFusingProperties() map[string]string {
	if fusingProperties, ok := p.Context.Value(contextFusingPropertiesKey{}).(map[string]string); ok {
		return fusingProperties
	}
	return nil
}

func (p *DirpcContext) GetCACert() string {
	if CACert, ok := p.Context.Value(contextCACert{}).(string); ok {
		return CACert
	}
	return ""
}

func (p *DirpcContext) GetKeepAlive() *bool {
	if KeepAlive, ok := p.Context.Value(contextKeepAlive{}).(bool); ok {
		return &KeepAlive
	}
	return nil
}

func (p *DirpcContext) GetScheme() string {
	if Scheme, ok := p.Context.Value(contextScheme{}).(string); ok {
		return Scheme
	}
	return ""
}

func (p *DirpcContext) GetRetryNum() *int {
	if RetryNum, ok := p.Context.Value(contextRetryNum{}).(int); ok {
		return &RetryNum
	}
	return nil
}

func (p *DirpcContext) GetRestfulUrl() string {
	if RestfulUrl, ok := p.Context.Value(UriForCallee{}).(string); ok {
		return RestfulUrl
	}
	return ""
}
