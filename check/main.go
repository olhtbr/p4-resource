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

	driver := driver.PerforceDriver{}
	err = driver.Login(request.Source.Server, request.Source.User, request.Source.Password)
	if err != nil {
		log.Fatalln(err)
	}

	if request.Version.Changelist == "" {
		cl, err := driver.GetLatestChangelist(request.Source.Filespec)
		if err != nil {
			log.Fatalln(err)
		}

		response = models.CheckResponse{
			{Changelist: cl},
		}
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
