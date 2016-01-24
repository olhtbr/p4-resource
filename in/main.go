package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/olhtbr/p4-resource/models"
)

func main() {
	var request models.InRequest

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalln(err)
	}

	if request.Version.Changelist == "" {
		log.Fatalln("Requested version is empty")
	}
}
