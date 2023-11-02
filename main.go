package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/veraison/corim/comid"
)

var sampleText = `
{
      "name": "Acme Ltd.",
      "regid": "https://acme.example",
      "roles": [
        "tagCreator",
	"owner"
      ]
}
`

func main() {
	comid.RegisterRole(4, "owner")

	var entity comid.Entity

	if err := json.Unmarshal([]byte(sampleText), &entity); err != nil {
		log.Fatalf("ERROR: %s", err.Error())
	}

	if err := entity.Valid(); err != nil {
		log.Fatalf("failed to validate: %s", err.Error())
	} else {
		fmt.Println("validation succeeded")
	}

	fmt.Println("roles:")
	for _, role := range entity.Roles {
		fmt.Printf("\t%s\n", role.String())
	}
}
