package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/olhtbr/p4-resource/models"
)

func main() {
	var request models.CheckRequest
	var response models.CheckResponse

	err := json.NewDecoder(os.Stdin).Decode(&request)
	if err != nil {
		log.Fatalln(err)
	}

	if request.Version.Changelist == "" {
		// TODO Get latest changelist from the driver
		response = models.CheckResponse{
			{Changelist: "123456"},
		}
	}

	json.NewEncoder(os.Stdout).Encode(response)
}
