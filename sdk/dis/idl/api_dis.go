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
	Type      AuthorizationType `json:"type,omitempty"`       // 权属类型。0开头表示所有者，1开头表示使用者
	WhoisData *ApiWhoisData     `json:"whois_data,omitempty"` // WHOIS注册数据
	Sign      string            `json:"sign,omitempty"`       // 使用DW私钥对其自身的DOI签名
}

// 数据对象属性更新
type ApiDOUpdateRequest struct {
	Doi                      string                       `json:"doi,omitempty"`
	NewDoi                   string                       `json:"new_doi,omitempty"` // 更新后的DO标识
	DwDoi                    string                       `json:"dw_doi,omitempty"`
	PubKey                   string                       `json:"pub_key,omitempty"`
	Dar                      string                       `json:"dar,omitempty"`                    // DOI地址
	Digest                   *ApiDigest                   `json:"digest,omitempty"`                 // 数据内容摘要
	Type                     AuthorizationType            `json:"type,omitempty"`                   // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation             string                       `json:"confirmation,omitempty"`           // 确权信息。DW私钥对数据摘要的签名
	Description              *ApiDescription              `json:"description,omitempty"`            // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	ClassificationAndGrading *ApiClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
	WhoisData                *ApiWhoisData                `json:"whois_data,omitempty"`             // WHOIS注册数据
	Sign                     string                       `json:"sign,omitempty"`                   // 使用DW私钥对其自身的DOI签名
}

// 数据对象属性删除
type ApiDODeleteRequest struct {
	Doi   string `json:"doi,omitempty"`
	DwDoi string `json:"dw_doi,omitempty"` // 更新后的DO标识
}

// 授权发起
type ApiAuthInitRequest struct {
	DataDoi      string            `json:"data_doi,omitempty"`
	DuDoi        string            `json:"du_doi,omitempty"`       // 被授权的数据使用者身份标识
	Type         AuthorizationType `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string            `json:"confirmation,omitempty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string            `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	Sign         string            `json:"sign,omitempty"`         // 所有者DW或授权发起者私钥对其自身的DOI签名
}

// 授权确认
type ApiAuthConfRequest struct {
	DataDoi      string            `json:"data_doi,omitempty"`
	DuDoi        string            `json:"du_doi,omitempty"`       // 被授权的数据使用者身份标识
	Type         AuthorizationType `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string            `json:"confirmation,omitempty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string            `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
	Sign         string            `json:"sign,omitempty"`         // 使用DU私钥对其自身的DOI签名
}

// 注册数据查询

type ApiRegDataQueryResponse struct {
	Errno  DisRespErrno  `json:"errno"`
	Errmsg string        `json:"errmsg"`
	Data   *ApiWhoisData `json:"data,omitempty"`
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

// 权地址查询响应
type ApiDOQueryResponse struct {
	Errno  DisRespErrno            `json:"errno"`
	Errmsg string                  `json:"errmsg"`
	Data   *ApiDOQueryResponseData `json:"data,omitempty"`
}

type ApiDOQueryResponseData struct {
	PubKey                   string                       `json:"pub_key,omitempty"`
	Dar                      string                       `json:"dar,omitempty"`                    // DOI地址
	Authorization            *ApiAuthorization            `json:"authorization,omitempty"`          // 权属
	Digest                   *ApiDigest                   `json:"digest,omitempty"`                 // 数据内容摘要
	ClassificationAndGrading *ApiClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
}

// 数据对象属性定义
type ApiDigest struct {
	Algorithm string `json:"algorithm,omitempty"` // 摘要算法
	Result    string `json:"result,omitempty"`    // 摘要计算结果
}

type ApiClassificationAndGrading struct {
	Class uint16 `json:"class,omitempty"` // 数据分类
	Grade uint16 `json:"grade,omitempty"` // 数据分级
}

type ApiWhoisData struct {
	Doi          string   `json:"doi,omitempty"`
	Organization []string `json:"organization,omitempty"` // 组织
	Contact      []string `json:"contact,omitempty"`      // 联系方式
	IP           []string `json:"ip,omitempty"`           // IP地址
	ASN          []string `json:"asn,omitempty"`
}

type ApiAuthorization struct {
	Type         AuthorizationType `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string            `json:"confirmation,omitempty"` // 确权信息。DW私钥对数据摘要的签名
	Description  string            `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
}

type ApiDescription struct {
	PermissionDoi string `json:"permission_doi,omitempty"` // 权限定义DOI
	CreatorDoi    string `json:"creator_doi,omitempty"`    // 权限创建者DOI
}
