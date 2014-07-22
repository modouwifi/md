package api

type App struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Icon        string `json:"icon,omitempty"`
	Source      string `json:"source,omitempty"`
	Version     string `json:"size,omitempty"`
	Size        int    `json:"size"`
	IsRunning   bool   `json:"isRunning"`
	Chksum      string `json:"chksum,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	InstallDate string `json:"installDate"`
	Description string `json:"description,omitempty"`
	Instruction string `json:"instruction,omitempty"`
}

type Apps struct {
	ResponseMessage
	Plugins []App `json:"plugins"`
}
