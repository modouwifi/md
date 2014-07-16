package api

type Wan struct {
	ResponseMessage
	Type            string `json:"type"`
	IP              string `json:"ip"`
	Mask            string `json:"mask"`
	Gateway         string `json:"gateway"`
	DNS1            string `json:"dns1"`
	DNS2            string `json:"dns2"`
	MTU             string `json:"mtu"`
	STP             bool   `json:"stp"`
	Account         string `json:"account"`
	Password        string `json:"password"`
	MACCloneEnabled bool   `json:"macCloneEnabled"`
	MACCloneMac     string `json:"macCloneMac"`
	Uptime          int    `json:"uptime"`
}
