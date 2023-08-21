package middleware

import (
	"crypto/md5"
	"fmt"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

const (
	TraceIDKeyName      = "Dip-TraceID"
	SpanIDKeyName       = "Dip-SpanID"
	ParentSpanIDKeyName = "Dip-ParentSpanID"
)

func AddTraceId() gin.HandlerFunc {
	return func(context *gin.Context) {
		traceId := context.Request.Header.Get(TraceIDKeyName)
		if traceId == "" {
			h := md5.New()
			h.Write([]byte(uuid.NewString()))
			re := h.Sum(nil)
			traceId = fmt.Sprintf("%x", re)
		}
		context.Set(TraceIDKeyName, traceId)
		context.Header(TraceIDKeyName, traceId)
		context.Next()
	}
}

func AddSpanId() gin.HandlerFunc {
	return func(context *gin.Context) {
		spanId := context.Request.Header.Get(SpanIDKeyName)
		if spanId == "" {
			h := md5.New()
			h.Write([]byte(uuid.NewString()))
			re := h.Sum(nil)
			spanId = fmt.Sprintf("%x", re)
			spanId = spanId[0:16]
		}
		context.Set(SpanIDKeyName, spanId)
		context.Header(SpanIDKeyName, spanId)
		context.Next()
	}
}

func AddParentSpanId() gin.HandlerFunc {
	return func(context *gin.Context) {
		spanId := context.Request.Header.Get(ParentSpanIDKeyName)
		if spanId == "" {
			h := md5.New()
			h.Write([]byte(uuid.NewString()))
			re := h.Sum(nil)
			spanId = fmt.Sprintf("%x", re)
			spanId = spanId[0:16]
		}
		context.Set(ParentSpanIDKeyName, spanId)
		context.Header(ParentSpanIDKeyName, spanId)
		context.Next()
	}
}

func InitRequestHeaders(ctx *gin.Context, req *http.Request) {
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)
	if ctx.Request != nil {
		req.Header.Add(ParentSpanIDKeyName, ctx.Request.Header.Get(ParentSpanIDKeyName))
		req.Header.Add(SpanIDKeyName, ctx.Request.Header.Get(SpanIDKeyName))
		req.Header.Add(TraceIDKeyName, ctx.Request.Header.Get(TraceIDKeyName))
	}
}
