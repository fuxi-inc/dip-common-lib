package idl

type RegistrationData struct {
	Doi          string   `json:"doi,omitempty"`
	Organization []string `json:"organization,omitempty"` // 组织
	Contact      []string `json:"contact,omitempty"`      // 联系方式
	IP           []string `json:"ip,omitempty"`           // IP地址
	ASN          []string `json:"asn,omitempty"`
}
