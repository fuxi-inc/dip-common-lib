package idl

import (
	"github.com/fuxi-inc/dip-common-lib/IDL"
	"github.com/fuxi-inc/dip-common-lib/utils/converter"
)

// 通用响应（无返回data）。在没有特别定义时，用这种响应格式
type ApiDisResponse struct {
	Errno  IDL.RespCodeType `json:"errno"`
	Errmsg string           `json:"errmsg"`
}

// 数据对象属性注册
type ApiDOCreateRequest struct {
	Doi               string            `json:"doi,omitempty"`
	DwDoi             string            `json:"dw_doi,omitempty"`
	PubKey            string            `json:"pub_key,omitempty"`
	WhoisData         *RegistrationData `json:"whois_data,omitempty"` // 类WHOIS注册数据
	IDL.SignatureData                   //统一共用的加签验签结构，字段均为必填项
}

type CreateRequestData struct {
	Doi       string            `json:"doi,omitempty"`
	DwDoi     string            `json:"dw_doi,omitempty"`
	PubKey    string            `json:"pub_key,omitempty"`
	WhoisData *RegistrationData `json:"whois_data,omitempty"` // 类WHOIS注册数据
}

// 数据对象属性批量注册
type ApiDOCreateBatchRequest struct {
	BatchData         []*CreateRequestData `json:"batch_data,omitempty"`
	IDL.SignatureData                      //统一共用的加签验签结构，字段均为必填项
}

// 数据对象属性更新
type ApiDOUpdateRequest struct {
	Doi                      string                    `json:"doi,omitempty"`
	NewDoi                   string                    `json:"new_doi,omitempty"` // 更新后的DO标识
	DwDoi                    string                    `json:"dw_doi,omitempty"`  //更新所有者
	PubKey                   string                    `json:"pub_key,omitempty"`
	Dar                      string                    `json:"dar,omitempty"`                    // DOI地址
	Digest                   *DataDigest               `json:"digest,omitempty"`                 // 数据内容摘要
	Authorization            *DataAuthorization        `json:"authorization,omitempty"`          // 授权信息数组
	ClassificationAndGrading *ClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
	WhoisData                *RegistrationData         `json:"whois_data,omitempty"`             // WHOIS注册数据	IDL.SignatureData //统一共用的加签验签结构，字段均为必填项
	IDL.SignatureData                                  //统一共用的加签验签结构，字段均为必填项
}

type UpdateRequestData struct {
	Doi                      string                    `json:"doi,omitempty"`
	NewDoi                   string                    `json:"new_doi,omitempty"` // 更新后的DO标识
	DwDoi                    string                    `json:"dw_doi,omitempty"`  //更新所有者
	PubKey                   string                    `json:"pub_key,omitempty"`
	Dar                      string                    `json:"dar,omitempty"`                    // DOI地址
	Digest                   *DataDigest               `json:"digest,omitempty"`                 // 数据内容摘要
	Authorization            *DataAuthorization        `json:"authorization,omitempty"`          // 授权信息数组
	ClassificationAndGrading *ClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
	WhoisData                *RegistrationData         `json:"whois_data,omitempty"`             // WHOIS注册数据
}

// 数据对象属性批量更新
type ApiDOUpdateBatchRequest struct {
	BatchData         []*UpdateRequestData `json:"batch_data,omitempty"`
	IDL.SignatureData                      //统一共用的加签验签结构，字段均为必填项
}

// 数据对象属性删除
type ApiDODeleteRequest struct {
	Doi               string `json:"doi,omitempty"`
	IDL.SignatureData        //统一共用的加签验签结构，字段均为必填项
}

// // TODO: WHOIS数据更新
// type ApiRegistrationDataUpdateRequest struct {
// }

// 授权发起
type ApiAuthInitRequest struct {
	DataDoi           string            `json:"data_doi,omitempty"`
	Authorization     DataAuthorization `json:"authorization,omitempty"` // 授权信息数组
	Fields            map[string]string `json:"fields,omitempty"`        // 扩展字段，用于发送通知
	IDL.SignatureData                   //统一共用的加签验签结构，字段均为必填项
}

// 授权确认
type ApiAuthConfRequest struct {
	DataDoi           string            `json:"data_doi,omitempty"`
	Authorization     DataAuthorization `json:"authorization,omitempty"` // 授权信息数组
	Fields            map[string]string `json:"fields,omitempty"`        //扩展字段，用于发送通知
	IDL.SignatureData                   //统一共用的加签验签结构，字段均为必填项
}

// 授权撤销
type ApiAuthRevRequest struct {
	DataDoi           string            `json:"data_doi,omitempty"`
	DuDoi             string            `json:"du_doi,omitempty"` // 数据使用者DOI
	Fields            map[string]string `json:"fields,omitempty"` //扩展字段，用于发送通知
	IDL.SignatureData                   //统一共用的加签验签结构，字段均为必填项
}

type ApiRegDataRequest struct {
	DataDoi           string `json:"data_doi,omitempty"`
	IDL.SignatureData        //统一共用的加签验签结构，字段均为必填项
}

// 注册数据查询

type ApiRegDataQueryResponse struct {
	Errno  IDL.RespCodeType  `json:"errno"`
	Errmsg string            `json:"errmsg"`
	Data   *RegistrationData `json:"data,omitempty"`
}

// 数据对象TX ID查询
type ApiTransactionInfoRequest struct {
	DataDoi           string `json:"data_doi"`
	IDL.SignatureData        //统一共用的加签验签结构，字段均为必填项
}

// 数据对象属性查询
type SearchType string

const (
	Dar        SearchType = "dar"        // 存储地址
	Owner      SearchType = "owner"      // 所有者DO
	PubKey     SearchType = "pubkey"     // 公钥
	Auth       SearchType = "auth"       // 权属
	Digest     SearchType = "digest"     // 数据内容摘要
	ClassGrade SearchType = "classgrade" // 数据分类分级
)

type ApiDOQueryRequest struct {
	Doi         string       `json:"doi,omitempty"`
	Type        []SearchType `json:"type,omitempty"` // 查询类型
	DirectQuery bool         `json:"directquery,omitempty"`
}

// 数据对象权属查询
type ApiDOAuthQueryRequest struct {
	Doi         string       `json:"doi,omitempty"`
	DuDoi       string       `json:"dudoi,omitempty"`
	Type        []SearchType `json:"type,omitempty"` // 查询类型，只包括auth
	DirectQuery bool         `json:"directquery,omitempty"`
}

// 权地址查询响应
type ApiDOQueryResponse struct {
	Errno  IDL.RespCodeType        `json:"errno"`
	Errmsg string                  `json:"errmsg"`
	Data   *ApiDOQueryResponseData `json:"data"`
}

type ApiDOQueryResponseData struct {
	PubKey                   string                       `json:"pub_key"`
	Owner                    string                       `json:"owner"`
	Dar                      string                       `json:"dar"`                    // DOI地址
	Auth                     map[string]DataAuthorization `json:"authorization"`          // 权属，key的内容也为权属对象
	Digest                   *DataDigest                  `json:"digest"`                 // 数据内容摘要
	ClassificationAndGrading *ClassificationAndGrading    `json:"classification_grading"` // 数据分类分级信息
}

func (s *ApiDOQueryRequest) ToString() string {
	return converter.ToString(s)
}

// WHOIS数据更新
type ApiWhoisUpdateRequest struct {
	WhoisData         *RegistrationData `json:"whois_data,omitempty"` // WHOIS注册数据
	IDL.SignatureData                   //统一共用的加签验签结构，字段均为必填项
}

type ApiHashManageRequest struct {
}

type ApiWhoisManageRequest struct {
	/* nothing */
}
