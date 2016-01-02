package models

type Version struct {
	Changelist string `json:"string"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type CheckResponse []Version

type Source struct {
	Port     Port     `json:"port"`
	User     string   `json:"string"`
	Ticket   string   `json:"string"`
	Filespec Filespec `json:"filespec"`
}

type Port struct {
	Protocol string `json:"string"`
	Host     string `json:"string"`
	Port     uint16 `json:"number"`
}

type Filespec struct {
	Depot  string `json:"string"`
	Stream string `json:"string"`
	Path   string `json:"string"`
}
