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

	driver := &driver.PerforceDriver{}
	err = driver.Login(request.Source.Server, request.Source.User, request.Source.Password)
	if err != nil {
		log.Fatalln(err)
	}

	if request.Version.Changelist == "" {
		response = getResponseWithLatest(driver, request.Source.Filespec)
	} else {
		cls, err := driver.GetChangelistsNewerThan(request.Version.Changelist, request.Source.Filespec)
		if err != nil {
			// Fallback to latest changelist
			response = getResponseWithLatest(driver, request.Source.Filespec)
		} else {
			response = models.CheckResponse{}
			for _, cl := range cls {
				response = append(response, models.Version{
					Changelist: cl,
				})
			}
		}
	}

	json.NewEncoder(os.Stdout).Encode(response)
}

func getResponseWithLatest(driver driver.Driver, filespec models.Filespec) (response models.CheckResponse) {
	cl, err := driver.GetLatestChangelist(filespec)
	if err != nil {
		log.Fatalln(err)
	}

	response = models.CheckResponse{
		{Changelist: cl},
	}

	return
}
