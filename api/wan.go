package api

type Wan struct {
	ResponseMessage
	Type            string `json:"type"`
	IP              string `json:"ip"`
	Mask            string `json:"mask"`
	Gateway         string `json:"gateway"`
	DNS1            string `json:"dns1"`
	DNS2            string `json:"dns2"`
	MTU             int    `json:"mtu"`
	STP             bool   `json:"stp"`
	Account         string `json:"account"`
	Password        string `json:"password"`
	MACCloneEnabled bool   `json:"macCloneEnabled"`
	MACCloneMac     string `json:"macCloneMac"`
	Uptime          int    `json:"uptime"`
}

type WanDHCP struct {
	ResponseMessage
	DNS1 string `json:"dns1"`
	DNS2 string `json:"dns2"`
	MTU  int    `json:"mtu"`
	STP  bool   `json:"stp"`
}

type WanPPPOE struct {
	ResponseMessage
	Account      string `json:"account"`
	Password     string `json:"password"`
	PPPOEMethod  string `json:"pppoe_method"`
	PedialPeriod int    `json:"pedial_period"`
	IdleTime     int    `json:"idle_time"`
	Status       int    `json:"status"`
}

type WanSTATIC struct {
	ResponseMessage
	IP      string `json:"ip"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
	DNS1    int    `json:"dns1"`
	DNS2    int    `json:"dns2"`
	MTU     int    `json:"mtu"`
	STP     bool   `json:"stp"`
}
