package dirpc

import (
	"bytes"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/exception"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/service"
	"github.com/fuxi-inc/dip-common-lib/utils/common-rpc/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"time"
)

// NewHttpClient 创建HttpClient,name "http://baidu.com"
func NewHttpClient(logger *zap.Logger) *HttpClient {
	c := &HttpClient{
		Logger: logger,
	}
	return c
}

type HttpClient struct {
	Logger *zap.Logger
}

// (t *HttpClient) PostForm
//  发送POST请求，ContentType为application/x-www-form-urlencoded
func (c *HttpClient) PostForm(ctx *gin.Context, uri string, body []byte) ([]byte, error) {
	return c.Post(ctx, uri, "application/x-www-form-urlencoded", body)
}

func (c *HttpClient) PostFormWithResp(ctx *gin.Context, uri string, body []byte) (retData []byte, retResp *http.Response, retErr error) {
	return c.PostWithResp(ctx, uri, "application/x-www-form-urlencoded", body)
}

// (t *HttpClient) PostJson
//  发送POST请求，ContentType为application/json
func (c *HttpClient) PostJson(ctx *gin.Context, uri string, body []byte) ([]byte, error) {
	return c.Post(ctx, uri, "application/json", body)
}

func (c *HttpClient) PostJsonWithResp(ctx *gin.Context, uri string, body []byte) (retData []byte, retResp *http.Response, retErr error) {
	return c.PostWithResp(ctx, uri, "application/json", body)
}

// (t *HttpClient) Get
//  发送GET请求
//  queryString: 不要带?前缀
func (c *HttpClient) Get(ctx *gin.Context, uri string, queryString []byte) ([]byte, error) {
	urlReq, err := url.Parse(uri)
	if err != nil {
		return nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Get fail, url.Parse err=%s", err)
	}
	if urlReq.RawQuery != "" && len(queryString) != 0 {
		urlReq.RawQuery = urlReq.RawQuery + "&"
	}
	urlReq.RawQuery = urlReq.RawQuery + string(queryString)

	r, err := http.NewRequest("GET", urlReq.String(), nil)
	if err != nil {
		return nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Get fail, err=%s", err)
	}
	return c.Do2(ctx, nil, r)
}

func (c *HttpClient) GetWithResp(ctx *gin.Context, uri string, queryString []byte) (retData []byte, retResp *http.Response, retErr error) {
	urlReq, err := url.Parse(uri)
	if err != nil {
		return nil, nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Get fail, url.Parse err=%s", err)
	}
	if urlReq.RawQuery != "" && len(queryString) != 0 {
		urlReq.RawQuery = urlReq.RawQuery + "&"
	}
	urlReq.RawQuery = urlReq.RawQuery + string(queryString)

	r, err := http.NewRequest("GET", urlReq.String(), nil)
	if err != nil {
		return nil, nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Get fail, err=%s", err)
	}
	return c.do(ctx, nil, r)
}

// (t *HttpClient) Post
//  发送POST请求
func (c *HttpClient) Post(ctx *gin.Context, url string, contentType string, body []byte) ([]byte, error) {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Post fail, err=%s", err)
	}
	r.Header.Set("Content-Type", contentType)
	return c.Do2(ctx, body, r)
}

func (c *HttpClient) PostWithResp(ctx *gin.Context, url string, contentType string, body []byte) (retData []byte, retResp *http.Response, retErr error) {
	r, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "Post fail, err=%s", err)
	}
	r.Header.Set("Content-Type", contentType)
	return c.do(ctx, body, r)
}

func (c *HttpClient) Do(ctx *gin.Context, r *http.Request, ) (retData []byte, retErr error) {
	// 获取body
	var bodyByte []byte
	if r.Body != nil {
		bodyByte, retErr = ioutil.ReadAll(r.Body)
		if retErr != nil {
			retErr = exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "ioutil.ReadAll r.Body fail, err=%s", retErr)
			return nil, retErr
		}
	}
	return c.Do2(ctx, bodyByte, r)
}

func (c *HttpClient) DoWithHttpResponse(ctx *gin.Context, r *http.Request) (retResp *http.Response, retErr error) {
	// 获取body
	var bodyByte []byte
	if r.Body != nil {
		bodyByte, retErr = ioutil.ReadAll(r.Body)
		if retErr != nil {
			retErr = exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "ioutil.ReadAll r.Body fail, err=%s", retErr)
			return nil, retErr
		}
	}
	return c.DoWithHttpResponse2(ctx, bodyByte, r)
}

func (c *HttpClient) DoWithHttpResponse2(ctx *gin.Context, bodyByte []byte, r *http.Request) (retResp *http.Response, retErr error) {
	_, retResp, retErr = c.do(ctx, bodyByte, r)
	return
}

func (c *HttpClient) Do2(ctx *gin.Context, bodyByte []byte, r *http.Request) (retData []byte, retErr error) {
	retData, _, retErr = c.do(ctx, bodyByte, r)
	return
}

func (c *HttpClient) do(ctx *gin.Context, bodyByte []byte, r *http.Request) (retData []byte, retResp *http.Response, retErr error) {
	defer func() {
		if r := recover(); r != nil {
			stack := make([]byte, 1024)
			length := runtime.Stack(stack, true)
			retErr = exception.NewDirpcExceptionf(exception.DIRPC_UNKNOW_EXCEPTION, "panic happens when Flush http\n %s", stack[:length])
		}
	}()

	// 获取http scheme
	scheme := r.URL.Scheme

	// 默认是http
	if len(scheme) == 0 {
		scheme = service.ProtoHttp
	}

	// 判断传入scheme合法性
	if scheme != service.ProtoHttp && scheme != service.ProtoHttps {
		retErr = exception.NewDirpcExceptionf(exception.DIRPC_HTTP_SCHEME_ERROR, "Invalid http scheme. scheme=%s, need http or https.", scheme)
		return nil, nil, retErr
	}

	// 获取body
	if bodyByte != nil {
		rpcContext.CallInfo.Body = utils.Bytes2str(bodyByte)
	}

	rpcContext.CallInfo.Path = itfcName
	rpcContext.CallInfo.Url = r.URL.RequestURI()

	for i := 0; i < retryNum+1; i++ {
		retErr = nil
		// 第几次重试
		rpcContext.CallInfo.RetryFlag = i

		beginTime := time.Now()
		rpcContext.CallInfo.Ip = addr.Ip
		urlPath.Host = addr.Addr
		dialAddr := addr.Addr

		var (
			body    io.ReadCloser
			latency time.Duration
		)
		if bodyByte != nil {
			body = ioutil.NopCloser(bytes.NewReader(bodyByte))
		}
		retResp, err := c.doOnceWithHttpResponse(ctx, r.Method, urlPath, body, r.Header, r.ContentLength, callOpts)

		latency = time.Since(beginTime)

		if retResp.StatusCode < 200 || retResp.StatusCode >= 300 {
			retErr = exception.NewDirpcExceptionf(retResp.StatusCode, "non-2xx response, StatusCode=%d", retResp.StatusCode)
			retResp.Body.Close()
			continue
		}
		retData, err := ioutil.ReadAll(retResp.Body)
		if err != nil {
			retErr = exception.NewDirpcExceptionf(exception.DIRPC_HTTP_REQUEST_FAIL, "ioutil.ReadAll fail, err=%s", err)
			retResp.Body.Close()
			continue
		}

		return retData, retResp, nil
	}
	return nil, nil, retErr
}

func (c *HttpClient) doOnceWithHttpResponse(rpcContext *gin.Context, method string, url url.URL, body io.Reader, header http.Header) (resp *http.Response, retErr error) {
	//addr := url.Host
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		retErr = exception.NewDirpcException(exception.DIRPC_HTTP_CREATE_REQUEST_ERROR, "NewRequest fail")
		return nil, retErr
	}

	return response, nil
}
