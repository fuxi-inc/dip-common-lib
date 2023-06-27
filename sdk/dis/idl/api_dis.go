package idl

// 通用响应（无返回data）。在没有特别定义时，用这种响应格式
type ApiDisResponse struct {
	Errno  DisRespErrno `json:"errno"`
	Errmsg string       `json:"errmsg"`
}

// 数据对象属性注册
type ApiDOCreateRequest struct {
	Doi       string            `json:"doi,omitempty"`
	DwDoi     string            `json:"dw_doi,omitempty"`
	PubKey    string            `json:"pub_key,omitempty"`
	WhoisData *RegistrationData `json:"registration_data,omitempty"` // 类WHOIS注册数据
	Sign      string            `json:"sign,omitempty"`              // 使用DW私钥对其自身的DOI签名
}

// 数据对象属性更新
type ApiDOUpdateRequest struct {
	Doi                      string                    `json:"doi,omitempty"`
	NewDoi                   string                    `json:"new_doi,omitempty"` // 更新后的DO标识
	DwDoi                    string                    `json:"dw_doi,omitempty"`
	PubKey                   string                    `json:"pub_key,omitempty"`
	Dar                      string                    `json:"dar,omitempty"`                    // DOI地址
	Digest                   *DataDigest               `json:"digest,omitempty"`                 // 数据内容摘要
	Authorization            *[]DataAuthorization      `json:"authorization,omitempty"`          // 授权信息数组
	ClassificationAndGrading *ClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
	WhoisData                *RegistrationData         `json:"whois_data,omitempty"`             // WHOIS注册数据
	Sign                     string                    `json:"sign,omitempty"`                   // 使用DW私钥对其自身的DOI签名
}

// 数据对象属性删除
type ApiDODeleteRequest struct {
	Doi   string `json:"doi,omitempty"`
	DwDoi string `json:"dw_doi,omitempty"` // 更新后的DO标识
	Sign  string `json:"sign,omitempty"`   // 使用DW私钥对其自身的DOI签名
}

// 授权发起
type ApiAuthInitRequest struct {
	DataDoi       string            `json:"data_doi,omitempty"`
	Authorization DataAuthorization `json:"authorization,omitempty"` // 授权信息数组
	InitiatorDOI  string            `json:"initiator_doi,omitempty"` //发起者DW或DU的DOI
	Sign          string            `json:"sign,omitempty"`          // 发起者私钥对其自身的DOI签名
}

// 授权确认
type ApiAuthConfRequest struct {
	DataDoi       string            `json:"data_doi,omitempty"`
	Authorization DataAuthorization `json:"authorization,omitempty"` // 授权信息数组
	ConfirmerDOI  string            `json:"confirmer_doi,omitempty"` //确认者DW或DU的DOI
	Sign          string            `json:"sign,omitempty"`          // 使用确认这私钥对其自身的DOI签名
}

// 注册数据查询

type ApiRegDataQueryResponse struct {
	Errno  DisRespErrno      `json:"errno"`
	Errmsg string            `json:"errmsg"`
	Data   *RegistrationData `json:"data,omitempty"`
}

// 数据对象属性查询
type SearchType string

const (
	Dar        SearchType = "dar"         // 存储地址
	PubKey     SearchType = "pub_key"     // 公钥
	Auth       SearchType = "auth"        // 权属
	Digest     SearchType = "digest"      // 数据内容摘要
	ClassGrade SearchType = "class_grade" // 数据分类分级
)

type ApiDOQueryRequest struct {
	Doi  string       `json:"doi,omitempty"`
	Type []SearchType `json:"type,omitempty"` // 查询类型
}

// 数据对象权属查询
type ApiDOAuthQueryRequest struct {
	Doi   string       `json:"doi,omitempty"`
	DuDoi string       `json:"dudoi,omitempty"`
	Type  []SearchType `json:"type,omitempty"` // 查询类型
}

// 权地址查询响应
type ApiDOQueryResponse struct {
	Errno  DisRespErrno            `json:"errno"`
	Errmsg string                  `json:"errmsg"`
	Data   *ApiDOQueryResponseData `json:"data,omitempty"`
}

type ApiDOQueryResponseData struct {
	PubKey                   string                    `json:"pub_key,omitempty"`
	Dar                      string                    `json:"dar,omitempty"`                    // DOI地址
	Auth                     *[]DataAuthorization      `json:"authorization,omitempty"`          // 权属
	Digest                   *DataDigest               `json:"digest,omitempty"`                 // 数据内容摘要
	ClassificationAndGrading *ClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
}
