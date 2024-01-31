package config

import (
	"encoding/json"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/exception"
	"net/http"
	"strings"
	"sync"
)

const recvTimeoutMsecMagicNum int = 188
const (
	defaultServiceConfigJson = ` {
    "disableLog":           false,
	"logMaxHttpBodyLen": 		0,
	"logMaxHttpRespLen": 		0,
	"disableMetric":        false,
    "connectTimeoutMsec":   100,
    "sendTimeoutMsec":      200,
    "recvTimeoutMsec":      188,
    "timeoutMsec":          400,

    "retry":                0,
    "healthyThreshold":     5,
    "maxCooldownTime":      60,
    "minHealthyRatio":      0.67,
	
    "thriftLog": 			false,
	"sdkSupportThriftLog":	false,
	"smartRetry": 			false,
	
	"keepAlives":           false,
	"maxIdleConn":          5,
	"idleConnTimeout":      3,

	"fusing":false,
	"mock":false,

    "protocol":             "TBinaryProtocol",
    "transport":            "TFramedTransport",
	"scheme":               "http",
	"cacert":				"",
	"headers":              [],
	"signType":             "",
	"wardenCustom":         false,
	"wardenAccessKey":      "",
	"wardenSecretKey":      "",
	"hashRingLen":  		1024
}`
)

const (
	TBinaryProtocol  = "TBinaryProtocol"
	TCompactProtocol = "TCompactProtocol"

	TSocket          = "TSocket"
	TFramedTransport = "TFramedTransport"
)

var (
	defaultServiceConfig = &ServiceConfig{}
	defaultDirpcConfig   = NewDirpcConfig()
)

func init() {
	// 初始化默认配置
	if err := json.Unmarshal([]byte(defaultServiceConfigJson), defaultServiceConfig); err != nil {
		panic(err)
	}
}

func GetLidc() string {
	return defaultDirpcConfig.GetLidc()
}

func GetSu() string {
	return defaultDirpcConfig.GetSu()
}

func GetUsn() string {
	return defaultDirpcConfig.GetUsn()
}

func DisableTrace() {
	defaultDirpcConfig.DisableTrace()
}

func SetMeshDegradeErrRate(MeshDegradeErrRate float64) {
	defaultDirpcConfig.SetMeshDegradeErrRate(MeshDegradeErrRate)
}

func GetMeshDegradeErrRate() float64 {
	return defaultDirpcConfig.GetMeshDegradeErrRate()
}

func SetMeshDegrade(ifDegrade bool) {
	defaultDirpcConfig.SetMeshDegrade(ifDegrade)
}

func GetMeshDegrade() bool {
	return defaultDirpcConfig.GetMeshDegrade()
}

func IsTraceDisabled() bool {
	return defaultDirpcConfig.IsTraceDisabled()
}

type BaseConfig struct {
	// log switch
	DisableLog        *bool `json:"disableLog,omitempty"`
	LogMaxHttpBodyLen *int  `json:"logMaxHttpBodyLen,omitempty"`
	LogMaxHttpRespLen *int  `json:"logMaxHttpRespLen,omitempty"`
	DisableMetric     *bool `json:"disableMetric,omitempty"`

	// timeout
	ConnectTimeoutMsec *int `json:"connectTimeoutMsec,omitempty" validate:"min=1"`
	SendTimeoutMsec    *int `json:"sendTimeoutMsec,omitempty" validate:"min=1"`
	RecvTimeoutMsec    *int `json:"recvTimeoutMsec,omitempty" validate:"min=1"`
	TimeoutMsec        *int `json:"timeoutMsec,omitempty" validate:"min=1"`

	// balancer && fault tolerance
	Retry            *int     `json:"retry,omitempty" validate:"min=0,max=4"`
	MinHealthyRatio  *float64 `json:"minHealthyRatio,omitempty" validate:"min=0,max=1"`
	HealthyThreshold *int     `json:"healthyThreshold,omitempty" validate:"min=1"`
	MaxCooldownTime  *int     `json:"maxCooldownTime,omitempty" validate:"min=1"`

	// keepalive, only support thrift
	KeepAlives      *bool `json:"keepAlives,omitempty"`
	MaxIdleConn     *int  `json:"maxIdleConn,omitempty" validate:"min=1"`
	IdleConnTimeout *int  `json:"idleConnTimeout,omitempty" validate:"min=1"`

	// fuse switch
	Fusing *bool `json:"fusing,omitempty"`

	// mock调用
	Mock *bool `json:"mock,omitempty"`

	// thrift specific
	Transport *string `json:"transport,omitempty"`
	Protocol  *string `json:"protocol,omitempty"`

	ThriftLog           *bool `json:"thriftLog,omitempty"`
	SDKSupportThriftLog *bool `json:"sdkSupportThriftLog,omitempty"`
	// used for dirpc persistent connection closed by peer retry
	SmartRetry *bool `json:"smartRetry,omitempty"`
	// http specific
	Scheme  *string  `json:"scheme,omitempty"`
	CACert  *string  `json:"cacert,omitempty"`
	Headers []string `json:"headers,omitempty"`

	// sign鉴权配置，当前只支持warden
	SignType        *string `json:"signType,omitempty"`
	WardenCustom    *bool   `json:"wardenCustom,omitempty"`
	WardenAccessKey *string `json:"wardenAccessKey,omitempty"`
	WardenSecretKey *string `json:"wardenSecretKey,omitempty"`

	// load balance type
	LbType      *string `json:"lbType,omitempty"`
	HashRingLen *int    `json:"hashRingLen,omitempty" validate:"min=256"`
}

// ServiceConfig表示单个服务的配置信息
type ServiceConfig struct {
	// 配置结构体
	BaseConfig
	ActionList map[string]BaseConfig `json:"actionList,omitempty"`

	// ServiceConfig 标识
	Usn    string
	Disf   string
	rwlock sync.RWMutex
}

func (p *ServiceConfig) GetLogMaxHttpBodyLen() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.LogMaxHttpBodyLen
}
func (p *ServiceConfig) GetLogMaxHttpRespLen() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.LogMaxHttpRespLen
}
func (p *ServiceConfig) GetHealthyThreshold() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.HealthyThreshold
}

func (p *ServiceConfig) GetHashRingLen() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.HashRingLen
}

func (p *ServiceConfig) GetNewSDKSupportThriftLog() bool {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.SDKSupportThriftLog
}

func (p *ServiceConfig) GetSmartRetry() bool {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.SmartRetry
}

func (p *ServiceConfig) GetMaxCooldownTime() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.MaxCooldownTime
}

func (p *ServiceConfig) GetKeepAlives() bool {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.KeepAlives
}

func (p *ServiceConfig) GetMaxIdleConn() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.MaxIdleConn
}

func (p *ServiceConfig) GetIdleConnTimeout() int {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.IdleConnTimeout
}

func (p *ServiceConfig) GetTransport() string {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.Transport
}
func (p *ServiceConfig) GetProtocol() string {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	return *p.Protocol
}

func (p *ServiceConfig) GetHeaders() http.Header {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()

	result := http.Header{}
	h := p.Headers
	for _, v := range h {
		hSplit := strings.Split(v, ":")
		if len(hSplit) != 2 {
			continue
		}
		result.Add(hSplit[0], hSplit[1])
	}
	return result
}

func NewDirpcConfig() *DirpcConfig {
	dirpcConfig := &DirpcConfig{}
	dirpcConfig.runtimeConfig = make(map[string]*ServiceConfig)
	dirpcConfig.consumerConfig = make(map[string]*ServiceConfig)
	dirpcConfig.providerConfig = make(map[string]*ServiceConfig)
	dirpcConfig.finalConfig = make(map[string]*ServiceConfig)
	dirpcConfig.defaultConfig = defaultServiceConfig

	return dirpcConfig
}

type DirpcConfig struct {
	mu sync.RWMutex

	// 部署信息
	deploySu   string
	deployUsn  string
	deployLidc string

	// 是否关闭trace功能
	disableTrace       bool
	MeshDegradeErrRate float64
	MeshDegrade        bool
	// 优先级由上到下，依次降低
	runtimeConfig      map[string]*ServiceConfig // provider运行时配置
	consumerConfig     map[string]*ServiceConfig // consumer差异化配置
	providerConfig     map[string]*ServiceConfig // provider配置
	consumerMetaConfig *ServiceConfig            // consumer默认配置
	defaultConfig      *ServiceConfig            // 全局默认配置，兜底用

	// 最终配置：provider usn=》*ServiceConfig
	finalConfig map[string]*ServiceConfig
}

func (p *DirpcConfig) GetSu() string {
	return p.deploySu
}

func (p *DirpcConfig) GetUsn() string {
	return p.deployUsn
}

func (p *DirpcConfig) GetLidc() string {
	return p.deployLidc
}

func (p *DirpcConfig) DisableTrace() {
	p.disableTrace = true
}

func (p *DirpcConfig) IsTraceDisabled() bool {
	return p.disableTrace
}

func (p *DirpcConfig) SetMeshDegradeErrRate(MeshDegradeErrRate float64) {
	p.MeshDegradeErrRate = MeshDegradeErrRate
}

func (p *DirpcConfig) GetMeshDegradeErrRate() float64 {
	return p.MeshDegradeErrRate
}

func (p *DirpcConfig) SetMeshDegrade(ifDegrade bool) {
	p.MeshDegrade = ifDegrade
}

func (p *DirpcConfig) GetMeshDegrade() bool {
	return p.MeshDegrade
}

func (p *DirpcConfig) tryGetServiceConfig(usn string) (*ServiceConfig, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	sc, ok := p.finalConfig[usn]
	if !ok {
		return nil, exception.NewDirpcExceptionf(exception.DIRPC_CONFIG_GET_ERROR, "serviceConfig not registered, usn=%s", usn)
	}

	return sc, nil
}
