package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
)

func main() {

	config, err := internal.LoadConfigs()
	if err != nil {
		log.Fatal("Could no load configs")
	}

	// Create Request
	getCardsURL := fmt.Sprintf("%s/sv8/1.png", config.PokemonTCGImagesURL)
	req, err := http.NewRequest("GET", getCardsURL, nil)
	if err != nil {
		log.Fatal("Could not create request")
	}

	req.Header.Set("x-api-key", config.PokemonTCGAPIKey)

	// Make Request
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal("Request not made")
	}

	defer resp.Body.Close()

	file, err := os.Create("sv8-1.png")
	if err != nil {
		log.Fatal("Could not create file")
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal("Error reading body")
	}

}
