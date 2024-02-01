package dirpc

import (
	"net/http"
)

//Deprecated: Use Context.context instead
type callOptions struct {
	header             http.Header
	connectTimeoutMsec *int
	timeoutMsec        *int
	retry              *int
	keepAlive          *bool
	rpcCluster         *string
	scheme             *string
	unixaddr           *string
	caCert             *string
	fusingProperties   map[string]string
}

//Deprecated: Use context.Context instead
type CallOption func(*callOptions)

//Deprecated: Use SetFusingProperties instead
func WithFusingProperties(properties map[string]string) CallOption {
	return func(opt *callOptions) {
		opt.fusingProperties = properties
	}
}

//Deprecated: Use SetConnectTimeoutMsec instead
func WithConnectTimeoutMsec(tm int) CallOption {
	return func(opt *callOptions) {
		localtm := tm
		opt.connectTimeoutMsec = &localtm
	}
}

//Deprecated: Use SetTimeoutMsec instead
func WithTimeoutMsec(tm int) CallOption {
	return func(opt *callOptions) {
		localtm := tm
		opt.timeoutMsec = &localtm
	}
}

//Deprecated: Use SetRetryNum instead
func WithRetry(retry int) CallOption {
	return func(opt *callOptions) {
		localRetry := retry
		opt.retry = &localRetry
	}
}

//Deprecated: Use SetKeepAlive instead
func WithKeepAlive(keepalive bool) CallOption {
	return func(opt *callOptions) {
		localKA := keepalive
		opt.keepAlive = &localKA
	}
}

//Deprecated: Use SetRpcCluster instead
func WithRpcCluster(cluster string) CallOption {
	return func(opt *callOptions) {
		localCluster := cluster
		opt.rpcCluster = &localCluster
	}
}

//Deprecated: Use SetScheme instead
func WithScheme(scheme string) CallOption {
	return func(opt *callOptions) {
		localScheme := scheme
		opt.scheme = &localScheme
	}
}

//Deprecated: Use SetCACert instead
func WithCACert(cacert string) CallOption {
	return func(opt *callOptions) {
		localCACert := cacert
		opt.caCert = &localCACert
	}
}
