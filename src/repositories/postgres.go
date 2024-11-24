package repositories

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/matheusjorge/pokemon-tcg-dex/src/models"
)

type PgRepo struct {
	Conn *pgx.Conn
}

func PgConnect(databaseURL string) (*PgRepo, error) {
	conn, err := pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		return nil, err
	}

	return &PgRepo{
		Conn: conn,
	}, nil
}

func (r *PgRepo) InsertCard(card models.Card) error {
	query := `
  INSERT INTO cards ( 
    id,
    name,
    set_id,
    number,
    artist,
    rarity,
    national_pokedex_number,
    image_small,
    image_large,
 supertype,
    subtypes,
    level,
    hp,
    types,
    evolves_from,
    evolves_to,
    rules,
    ancient_trait,
    abilities,
    attacks,
    weaknesses,
    resistances,
    retreat_cost,
    converted_retreat_cost
  ) VALUES (
    @id,
    @name,
    @set_id,
    @number,
    @artist,
    @rarity,
    @national_pokedex_number,
    @image_small,
    @image_large,
    @supertype,
    @subtypes,
    @level,
    @hp,
    @types,
    @evolves_from,
    @evolves_to,
    @rules,
    @ancient_trait,
    @abilities,
    @attacks,
    @weaknesses,
    @resistances,
    @retreat_cost,
    @converted_retreat_cost
  )
  `

	// bytesRules, _ := json.Marshal(card.Rules)
	args := pgx.NamedArgs{
		"id":                      card.Id,
		"name":                    card.Name,
		"set_id":                  card.SetId,
		"number":                  card.Number,
		"artist":                  card.Artist,
		"rarity":                  card.Rarity,
		"national_pokedex_number": card.NationalPokedexNumber,
		"image_small":             card.Images.Small,
		"image_large":             card.Images.Large,
		"supertype":               card.Supertype,
		"subtypes":                card.Subtypes,
		"level":                   card.Level,
		"hp":                      card.HP,
		"types":                   card.Types,
		"evolves_from":            card.EvolvesFrom,
		"evolves_to":              card.EvolvesTo,
		"rules":                   card.Rules,
		"ancient_trait":           utils.JsonMarshal(card.AncientTrait),
		"abilities":               utils.JsonMarshal(card.Abilities),
		"attacks":                 utils.JsonMarshal(card.Attacks),
		"weaknesses":              utils.JsonMarshal(card.Weaknessess),
		"resistances":             utils.JsonMarshal(card.Resistances),
		"retreat_cost":            card.RetreatCost,
		"converted_retreat_cost":  card.ConvertedRetreatCost,
	}

	_, err := r.Conn.Exec(context.Background(), query, args)
	if err != nil {
		log.Printf("Error inserting into pg card %s: %s", card.Id, err)
		return err
	}

	return nil
}

func (r *PgRepo) InsertManyCards(cards []models.Card) {
	for i, card := range cards {
		err := r.InsertCard(card)
		if err != nil {
			log.Printf("Failed to insert card no %d", i)
		}
	}
}

func (r *PgRepo) FetchAllCards() []models.Card {
	query := `
  SELECT * FROM cards
  `
	rows, _ := r.Conn.Query(context.Background(), query)
	defer rows.Close()

	var cards []models.Card
	for rows.Next() {
		var card models.Card
		var images models.Images
		var ancientTrait string
		var abilities string
		var attacks string
		var weaknesses string
		var resistances string

		err := rows.Scan(
			&card.Id,
			&card.Name,
			&card.SetId,
			&card.Number,
			&card.Artist,
			&card.Rarity,
			&card.NationalPokedexNumber,
			&images.Small,
			&images.Large,
			&card.Supertype,
			&card.Subtypes,
			&card.Level,
			&card.HP,
			&card.Types,
			&card.EvolvesFrom,
			&card.EvolvesTo,
			&card.Rules,
			&ancientTrait,
			&abilities,
			&attacks,
			&weaknesses,
			&resistances,
			&card.RetreatCost,
			&card.ConvertedRetreatCost,
		)
		if err != nil {
			log.Printf("Failed to parse card: %s", err)
			continue
		}

		card.Images = images
		utils.JsonUnmarshal(ancientTrait, &card.AncientTrait)
		utils.JsonUnmarshal(abilities, &card.Abilities)
		utils.JsonUnmarshal(attacks, &card.Attacks)
		utils.JsonUnmarshal(weaknesses, &card.Weaknessess)
		utils.JsonUnmarshal(resistances, &card.Resistances)

		cards = append(cards, card)

	}

	return cards
}
