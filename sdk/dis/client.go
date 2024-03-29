package dis

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/middleware"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"

	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/net/idna"
)

/**
*  请求方式类如：
*    response,err := NewClient().
					InitLogger(logger).
					InitDis("http://dis.viv.cn").
					ApiRequestDemo( &idl.ApiDemoRequest{})
*/

type Client struct {
	Logger   *zap.Logger //日志组件
	DisHost  string
	DisQHost string
	DaoHost  string
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) InitLogger(logger *zap.Logger) *Client {
	c.Logger = logger
	return c
}

func (c *Client) InitDis(disHost string) *Client {
	c.DisHost = disHost
	return c
}

func (c *Client) InitDisQ(disqHost string) *Client {
	c.DisQHost = disqHost
	return c
}

func (c *Client) InitDao(daoHost string) *Client {
	c.DaoHost = daoHost
	return c
}

// 数据对象属性注册
func (c *Client) ApiDOCreate(ctx *gin.Context, request *idl.ApiDOCreateRequest) (*idl.ApiDisResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Doi = doi

	// punycode编码
	dwdoi, err := Encode_Punycode(request.DwDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("dwdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DwDoi = dwdoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	if request.WhoisData != nil {
		// punycode编码
		whoisdoi, err := Encode_Punycode(request.WhoisData.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("whoisdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		request.WhoisData.Doi = whoisdoi
	}

	disurl := c.DisHost + "/dip/dis-r/doi/register"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiDOCreate] request=%s, response=%s", converter.ToString(request), string(body)))

	fmt.Println("body", string(body), res.StatusCode)
	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", response.Errmsg)
	}
	fmt.Println("reponse", response)
	return response, nil

}

// 数据对象属性批量注册
func (c *Client) ApiDOCreateBatch(ctx *gin.Context, request *idl.ApiDOCreateBatchRequest) (*idl.ApiDisResponse, error) {

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi
	for _, value := range request.BatchData {
		// punycode编码
		doi, err := Encode_Punycode(value.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		value.Doi = doi

		// punycode编码
		dwdoi, err := Encode_Punycode(value.DwDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("dwdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		value.DwDoi = dwdoi
		if value.WhoisData != nil {
			// punycode编码
			whoisdoi, err := Encode_Punycode(value.WhoisData.Doi)
			if err != nil {
				c.Logger.Error(fmt.Sprintf("whoisdoi punycode编码错误：%s", err.Error()))
				return nil, err
			}
			value.WhoisData.Doi = whoisdoi
		}
	}

	disurl := c.DisHost + "/dip/dis-r/doi/batchregister"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiDOCreateBatch] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", response.Errmsg)
	}
	fmt.Println("reponse", response)
	return response, nil

}

// 数据对象属性更新
func (c *Client) ApiDOUpdate(ctx *gin.Context, request *idl.ApiDOUpdateRequest) (*idl.ApiDisResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Doi = doi

	// punycode编码
	if request.NewDoi != "" {
		newdoi, err := Encode_Punycode(request.NewDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("newdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		request.NewDoi = newdoi
	}

	// punycode编码
	dwdoi, err := Encode_Punycode(request.DwDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("dwdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DwDoi = dwdoi

	if request.Authorization != nil {
		// punycode编码
		authdoi, err := Encode_Punycode(request.Authorization.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		request.Authorization.Doi = authdoi
	}

	if request.WhoisData != nil {
		// punycode编码
		whoisdoi, err := Encode_Punycode(request.WhoisData.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("whoisdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		request.WhoisData.Doi = whoisdoi
	}

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-r/doi/update"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiDOUpdate] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil
}

// 数据对象属性批量更新
func (c *Client) ApiDOUpdateBatch(ctx *gin.Context, request *idl.ApiDOUpdateBatchRequest) (*idl.ApiDisResponse, error) {

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi
	for _, value := range request.BatchData {
		// punycode编码
		doi, err := Encode_Punycode(value.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		value.Doi = doi

		// punycode编码
		if value.NewDoi != "" {
			newdoi, err := Encode_Punycode(value.NewDoi)
			if err != nil {
				c.Logger.Error(fmt.Sprintf("newdoi punycode编码错误：%s", err.Error()))
				return nil, err
			}
			value.NewDoi = newdoi
		}

		// punycode编码
		dwdoi, err := Encode_Punycode(value.DwDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("dwdoi punycode编码错误：%s", err.Error()))
			return nil, err
		}
		value.DwDoi = dwdoi

		if value.Authorization != nil {
			// punycode编码
			authdoi, err := Encode_Punycode(value.Authorization.Doi)
			if err != nil {
				c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
				return nil, err
			}
			value.Authorization.Doi = authdoi
		}

		if value.WhoisData != nil {
			// punycode编码
			whoisdoi, err := Encode_Punycode(value.WhoisData.Doi)
			if err != nil {
				c.Logger.Error(fmt.Sprintf("whoisdoi punycode编码错误：%s", err.Error()))
				return nil, err
			}
			value.WhoisData.Doi = whoisdoi
		}
	}
	disurl := c.DisHost + "/dip/dis-r/doi/batchupdate"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiDOUpdateBatch] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil
}

// 注册数据查询（链查询）
func (c *Client) ApiGetRegistrationData(ctx *gin.Context, request *idl.ApiRegDataRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = doi

	disurl := c.DisHost + "/dip/dis-w/whois/getwhoisdata"
	method := constants.GET
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiGetRegistrationData] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil
}

// WHOIS数据局部查询1：根据org进行查询（mongoDB）
func (c *Client) ApiGetRegistrationDataByOrg(ctx *gin.Context, request *idl.ApiWhoisByOrgRequest) (*IDL.CommonResponse, error) {
	org, err := Encode_Punycode(request.Organization)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("org punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Organization = org

	disurl := c.DisHost + "/dip/dis-w/whois/getwhois_by_org"
	method := constants.GET
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiGetRegistrationDataByOrg] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil
}

// WHOIS数据更新
func (c *Client) ApiRegistrationDataUpdate(ctx *gin.Context, request *idl.ApiWhoisUpdateRequest) (*idl.ApiDisResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.WhoisData.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.WhoisData.Doi = doi

	disurl := c.DisHost + "/dip/dis-w/whois/update"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiRegistrationDataUpdate] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil
}

// 查询transaction
func (c *Client) ApiTransactionGet(ctx *gin.Context, request *idl.ApiTransactionInfoRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误:%s", err.Error()))
		return nil, err
	}
	request.DataDoi = doi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误:%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-r/transcationinfo/get"
	method := constants.GET
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}

	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	c.Logger.Info(fmt.Sprintf("update response, %s", converter.ToString(response)))
	return response, nil
}

// 数据对象属性删除
func (c *Client) ApiDODelete(ctx *gin.Context, request *idl.ApiDODeleteRequest) (*idl.ApiDisResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Doi = doi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-r/doi/delete"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiDODelete] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, err
	}
	return response, nil

}

// 授权发起
func (c *Client) ApiAuthInit(ctx *gin.Context, request *idl.ApiAuthInitRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	datadoi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("datadoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = datadoi

	// punycode编码
	authdoi, err := Encode_Punycode(request.Authorization.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Authorization.Doi = authdoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-x/auth/init"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiAuthInit] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if !response.Code.IsSuccess() {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return response, fmt.Errorf("Error response.Errno,error:%s", response.Message)
	}
	return response, nil
}

// 授权发起
func (c *Client) ApiExchangeOwnershipInit(ctx *gin.Context, request *idl.ApiExchangeOwnershipRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	datadoi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("datadoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = datadoi

	dudoi, err := Encode_Punycode(request.DuDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("dudoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DuDoi = dudoi

	// punycode编码
	authdoi, err := Encode_Punycode(request.Authorization.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Authorization.Doi = authdoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-x/exchange_ownership/init"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiAuthInit] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if !response.Code.IsSuccess() {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return response, fmt.Errorf("Error response.Errno,error:%s", response.Message)
	}
	return response, nil
}

// 授权确认
func (c *Client) ApiAuthConf(ctx *gin.Context, request *idl.ApiAuthConfRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	datadoi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("datadoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = datadoi

	// punycode编码
	authdoi, err := Encode_Punycode(request.Authorization.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Authorization.Doi = authdoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-x/auth/confirm"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiAuthConf] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if !response.Code.IsSuccess() {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return response, fmt.Errorf("Error response.Errno,error:%s", response.Message)
	}
	return response, nil
}

// 交换所有权确认
func (c *Client) ApiExchangeOwnershipConf(ctx *gin.Context, request *idl.ApiExchangeOwnershipRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	datadoi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("datadoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = datadoi

	// punycode编码
	authdoi, err := Encode_Punycode(request.Authorization.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("authdoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Authorization.Doi = authdoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-x/exchangeownership/confirm"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiExchangeOwnershipConf] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if !response.Code.IsSuccess() {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return response, fmt.Errorf("Error response.Errno,error:%s", response.Message)
	}
	return response, nil
}

func (c *Client) ApiAuthRevoke(ctx *gin.Context, request *idl.ApiAuthRevRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	datadoi, err := Encode_Punycode(request.DataDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("datadoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DataDoi = datadoi

	// punycode编码
	dudoi, err := Encode_Punycode(request.DuDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("dudoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.DuDoi = dudoi

	// punycode编码
	operatordoi, err := Encode_Punycode(request.SignatureData.OperatorDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("operatordoi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.SignatureData.OperatorDoi = operatordoi

	disurl := c.DisHost + "/dip/dis-x/auth/revoke"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do,error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll,error:%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiAuthRevoke] request=%s, response=%s", converter.ToString(request), string(body)))

	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if !response.Code.IsSuccess() {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return response, fmt.Errorf("Error response.Errno,error:%s", response.Message)
	}
	return response, nil
}

// 数据对象属性查询API返回（对内）
type doqueryresponse struct {
	Code int64                  `json:"code"` // 业务编码
	Data map[string]interface{} `json:"data"` // 成功时返回的数据
	Msg  string                 `json:"msg"`  // 错误描述
}

// 数据对象属性查询
func (c *Client) ApiDOQuery(ctx *gin.Context, request *idl.ApiDOQueryRequest) (*idl.ApiDOQueryResponse, error) {

	method := constants.GET
	baseurl := c.DisQHost + "/DataObject"

	// 将结构体字段转换为字符串
	var typeStrings []string
	for _, t := range request.Type {
		typeStrings = append(typeStrings, string(t))
	}
	typesString := strings.Join(typeStrings, ", ")

	// 构建查询参数
	queryParams := url.Values{}

	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("punycode编码错误：%s", err.Error()))
		return nil, err
	}

	queryParams.Set("doi", doi)
	queryParams.Set("type", typesString)

	// 设置直接查询
	if request.DirectQuery {
		queryParams.Set("direct_query", "true")
	}

	// 将查询参数附加到URL
	url := baseurl + "?" + queryParams.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("创建请求错误：%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	// 发送 GET 请求
	resp, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("请求发送失败：%s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的 Body 内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("读取响应内容失败：%s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[dis-ApiDOQuery] request=%s, response=%s", url, string(body)))

	// TODO: 需要测试返回是否在成功
	var m doqueryresponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("返回内容unmarshal失败：%s", err.Error()))
		return nil, err
	}

	// TODO: 测试响应，后面删掉
	// c.Logger.Error(fmt.Sprintf("Response msg: ", m.Msg)

	response := &idl.ApiDOQueryResponse{}

	response_data := &idl.ApiDOQueryResponseData{}
	response_digest := &idl.DataDigest{}
	response_classgrade := &idl.ClassificationAndGrading{}

	// 遍历返回结果
	for key, value := range m.Data {
		// TODO: 测试用，可以删掉
		log.Printf("键: %s\n", key)

		// 进行断言
		switch key {
		case "dar":
			if str, ok := value.(string); ok {
				response_data.Dar = str
			}
		case "owner":
			if str, ok := value.(string); ok {

				// punycode解码
				newstr, err := Decode_Punycode(str)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("punycode解码错误：%s", err.Error()))
					return nil, err
				}

				response_data.Owner = newstr
			}
		case "pubkey":
			if str, ok := value.(string); ok {
				response_data.PubKey = str
			}
		case "digest":
			if str, ok := value.(string); ok {
				decodedBytes, err := base64.StdEncoding.DecodeString(str)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("digest解码错误: %s", err.Error()))
					return nil, err
				}

				err = json.Unmarshal(decodedBytes, &response_digest)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("digest unmarshal失败: %s", err.Error()))
					return nil, err
				}

				// 	TODO: 需要测试
				response_data.Digest = response_digest
			}

		case "class":
			if str, ok := value.(string); ok {

				num, err := strconv.ParseUint(str, 10, 16)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("uint16转换失败: %s", err.Error()))
					return nil, err
				}

				uint16Val := uint16(num)

				response_classgrade.Class = uint16Val
			}

		case "grade":
			if str, ok := value.(string); ok {

				num, err := strconv.ParseUint(str, 10, 16)
				if err != nil {
					c.Logger.Error(fmt.Sprintf("uint16转换失败: %s", err.Error()))
					return nil, err
				}

				uint16Val := uint16(num)

				response_classgrade.Grade = uint16Val

				response_data.ClassificationAndGrading = response_classgrade

			}

		}

	}

	response.Data = response_data
	response.Errno = IDL.RespCodeType(m.Code)
	response.Errmsg = m.Msg

	return response, nil
}

// 数据对象权属查询
func (c *Client) ApiDOAuthQuery(ctx *gin.Context, request *idl.ApiDOAuthQueryRequest) (*idl.ApiDOQueryResponse, error) {

	method := constants.GET
	baseurl := c.DisQHost + "/DataObject"

	// 构建查询参数
	queryParams := url.Values{}

	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("punycode编码错误：%s", err.Error()))
		return nil, err
	}
	queryParams.Set("doi", doi)

	// punycode编码
	dudoi, err := Encode_Punycode(request.DuDoi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("punycode编码错误：%s", err.Error()))
		return nil, err
	}
	queryParams.Set("dudoi", dudoi)

	queryParams.Set("type", "auth")

	// 设置直接查询
	if request.DirectQuery {
		queryParams.Set("direct_query", "true")
	}

	// 将查询参数附加到URL
	url := baseurl + "?" + queryParams.Encode()

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("创建请求错误：%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	// 发送 GET 请求
	resp, err := client.Do(req)
	if err != nil {
		// TODO: 错误返回格式统一
		c.Logger.Error(fmt.Sprintf("请求发送失败: %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的 Body 内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("读取响应内容失败: %s", err.Error()))
		return nil, err
	}
	c.Logger.Info(fmt.Sprintf("[Dis-ApiDOAuthQuery] request=%s, response=%s", url, string(body)))

	// TODO: 需要测试返回是否在成功
	var m doqueryresponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("返回内容unmarshal失败: %s", err.Error()))
		return nil, err
	}

	// TODO: 测试响应，后面删掉
	// c.Logger.Error(fmt.Sprintf("Response msg: ", m.Msg)

	response := &idl.ApiDOQueryResponse{}

	response_data := &idl.ApiDOQueryResponseData{}
	response_auth := make(map[string]idl.DataAuthorization)

	value := m.Data["auth"]

	// c.Logger.Error(fmt.Sprintf("7test-value", value)

	if str, ok := value.(string); ok {

		// c.Logger.Error(fmt.Sprintf("7test-split str", str)

		decodedBytes, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("auth解码错误: %s", err.Error()))
			return nil, err
		}

		var au idl.DataAuthorization
		err = json.Unmarshal(decodedBytes, &au)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("authorization unmarshal失败: %s", err.Error()))
			return nil, err
		}

		// punycode解码
		newdoi, err := Decode_Punycode(au.Doi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("punycode解码错误：%s", err.Error()))
			return nil, err
		}
		au.Doi = newdoi

		newcreatordoi, err := Decode_Punycode(au.Description.CreatorDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("punycode解码错误：%s", err.Error()))
			return nil, err
		}
		au.Description.CreatorDoi = newcreatordoi

		newparentdoi, err := Decode_Punycode(au.Description.ParentDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("punycode解码错误：%s", err.Error()))
			return nil, err
		}
		au.Description.ParentDoi = newparentdoi

		newpermissiondoi, err := Decode_Punycode(au.Description.PermissionDoi)
		if err != nil {
			c.Logger.Error(fmt.Sprintf("punycode解码错误：%s", err.Error()))
			return nil, err
		}
		au.Description.PermissionDoi = newpermissiondoi

		response_auth[request.DuDoi] = au
		response_data.Auth = response_auth

	}

	response.Data = response_data
	response.Errno = IDL.RespCodeType(m.Code)
	response.Errmsg = m.Msg

	return response, nil
}

func (c *Client) ApiHashManagement(ctx *gin.Context, request *idl.ApiHashManageRequest) (*idl.ApiDisResponse, error) {
	disurl := c.DisHost + "/dip/dis-r/hash/manage"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiHashManagement] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("hash manage response", response)
	return response, nil
}

func (c *Client) ApiWhoisManagement(ctx *gin.Context, request *idl.ApiWhoisManageRequest) (*idl.ApiDisResponse, error) {
	disurl := c.DisHost + "/dip/dis-w/whois/manage"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiWhoisManagement] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &idl.ApiDisResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Errno != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Errmsg))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("whois manage response", response)
	return response, nil

}

func (c *Client) GetPublicKey(ctx *gin.Context, request *idl.ApiPublicKeyRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Doi = doi

	disurl := c.DisHost + "/dip/dis-q/getpublickey"
	method := constants.GET
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiGetPublicKey] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil

}

func (c *Client) GetDataOwner(ctx *gin.Context, request *idl.ApiDataOwnerRequest) (*IDL.CommonResponse, error) {
	// punycode编码
	doi, err := Encode_Punycode(request.Doi)
	if err != nil {
		c.Logger.Error(fmt.Sprintf("doi punycode编码错误：%s", err.Error()))
		return nil, err
	}
	request.Doi = doi

	disurl := c.DisHost + "/dip/dis-q/getdataowner"
	method := constants.GET
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request, error:%s", err.Error()))
		return nil, err
	}
	middleware.InitRequestHeaders(ctx, req)

	res, err := client.Do(req)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error Client.Do, error:%s", err.Error()))
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error ioutil.ReadAll, error:%s", err.Error()))
		return nil, err
	}

	c.Logger.Info(fmt.Sprintf("[Dis-ApiGetDataOwner] request=%s, response=%s", converter.ToString(request), string(body)))
	response := &IDL.CommonResponse{}
	err = json.Unmarshal(body, response)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error response.Unmarshal,error:%s", err.Error()))
		return nil, err
	}
	if response.Code != 0 {
		c.Logger.Error(fmt.Sprintf("Error response.Errno,error:%s", response.Message))
		return nil, fmt.Errorf("Error response.Errno,error:%s", converter.ToString(response))
	}
	fmt.Println("update response", response)
	return response, nil

}

func Encode_Punycode(name string) (string, error) {
	// Punycode 编码
	punycode, err := idna.ToASCII(name)
	if err != nil {
		return punycode, err
	}

	return punycode, err

}

func Decode_Punycode(name string) (string, error) {
	// Punycode 解码
	decodedDomain, err := idna.ToUnicode(name)
	if err != nil {
		return name, err
	}
	return decodedDomain, nil

}
