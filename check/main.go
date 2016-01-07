package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/olhtbr/p4-resource/driver"
	"github.com/olhtbr/p4-resource/models"
)

func main() {
	var request models.CheckRequest
	var response models.CheckResponse

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalln(err)
	}

	driver := driver.PerforceDriver{
		Server: request.Source.Server,
		User:   request.Source.User,
		Ticket: request.Source.Ticket,
	}

	if request.Version.Changelist == "" {
		response = models.CheckResponse{
			{Changelist: driver.GetLatestChangelist(request.Source.Filespec)},
		}
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
