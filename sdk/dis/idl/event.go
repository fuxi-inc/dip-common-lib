package idl

// 全量更新事件
type ZoneFullUpdateEvent struct {
	ZoneList []string `json:"zone_list"`
}

// 增量更新事件
type ZoneUpdateEvent struct {
	Operations []ZoneUpdateOperation `json:"operations"`
}

// 增量更新操作
type ZoneUpdateOperation struct {
	Opcode uint8  `json:"opcode"`
	Zone   string `json:"zone"`  //e.g. "viv.cn"
	Label  string `json:"label"` //e.g. "alice"
	Rdata  string `json:"rdata"` //json of Record data
}

const (
	EventZoneUpdate     string = "ZoneUpdateEvent"
	EventZoneFullUpdate string = "ZoneFullUpdateEvent"
	OpcodeSet           uint8  = 0 //create or update
	OpcodeDelete        uint8  = 1 //delete
)

type LabelRecord struct {
	Label string `json:"label"`
	Rdata string `json:"rdata"`
}

type ZoneDataRedis struct {
	Zone      string        `json:"zone"`
	LabelData []LabelRecord `json:"label_data"`
}

type ZoneUpdateFunc = func(string, ZoneUpdateEvent) error
type ZoneFullUpdateFunc = func(string, ZoneFullUpdateEvent) error
