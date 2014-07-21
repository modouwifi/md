package api

type App struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Icon        string `json:"icon"`
	Source      string `json:"source"`
	Version     string `json:"size"`
	Size        int    `json:"size"`
	Chksum      string `json:"chksum"`
	ReleaseDate string `json:"releaseDate"`
	InstallDate string `json:"installDate"`
	Description string `json:"description"`
	Instruction string `json:"instruction"`
}

type Apps struct {
	ResponseMessage
	Plugins []App `json:"plugins"`
}
