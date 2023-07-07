package dis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/constants"
	"github.com/fuxi-inc/dip-common-lib/sdk/dis/idl"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	disurl := c.DisHost + "/dip/dis-r/doi/register"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	//req.Header.Add(constants.HeaderAuthorization, "<Authorization>")
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)

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

// 数据对象属性更新
func (c *Client) ApiDOUpdate(ctx *gin.Context, request *idl.ApiDOUpdateRequest) (*idl.ApiDisResponse, error) {
	disurl := c.DisHost + "/dip/dis-r/doi/update"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	//req.Header.Add(constants.HeaderAuthorization, "<Authorization>")
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)

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

// 数据对象属性删除
func (c *Client) ApiDODelete(ctx *gin.Context, request *idl.ApiDODeleteRequest) (*idl.ApiDisResponse, error) {
	disurl := c.DisHost + "/dip/dis-r/doi/delete"
	method := constants.POST
	payload := strings.NewReader(converter.ToString(request))

	client := &http.Client{}
	req, err := http.NewRequest(method, disurl, payload)

	if err != nil {
		c.Logger.Error(fmt.Sprintf("Error creating request,error:%s", err.Error()))
		return nil, err
	}
	//req.Header.Add(constants.HeaderAuthorization, "<Authorization>")
	req.Header.Add(constants.HeaderContentType, constants.MIMEApplicationJSON)

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
func (c *Client) ApiAuthInit(ctx *gin.Context, request *idl.ApiAuthInitRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 授权确认
func (c *Client) ApiAuthConf(ctx *gin.Context, request *idl.ApiAuthConfRequest) (*idl.ApiDisResponse, error) {
	return nil, nil
}

// 注册数据查询
func (c *Client) ApiRegDataQuery(ctx *gin.Context, request *idl.ApiRegDataRequest) (*idl.ApiRegDataQueryResponse, error) {
	return nil, nil
}

// 数据对象属性查询API返回（对内）
type doqueryresponse struct {
	Code int64                  `json:"code"` // 业务编码
	Data map[string]interface{} `json:"data"` // 成功时返回的数据
	Msg  string                 `json:"msg"`  // 错误描述
}

// 数据对象属性查询
func (c *Client) ApiDOQuery(ctx *gin.Context, request *idl.ApiDOQueryRequest) (*idl.ApiDOQueryResponse, error) {

	baseurl := c.DisQHost + "/DataObject"

	// 将结构体字段转换为字符串
	var typeStrings []string
	for _, t := range request.Type {
		typeStrings = append(typeStrings, string(t))
	}
	typesString := strings.Join(typeStrings, ", ")

	// 构建查询参数
	queryParams := url.Values{}
	queryParams.Set("doi", request.Doi)
	queryParams.Set("type", typesString)

	// 将查询参数附加到URL
	url := baseurl + "?" + queryParams.Encode()

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		// TODO: 错误返回格式统一
		log.Println("请求发送失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的 Body 内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应内容失败:", err)
		return nil, err
	}

	// TODO: 需要测试返回是否在成功
	var m doqueryresponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("返回内容unmarshal失败:", err)
		return nil, err
	}

	// TODO: 测试响应，后面删掉
	log.Println("Response msg: ", m.Msg)

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
				response_data.Owner = str
			}
		case "pubkey":
			if str, ok := value.(string); ok {
				response_data.PubKey = str
			}
		case "digest":
			if str, ok := value.(string); ok {
				b := []byte(str)
				err := json.Unmarshal(b, &response_digest)
				if err != nil {
					log.Println("digest unmarshal失败:", err)
					return nil, err
				}

				// 	TODO: 需要测试
				response_data.Digest = response_digest
			}

		case "class":
			if str, ok := value.(string); ok {

				num, err := strconv.ParseUint(str, 10, 16)
				if err != nil {
					log.Println("uint16转换失败:", err)
					return nil, err
				}

				uint16Val := uint16(num)

				response_classgrade.Class = uint16Val
			}

		case "grade":
			if str, ok := value.(string); ok {

				num, err := strconv.ParseUint(str, 10, 16)
				if err != nil {
					log.Println("uint16转换失败:", err)
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

	baseurl := c.DisQHost + "/DataObject"

	// 构建查询参数
	queryParams := url.Values{}
	queryParams.Set("doi", request.Doi)
	queryParams.Set("dudoi", request.DuDoi)

	// 将查询参数附加到URL
	url := baseurl + "?" + queryParams.Encode()

	// 发送 GET 请求
	resp, err := http.Get(url)
	if err != nil {
		// TODO: 错误返回格式统一
		log.Println("请求发送失败:", err)
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应的 Body 内容
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应内容失败:", err)
		return nil, err
	}

	// TODO: 需要测试返回是否在成功
	var m doqueryresponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Println("返回内容unmarshal失败:", err)
		return nil, err
	}

	// TODO: 测试响应，后面删掉
	log.Println("Response msg: ", m.Msg)

	response := &idl.ApiDOQueryResponse{}

	response_data := &idl.ApiDOQueryResponseData{}
	response_auth := make(map[string]idl.DataAuthorization)

	value := m.Data["auth"]

	if str, ok := value.(string); ok {
		var au idl.DataAuthorization
		err = json.Unmarshal([]byte(str), &au)
		if err != nil {
			log.Println("authorization unmarshal失败:", err)
			return nil, err
		}

		response_auth[request.DuDoi] = au
		response_data.Auth = response_auth

	}

	response.Data = response_data
	response.Errno = IDL.RespCodeType(m.Code)
	response.Errmsg = m.Msg

	return response, nil
}
