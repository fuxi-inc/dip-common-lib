package idl

type ZoneRR struct {
	SOARR  ResourceRecord   `json:"SOA_RR"`
	RRList []ResourceRecord `json:"RR_list,omitempty"`
}

type ServiceZonesInDIS struct {
	Zones []string `json:"zones,omitempty"`
}

type GetSOARequest struct {
	Zone string `json:"zone,omitempty"`
}

type UpdateSOASerialRequest struct {
	Zone string `json:"zone,omitempty"`
	Keys string `json:"keys,omitempty"`
}

type GetSOAIncreKeysRequest struct {
	Zone string `json:"zone,omitempty"`
}

type GetZoneRequest struct {
	Zone string `json:"zone,omitempty"`
}

type GetServiceZonesRequest struct {
}

type GetZoneRedisRequest struct {
	Zone string `json:"zone,omitempty"`
}
