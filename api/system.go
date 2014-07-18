package api

type SystemVersion struct {
	ResponseMessage
	Track    string `json:"track"`
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
}

type SystemLatestVersion struct {
	ResponseMessage
	Filename    string `json:"filename"`
	Version     string `json:"version"`
	Releasenote string `json:"releasenote"`
}

type SystemTime struct {
	ResponseMessage
	Time     int `json:"time"`
	TimeType int `json:"time_type"`
}

type SystemCable struct {
	ResponseMessage
	Wan  bool `json:"time"`
	Lan1 bool `json:"lan1"`
	Lan2 bool `json:"lan2"`
	USB  bool `json:"usb"`
}
