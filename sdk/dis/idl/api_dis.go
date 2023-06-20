package idl

// 通用响应（无返回data）。在没有特别定义时，用这种响应格式
type ApiDisResponse struct {
	Errno  int64  `json:"errno"`
	Errmsg string `json:"errmsg"`
}

// 权址登记
type ApiDORegistryRequest struct {
	Doi                      string                       `json:"doi,omitempty"`
	DwDoi                    string                       `json:"dw_doi,omitempty"`
	Pubkey                   string                       `json:"pubkey,omitempty"`
	Dar                      string                       `json:"dar,omitempty"`                    // DOI地址
	Digest                   *ApiDigest                   `json:"digest,omitempty"`                 // 数据内容摘要
	Type                     uint8                        `json:"type,omitempty"`                   // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation             string                       `json:"confirmation,omitwmpty"`           // 确权信息。DW私钥对数据摘要的签名
	Description              string                       `json:"description,omitempty"`            // json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	ClassificationAndGrading *ApiClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
	WhoisData                *ApiWhoisData                `json:"whois_data,omitempty"`             // WHOIS注册数据
	Sign                     string                       `json:"sign,omitempty"`                   // 使用DW私钥对其自身的DOI签名
}

// 授权发起
type ApiAuthInitRequest struct {
	DataDoi      string `json:"data_doi,omitempty"`
	DuDoi        string `json:"du_doi,omitempty"`       // 被授权的数据使用者身份标识
	Type         uint8  `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string `json:"confirmation,omitwmpty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	Sign         string `json:"sign,omitempty"`         // 所有者DW或授权发起者私钥对其自身的DOI签名
}

// 授权确认
type ApiAuthConfRequest struct {
	DataDoi      string `json:"data_doi,omitempty"`
	DuDoi        string `json:"du_doi,omitempty"`       // 被授权的数据使用者身份标识
	Type         uint8  `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string `json:"confirmation,omitwmpty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	Sign         string `json:"sign,omitempty"`         // 使用DU私钥对其自身的DOI签名
}

// 权址查询
type SearchType string

const (
	Dar        SearchType = "dar"         // 存储地址
	Pubkey     SearchType = "pubkey"      // 公钥
	Auth       SearchType = "auth"        // 权属
	Digest     SearchType = "digest"      // 数据内容摘要
	ClassGrade SearchType = "class_grade" // 数据分类分级
)

type ApiDOQueryRequest struct {
	Doi  string       `json:"doi,omitempty"`
	Type []SearchType `json:"type,omitempty"` // 查询类型
}

// 权地址查询响应
type ApiDOQueryResponse struct {
	Errno  int64                   `json:"errno"`
	Errmsg string                  `json:"errmsg"`
	Data   *ApiDOQueryResponseData `json:"data,omitempty"`
}

type ApiDOQueryResponseData struct {
	Pubkey                   string                       `json:"pubkey,omitempty"`
	Dar                      string                       `json:"dar,omitempty"`                    // DOI地址
	Authorization            *ApiAuthorization            `json:"authorization,omitempty"`          // 权属
	Digest                   *ApiDigest                   `json:"digest,omitempty"`                 // 数据内容摘要
	ClassificationAndGrading *ApiClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
}

// 数据对象属性定义
type ApiDigest struct {
	Algorithm string `json:"algorithm,omitempty"` // 摘要算法
	Result    string `json:"result,omitwmpty"`    // 摘要计算结果
}

type ApiClassificationAndGrading struct {
	Class uint16 `json:"class,omitempty"` // 数据分类
	Grade uint16 `json:"grade,omitwmpty"` // 数据分级
}

type ApiWhoisData struct {
	Organization []string `json:"organization,omitempty"` // 组织
	Contact      []string `json:"contact,omitwmpty"`      // 联系方式
	IP           []string `json:"ip,omitempty"`           // IP地址
	ASN          []string `json:"asn,omitempty"`
}

type ApiAuthorization struct {
	Type         uint8  `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string `json:"confirmation,omitwmpty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
}
