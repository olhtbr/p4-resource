package models

import "strconv"

type Version struct {
	Changelist string `json:"string"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type CheckResponse []Version

type Source struct {
	Server   Server   `json:"server"`
	User     string   `json:"string"`
	Ticket   string   `json:"string"`
	Filespec Filespec `json:"filespec"`
}

type Server struct {
	Protocol string `json:"string"`
	Host     string `json:"string"`
	Port     uint16 `json:"number"`
}

type Filespec struct {
	Depot  string `json:"string"`
	Stream string `json:"string"`
	Path   string `json:"string"`
}

func (s Server) String() (url string) {
	if s.Protocol != "" {
		url = s.Protocol + ":"
	}
	url += s.Host
	if s.Port != 0 {
		url += ":" + strconv.FormatUint(uint64(s.Port), 10)
	}
	return
}

func (f Filespec) String() string {
	return "//" + f.Depot + "/" + f.Stream + "/" + f.Path
}
