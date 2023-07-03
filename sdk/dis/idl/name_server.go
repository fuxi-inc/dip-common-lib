package idl

// name_server api
type DomainResponse struct {
	Domains []string `json:"all_zone"`
}

type SOAResponse struct {
	SOARR ResourceRecord `json:"soa"`
}

type GetZoneResponse struct {
	ZoneDatas []*ZoneData `json:"zone_data"`
}

type ZoneData struct {
	Identifier      string            `json:"identifier"`
	ResourceRecords []*ResourceRecord `json:"resource_records"`
}