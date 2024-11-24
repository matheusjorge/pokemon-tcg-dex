package main

import (
	"log"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/repositories"
	"github.com/matheusjorge/pokemon-tcg-dex/src/tools"
)

func main() {

	config, err := internal.LoadConfigs()
	if err != nil {
		log.Fatal("Could no load configs")

	}
	pgRepo, err := repositories.PgConnect(config.PostgresURL)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	// pgRepo.InsertManyCards(cardsPg)
	// cardsNew := pgRepo.FetchAllCards()
	// log.Print(cardsNew[162])

	// tools.FetchCardsData(*config)
	allCards := tools.LoadAllSets(*config)
	log.Println(len(allCards))

	pgRepo.InsertManyCards(allCards)
}
