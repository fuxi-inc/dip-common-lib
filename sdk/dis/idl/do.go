package idl

// 数据对象
type DO struct {
	Doi                      string                    `json:"doi,omitempty"`
	PubKey                   string                    `json:"pub_key,omitempty"`                //公钥
	Dar                      string                    `json:"dar,omitempty"`                    // 数据在DAO中的存储地址
	Digest                   *DataDigest               `json:"digest,omitempty"`                 // 数据内容摘要
	Authorization            *[]DataAuthorization      `json:"authorization,omitempty"`          // 权属信息
	ClassificationAndGrading *ClassificationAndGrading `json:"classification_grading,omitempty"` // 数据分类分级信息
}

// 数据摘要
type DataDigest struct {
	Algorithm string `json:"algorithm,omitempty"` // 摘要算法
	Result    string `json:"result,omitempty"`    // 摘要计算结果
}

// 权属信息
type DataAuthorization struct {
	Doi          string                 `json:"doi,omitempty"`          //权属对象。
	Type         AuthorizationType      `json:"type,omitempty"`         // 权属类型。0开头表示所有者，1开头表示使用者
	Confirmation string                 `json:"confirmation,omitempty"` // 确权信息。DW私钥对数据摘要的签名
	Description  *PermissionDescription `json:"description,omitempty"`  // 权益特征。json格式，包括权限定义DOI（permission），权限创建者DOI（creator），及解密密钥（key）
}

// 权属信息中的权益特征
type PermissionDescription struct {
	PermissionDoi string `json:"permission_doi,omitempty"` // 权限定义DOI
	CreatorDoi    string `json:"creator_doi,omitempty"`    // 权限创建者DOI
	Key           string `json:"key,omitempty"`            //权限密钥，权属对象公钥加密的数据内容加密对称密钥的16进制表示（长度为256）
}

// 分类分级
type ClassificationAndGrading struct {
	Class uint16 `json:"class,omitempty"` // 数据分类
	Grade uint16 `json:"grade,omitempty"` // 数据分级。如果是加密数据，第1位是1。
}
