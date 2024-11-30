package cmd

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/models"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/spf13/cobra"
)

func LoadSet(cfg *internal.Config, set string) []models.Card {
	slog.Debug("Loading set", slog.String("set_name", set))
	var cardsJson []models.CardJson
	filepath := fmt.Sprintf("%s/cards/%s.json", cfg.DataPath, set)
	err := utils.LoadJson(filepath, &cardsJson)
	if err != nil {
		return []models.Card{}
	}

	var cardsPg []models.Card
	for _, c := range cardsJson {
		cardPg, err := models.FromJsonToPg(&c)
		if err != nil {
			slog.Error("Error creating Pg data", slog.Any("err_msg", err))
			continue
		}
		(*cardPg).SetId = set
		cardsPg = append(cardsPg, *cardPg)
	}

	return cardsPg
}

func LoadAllSets(cfg *internal.Config) []models.Card {
	cardsChannel := make(chan []models.Card)
	var wg sync.WaitGroup

	var dataConfig map[string][]string
	err := utils.LoadJson(cfg.DataConfigFile, &dataConfig)
	if err != nil {
		return []models.Card{}
	}
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

	slog.Debug("Loaded cards", slog.Int("cards_loaded", len(cards)))
	return cards
}

func InsertCardsData(cfg *internal.Config, pgRepo *repositories.PgRepo) cobra.Command {
	return cobra.Command{
		Use:   "insert-cards",
		Short: "Insert cards data into database from json files",
		Run: func(cmd *cobra.Command, args []string) {
			allCards := LoadAllSets(cfg)
			pgRepo.InsertManyCards(allCards)
		},
	}
}
