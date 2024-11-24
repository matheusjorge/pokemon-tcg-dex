package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func LoadJson[T any](filepath string, placeholderVar *T) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal("Could not open json file")
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Could not read json file")
	}

	err = json.Unmarshal(byteValue, placeholderVar)
	if err != nil {
		log.Fatalf("Error parsing json %s: %s", filepath, err)
	}
}
