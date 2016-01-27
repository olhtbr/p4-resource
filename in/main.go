package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/olhtbr/p4-resource/driver"
	"github.com/olhtbr/p4-resource/models"
)

func main() {
	var request models.InRequest
	cl := request.Version.Changelist

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalln(err)
	}

	driver := new(driver.Driver)
	err = driver.Login(request.Source.Server, request.Source.User, request.Source.Password)
	if err != nil {
		log.Fatalln(err)
	}

	if cl == "" {
		log.Fatalln("Requested version is empty")
	} else {
		exists, err := driver.ChangelistExists(cl)
		if err != nil {
			log.Fatalln(err)
		}

		if !exists {
			log.Fatalln("Requested version (" + cl + ")does not exist")
		}
	}
}
