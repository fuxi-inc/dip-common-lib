package idl

type ResourceRecord struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Class string `json:"class"`
	Ttl   int    `json:"ttl"`
	Rdata string `json:"rdata"`
}

type SOAData struct {
	Mname   string `json:"mname"`
	Rname   string `json:"rname"`
	Serial  int    `json:"serial"`
	Refresh int    `json:"refresh"`
	Retry   int    `json:"retry"`
	Expire  int    `json:"expire"`
	Minimum int    `json:"minimum"`
}