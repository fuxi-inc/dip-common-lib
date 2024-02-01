package common

import (
	"bytes"
	"strconv"
	"time"
)

const RPC_BEGIN = "begin"
const RPC_END = "end"

const FUSING_AUTO = "auto"
const FUSING_MANUAL = "manual"

//降级原因
const CONNECT_ERR = "connect"
const RESPFLAG_ERR = "respFlag"
const STATISTIC_ERR = "errRatio"

type CallInfo struct {
	Type    string `key:"type"`
	Process string `key:"process"`

	CalleeLidc string `key:"callee_lidc"`

	// caller&callee
	CallerSu     string `key:"caller_su"`
	CalleeSu     string `key:"callee_su"`
	CalleeIp     string `key:"callee_ip"`
	CallerMethod string `key:"caller_method"`
	CalleeMethod string `key:"function_name"`

	// trace
	TraceId string `key:"traceid"`
	SpanId  string `key:"spanid"`
	CSpanId string `key:"cspanid"`

	// hint info
	HintCode    string `key:"hintCode"`
	HintContent string `key:"hintContent"`
	// press flow
	IsPressTraffic bool `key:"pressTraffic"`
	// total time for flush, including loadbalance
	Latency time.Duration `key:"proc_time"`
	Start   time.Time     `key:"start"` // rpc 开始时间

	// footprint
	Footprint *Footprint `key:"footprint"`

	// errinfo
	ErrNo  int    `key:"errno"`
	ErrMsg string `key:"errmsg"`

	// grpc
	Body string `key:"body"`
	Resp string `key:"resp"`

	// timeout
	Timeout        int `key:"timeout"`
	ConnectTimeout int `key:"connectTimeout"`
	SendTimeout    int `key:"sendTimeout"`
	RecvTimeout    int `key:"recvTimeout"`

	// retry
	Retry int `key:"retry"`

	//标准链路透传header
	StandardTraceHeader string `key:"standardTraceHeader"`
	RegionKeyName       string `key:"regionKeyName"`
	RegionKeyValue      string `key:"regionKeyValue"`

	// 其它
	RetryFlag       int                    `key:"retryFlag"`       // 当前是第几次重试
	Chaos           int                    `key:"chaos"`           // 是否被放火
	FusingType      string                 `key:"fusingType"`      // 熔断类型 取值：auto、manual
	MeshDegradeType string                 `key:"meshDegradeType"` // mesh降级类型
	MapForChaos     map[string]interface{} `key:"mapForChaos"`     // 给放火插件里传各类型的值
}

func NewCallInfo() *CallInfo {
	callinfo := &CallInfo{}
	callinfo.Footprint = NewFootprint()
	return callinfo
}

func (ci *CallInfo) GetLogEntry() string {
	buf := make([]byte, 0, 1024)
	log := bytes.NewBuffer(buf)
	// caller&callee
	log.WriteString("||caller=")
	log.WriteString(ci.CallerSu)
	log.WriteString("||callee=")
	log.WriteString(ci.CalleeSu)

	log.WriteString("||caller-func=")
	log.WriteString(ci.CallerMethod)
	log.WriteString("||callee-func=")
	log.WriteString(ci.CalleeMethod)

	// trace
	log.WriteString("||traceid=")
	log.WriteString(ci.TraceId)
	log.WriteString("||spanid=")
	log.WriteString(ci.SpanId)
	log.WriteString("||cspanid=")
	log.WriteString(ci.CSpanId)

	// hint info
	log.WriteString("||hintCode=")
	log.WriteString(ci.HintCode)
	log.WriteString("||hintContent=")
	log.WriteString(ci.HintContent)

	//xregioninfo
	log.WriteString("||xregionkeyname=")
	log.WriteString(ci.RegionKeyName)
	log.WriteString("||xregionkeyvalue=")
	log.WriteString(ci.RegionKeyValue)

	log.WriteString("||latency=")
	log.WriteString(strconv.FormatInt(int64(ci.Latency/time.Millisecond), 10))

	//MeshDegradeType
	if ci.MeshDegradeType != "" {
		log.WriteString("||MeshDegradeType=")
		log.WriteString(ci.MeshDegradeType)
	}

	// errinfo
	log.WriteString("||errno=")
	log.WriteString(strconv.Itoa(ci.ErrNo))
	log.WriteString("||errmsg=")
	log.WriteString(ci.ErrMsg)

	// footprint
	log.WriteString("||footprint=")
	log.WriteString(ci.Footprint.String())

	// grpc
	if ci.Type == "grpc" {
		log.WriteString("||body=")
		log.WriteString(ci.Body)
		log.WriteString("||resp=")
		log.WriteString(ci.Resp)
	}

	// timeout
	if ci.Type == "grpc" {
		log.WriteString("||protocol=grpc")
		log.WriteString("||timeout=")
		log.WriteString(strconv.Itoa(ci.Timeout))
		log.WriteString("||connectTimeout=")
		log.WriteString(strconv.Itoa(ci.ConnectTimeout))
	} else {
		//thrift
		//添加body和resp的打印
		log.WriteString("||body=")
		log.WriteString(ci.Body)
		log.WriteString("||resp=")
		log.WriteString(ci.Resp)

		log.WriteString("||protocol=thrift")
		log.WriteString("||connectTimeout=")
		log.WriteString(strconv.Itoa(ci.ConnectTimeout))
		log.WriteString("||sendTimeout=")
		log.WriteString(strconv.Itoa(ci.SendTimeout))
		log.WriteString("||recvTimeout=")
		log.WriteString(strconv.Itoa(ci.RecvTimeout))
	}

	// retry
	log.WriteString("||retry=")
	log.WriteString(strconv.Itoa(ci.Retry))
	log.WriteString("||retryFlag=")
	log.WriteString(strconv.Itoa(ci.RetryFlag))

	// chaos
	log.WriteString("||chaos=")
	log.WriteString(strconv.Itoa(ci.Chaos))

	if ci.StandardTraceHeader != "" {
		log.WriteString(ci.StandardTraceHeader)
	}
	return log.String()
}
