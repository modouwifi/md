package api

type SecurityEnabled struct {
	ResponseMessage
	Enabled bool `json:"enabled"`
}
