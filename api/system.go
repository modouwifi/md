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
