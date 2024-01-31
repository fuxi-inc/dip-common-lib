package utils

import (
	"google.golang.org/grpc/metadata"
	"net/http"
)

func CopyHttpHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}

func MergeHttpHeader(dst http.Header, src http.Header) {
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			dst[k] = v
		}
	}
}

func MergeHttpHeaderOverwrite(dst http.Header, src http.Header) {
	for k, _ := range src {
		dst.Set(k, src.Get(k))
	}
}

func MergeMD(dst metadata.MD, src http.Header) {
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			dst[k] = v
		}
	}
}
