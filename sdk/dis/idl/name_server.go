package idl

// name_server api
type AllZonesResponse struct {
	Zones []string `json:"zones"`
}

type SOAResponse struct {
	SOARR ResourceRecord `json:"soa"`
}

type ZoneResponse struct {
	ZoneDatas []*ZoneData `json:"zone_datas"`
}

type ZoneData struct {
	Identifier      string            `json:"identifier"`
	ResourceRecords []*ResourceRecord `json:"resource_records"`
}
