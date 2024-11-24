package tools

import (
	// "fmt"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
)

func FetchResource(url string, filepath string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Could not create request: %s", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Request not made: %s", err)
	}

	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		log.Printf("Could not create file: %s", err)
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Printf("Error reading body:  %s", err)
	}
}

func FetchCardsData(cfg internal.Config) {

	baseUrl := fmt.Sprintf("%s/cards/en", cfg.RemoteDataURL)
	var dataConfig map[string][]string
	utils.LoadJson(cfg.DataConfigFile, &dataConfig)

	var wg sync.WaitGroup

	for _, set := range dataConfig["cards_sets"] {
		url := fmt.Sprintf("%s/%s.json", baseUrl, set)
		filepath := fmt.Sprintf("%s/cards/%s.json", cfg.DataPath, set)
		wg.Add(1)
		go func() {
			defer wg.Done()
			FetchResource(url, filepath)
		}()
	}
	wg.Wait()
}
