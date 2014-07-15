package api

type SystemVersion struct {
	ResponseMessage
	Track    string `json:"track"`
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
}
