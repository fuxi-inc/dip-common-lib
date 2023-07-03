package idl

type ZoneRR struct {
	SOARR  ResourceRecord   `json:"SOA_RR"`
	RRList []ResourceRecord `json:"RR_list,omitempty"`
}

type ServiceZonesInDIS struct {
	Zones []string `json:"zones,omitempty"`
}
