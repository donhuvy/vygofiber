package sample4

import (
	"encoding/json"
	"fmt"
	"log"
)

// The same json tags will be used to encode data into JSON
type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func Main4() {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))

	manifestJson, _ := json.MarshalIndent(data, "", "  ")

	log.Println(string(manifestJson))

}

func EasyPrint(data interface{}) {
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}
	data2, _ := json.Marshal(pigeon)
	manifestJson, _ := json.MarshalIndent(data2, "", "  ")

	log.Println(string(manifestJson))
}
