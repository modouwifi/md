package api

type ResponseMessage struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
}
