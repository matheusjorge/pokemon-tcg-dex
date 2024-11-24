package tools

import (
	"fmt"
	"log"
	"sync"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/matheusjorge/pokemon-tcg-dex/src/models"
)

func LoadSet(cfg internal.Config, set string) []models.Card {
	log.Printf("Loading set %s", set)
	var cardsJson []models.CardJson
	filepath := fmt.Sprintf("%s/cards/%s.json", cfg.DataPath, set)
	utils.LoadJson(filepath, &cardsJson)

	var cardsPg []models.Card
	for _, c := range cardsJson {
		cardPg, err := models.FromJsonToPg(&c)
		if err != nil {
			log.Printf("Error creating Pg data: %s", err)
			continue
		}
		(*cardPg).SetId = set
		cardsPg = append(cardsPg, *cardPg)
	}

	return cardsPg
}

func LoadAllSets(cfg internal.Config) []models.Card {
	cardsChannel := make(chan []models.Card)
	var wg sync.WaitGroup

	var dataConfig map[string][]string
	utils.LoadJson(cfg.DataConfigFile, &dataConfig)
	for _, set := range dataConfig["cards_sets"] {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := LoadSet(cfg, set)
			cardsChannel <- c
		}()
	}

	go func() {
		wg.Wait()
		close(cardsChannel)
	}()

	var cards []models.Card
	for cardSets := range cardsChannel {
		cards = append(cards, cardSets...)
	}

	return cards
}
