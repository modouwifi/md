package api

type Wifi struct {
	ResponseMessage
	W2G Wifi2G `json:"2g"`
	W5G Wifi5G `json:"5g"`
}

type Wifi2G struct {
	Enabled        bool   `json:"enabled"`
	SSID           string `json:"ssid"`
	BroadcastSSID  bool   `json:"broadcastssid"`
	SecurityMode   string `json:"security_mode"`
	Encrypt        string `json:"encrypt"`
	Password       string `json:"password"`
	Power          int    `json:"power"`
	Channel        int    `json:"channel"`
	NetType        int    `json:"net_type"`
	BandWidthMode  int    `json:"band_width_mode"`
	MAC            string `json:"mac"`
	Beacon         int    `json:"beacon"`
	APSDEnabled    bool   `json:"apsd_enabled"`
	APEnabled      bool   `json:"ap_enabled"`
	ShortGIEnabled bool   `json:"shortgi_enabled"`
	WMMEnabled     bool   `json:"wmm_enabled"`
}

type Wifi5G struct {
	Wifi2G
	SameAs2G bool `json:"same_as_2g"`
}

type WifiStatus struct {
	ResponseMessage
	IsEnabled bool `json:"is_enabled"`
}
