package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Request interface {
	Setup([]byte) error
}

type Response interface {
	Clear()
}

type Version struct {
	Changelist string `json:"changelist"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type CheckResponse []Version

type InRequest struct {
	Source  Source  `json:"source"`
	Version Version `json:"version"`
}

type InResponse struct {
	Version Version `json:"version"`
}

type Source struct {
	Server   Server   `json:"server"`
	User     string   `json:"user"`
	Password string   `json:"password"`
	Filespec Filespec `json:"filespec"`
}

type Server struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
}

type Filespec struct {
	Depot  string `json:"depot"`
	Stream string `json:"stream"`
	Path   string `json:"path"`
}

func (s Server) String() string {
	url := ""
	if s.Protocol != "" {
		url += s.Protocol + ":"
	}
	url += s.Host
	if s.Port != 0 {
		url += ":" + strconv.FormatUint(uint64(s.Port), 10)
	} else {
		url += ":1666"
	}
	return fmt.Sprintf("%s", url)
}

func (f Filespec) String() string {
	return "//" + f.Depot + "/" + f.Stream + "/" + f.Path
}

func (r *CheckRequest) Setup(jsonBlob []byte) error {
	err := json.Unmarshal(jsonBlob, &r)
	return err
}

func (r *InRequest) Setup(jsonBlob []byte) error {
	err := json.Unmarshal(jsonBlob, &r)
	return err
}

func (r *CheckResponse) Clear() {
	*r = CheckResponse{}
}

func (r *InResponse) Clear() {
	*r = InResponse{}
}
