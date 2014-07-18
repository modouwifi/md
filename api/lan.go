package api

type Lan struct {
	IP          string `json:"ip"`
	NetMask     string `json:"net_mask"`
	DHCPEnabled bool   `json:"dhcp_enabled"`
	IPAddrStart string `json:"ipaddr_start"`
	IPAddrEnd   string `json:"ipaddr_end"`
	Gateway     string `json:"gateway"`
	DHCPNetMask string `json:"dhcp_net_mask"`
	DNS1Method  string `json:"dns1Method"`
	DNS2Method  string `json:"dns2Method"`
	DNS1        string `json:"dns1"`
	DNS2        string `json:"dns2"`
	Time        int    `json:"time"`
	MAC         string `json:"mac"`
}
