package exception

import (
	"errors"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"io"
	"net"
	"os"
	"syscall"
)

const (
	OK                     = iota
	DIRPC_UNKNOW_EXCEPTION = 1
)

const (
	// DIRPC transport/conn exception 10xxx
	DIRPC_CONNECTION_GET_FAIL = 10001 + iota
	DIRPC_RPC_ALLADDR_FAIL
	DIRPC_CALL_CONN_NOT_OPEN
	DIRPC_CALL_WRITE_ERROR
	DIRPC_CALL_READ_ERROR
	DIRPC_TRANSPORT_ERROR
	DIRPC_TRANSPORT_CREATE_ERROR
	DIRPC_SOCKET_BAD_ADDR
	DIRPC_SOCKET_NOT_OPEN
	DIRPC_SOCKET_ALREAD_OPEN
	DIRPC_SOCKET_TIME_OUT
	DIRPC_SOCKET_EOF
	DIRPC_PROTOCOL_FACTORY_CREATE_ERROR
	DIRPC_GET_TRANSPORT_ERROR
)

const (
	// DIRPC client exception 11xxx
	DIRPC_SETUP_ERROR = 11001 + iota
	DIRPC_CONFIG_GET_ERROR
	DIRPC_THRIFT_CLIENT_CREATE_ERROR
	DIRPC_HTTP_CLIENT_CREATE_ERROR
	DIRPC_INVALID_MODULE_NAME
	DIRPC_INVALID_SERVICE_CONFIG
	DIRPC_HTTP_REQUEST_FAIL
	DIRPC_HTTP_CREATE_REQUEST_ERROR
	DIRPC_THRIFT_NEW_TRANSPORT_ERROR
	DIRPC_CONFIG_UPDATE_ERROR
	DIRPC_GRPC_INVOKE_FAIL
	DIRPC_FUSING_ERROR // 熔断错误码
	DIRPC_HTTP_GETCAFILE_ERROR
	DIRPC_HTTP_SCHEME_ERROR
	DIRPC_INTERCEPTOR_ERROR // 拦截器错误码
	DIRPC_INVOKE_MOCK_ERROR
	DIRPC_MESH_REQUEST_ERROR
	DIRPC_MESH_RESPONSE_ERROR
	DIRPC_MESH_CONN_ERROR
)

const (
	// DIRPC service exception 12xxx
	DIRPC_INVLAID_SU = 12001 + iota
	DIRPC_SERVICE_REGISTER_ERROR
	DIRPC_SERVICE_STOPED
	DIRPC_SERVICE_GET_FAIL
	DIRPC_SERVICE_NEW_ERROR
	DIRPC_SERVICE_GETCONN_ERROR
)

const (
	// DIRPC naming exception 13xxx
	DIRPC_NAMING_REGISTER_ERROR = 13001 + iota
	DIRPC_NAMING_RESOLVE_ERROR
)

const (
	DIRPC_BALANCE_SELECT_ERROR = 14001 + iota
	DIRPC_INVALID_HTTP_REQ
	DIRPC_MESH_PARSEHIJACKADDR_ERROR
	DIRPC_MESHDEGRADE_ERROR
)

const (
	DIRPC_ADDRMANAGER_NEW_ERROR = 15001 + iota
	DIRPC_ADDRMANAGER_NOTIFY_ERROR
	DIRPC_ADDRMANAGER_UPDATEADDR_ERROR
	DIRPC_ADDRMANAGER_INIT_BALANCER_ERROR
)

const (
	DIRPC_POOLMANAGER_NEW_ERROR = 16001 + iota
	DIRPC_POOL_CLOSED_ERROR
	DIRPC_POOL_OPENED_ERROR
	DIRPC_POOL_DIAL_ERROR
)

const (
	// DIRPC internal exception 19xxx
	DIRPC_INTERNAL_VOTE_ERROR = 19001 + iota
	DIRPC_INTERNAL_UPDATENAMING_ERROR
	DIRPC_INTERNAL_TRIGGER_MINUSABLE
	DIRPC_INTERNAL_CONNPOOL_ERROR
	DIRPC_INTERNAL_PANIC
	DIRPC_NOT_SETUP_ERROR

	DIRPC_INTERNAL_DEBUG = 19999
)

const (
	DIRPC_SEND_METRIC_ERROR = 20001 + iota
)

const (
	DIRPC_HEALTH_CHECK_ERROR = 21001 + iota
)

const (
	DIRPC_CIRCUIT_ERROR = 22001 + iota
)

type timeoutable interface {
	Timeout() bool
}

type TDirpcException interface {
	thrift.TException
	TypeId() int
	DirpcType()
}

type tDirpcException struct {
	typeId  int
	message string
}

func (e *tDirpcException) TExceptionType() thrift.TExceptionType {
	//TODO implement me
	panic("implement me")
}

func (e *tDirpcException) TypeId() int {
	return e.typeId
}

func (e *tDirpcException) Error() string {
	return e.message
}

func (e *tDirpcException) DirpcType() {
}

func IsCaredErrorForMesh(err error) bool {
	if errors.Is(err, syscall.ECONNREFUSED) || errors.Is(err, syscall.ENOENT) || errors.Is(err, syscall.ETIMEDOUT) || errors.Is(err, syscall.ENOTSOCK) {
		return true
	}
	return false
}

func IsCaredNetError(err error) bool {
	netErr, ok := err.(net.Error)
	if !ok {
		return false
	}
	opErr, ok := netErr.(*net.OpError)
	if !ok {
		return false
	}
	switch t := opErr.Err.(type) {
	case *os.SyscallError:
		if errno, ok := t.Err.(syscall.Errno); ok {
			switch errno {
			case syscall.EPIPE:
				return true
			case syscall.ECONNRESET:
				return true
			}
		}
	}

	return false
}

func NewDirpcException(errType int, message string) TDirpcException {
	return &tDirpcException{errType, message}
}

func NewDirpcExceptionf(errType int, format string, args ...interface{}) TDirpcException {
	return NewDirpcException(errType, fmt.Sprintf(format, args...))
}

func WithMessagef(err error, format string, args ...interface{}) error {
	return WithMessage(err, fmt.Sprintf(format, args...))
}

// WithMessage 累加错误信息, 并返回新的error对象. 如果err的类型不是TTransportException, TProtocolException, TApplicationException
//, 则创建对应的TDirpcException对象
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}

	message = message + " ## " + err.Error()

	switch v := err.(type) {
	case TDirpcException:
		return NewDirpcException(v.TypeId(), message)
	case timeoutable:
		if v.Timeout() {
			return NewDirpcException(DIRPC_SOCKET_TIME_OUT, message)
		}
	}

	if err == io.EOF {
		return NewDirpcException(DIRPC_SOCKET_EOF, message)
	}

	return NewDirpcException(DIRPC_UNKNOW_EXCEPTION, message)
}
