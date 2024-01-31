package common

import (
	"fmt"
	"time"
)

type footprint struct {
	Code    int           `json:"errno"`
	Msg     string        `json:"errmsg"`
	Addr    string        `json:"addr"`
	Latency time.Duration `json:"latency"`
}

type Footprint struct {
	footprints []footprint
}

func NewFootprint() *Footprint {
	return &Footprint{
		footprints: make([]footprint, 0, 2),
	}
}

func (p *Footprint) Step(addr string, code int, msg string, latency time.Duration) {
	f := footprint{
		Code:    code,
		Msg:     msg,
		Addr:    addr,
		Latency: latency,
	}
	p.footprints = append(p.footprints, f)
}

func (p *Footprint) String() string {
	var str string
	for _, fp := range p.footprints {
		str = str + fmt.Sprintf("{addr=%s,latency=%d,errno=%d,errmsg=%s}", fp.Addr, int64(fp.Latency/time.Millisecond), fp.Code, fp.Msg)
	}
	return str
}
