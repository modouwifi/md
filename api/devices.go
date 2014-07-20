package api

type DevicesCables struct {
	ResponseMessage
	Devices []DeviceCable
}

type DevicesWifis struct {
	ResponseMessage
	Devices []DeviceWifi
}

type Device struct {
	Platform   string `json:"platform"`
	TotalSpeed int    `json:"total_speed"`
	Down       string `json:"down"`
	DownSpeed  int    `json:"down_speed"`
	DownLimit  int    `json:"down_limit"`
	Up         string `json:"up"`
	UpSpeed    int    `json:"up_speed"`
	UpLimit    int    `json:"up_limit"`
	Hostname   string `json:"host_name"`
	Type       int    `json:"type"`
	IP         string `json:"ip"`
	MAC        string `json:"mac"`
	LeftTime   int    `json:"leftTime"`
	Time       int    `json:"time"`
	Tag        string `json:"tag"`
	Local      int    `json:"local"`
	ConnType   int    `json:"conn_type"`
}

type DeviceCable struct {
	Device
}

type DeviceWifi struct {
	Device
	Single int `json:"single"`
	RSSI0  int `json:"rssi0"`
	RSSI1  int `json:"rssi1"`
	RSSI2  int `json:"rssi2"`
}

type DeviceDisk struct {
	ResponseMessage
	TotalSize int     `json:"totalSize"`
	Left      float32 `json:"left"`
	Name      string  `json:"name"`
	Time      int     `json:"time"`
}
