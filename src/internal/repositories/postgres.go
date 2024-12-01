package repositories

import (
	"context"
	"log"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/models"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/pgvector/pgvector-go"
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
	var rows [][]interface{}
	tableName := "cards"
	columns := []string{
		"id",
		"name",
		"set_id",
		"number",
		"artist",
		"rarity",
		"national_pokedex_number",
		"image_small",
		"image_large",
		"supertype",
		"subtypes",
		"level",
		"hp",
		"types",
		"evolves_from",
		"evolves_to",
		"rules",
		"ancient_trait",
		"abilities",
		"attacks",
		"weaknesses",
		"resistances",
		"retreat_cost",
		"converted_retreat_cost",
	}
	for _, card := range cards {
		rows = append(rows, []interface{}{
			card.Id,
			card.Name,
			card.SetId,
			card.Number,
			card.Artist,
			card.Rarity,
			card.NationalPokedexNumber,
			card.Images.Small,
			card.Images.Large,
			card.Supertype,
			card.Subtypes,
			card.Level,
			card.HP,
			card.Types,
			card.EvolvesFrom,
			card.EvolvesTo,
			card.Rules,
			utils.JsonMarshal(card.AncientTrait),
			utils.JsonMarshal(card.Abilities),
			utils.JsonMarshal(card.Attacks),
			utils.JsonMarshal(card.Weaknessess),
			utils.JsonMarshal(card.Resistances),
			card.RetreatCost,
			card.ConvertedRetreatCost,
		})
	}
	count, err := r.Conn.CopyFrom(context.Background(), pgx.Identifier{tableName}, columns, pgx.CopyFromRows(rows))
	if err != nil {
		slog.Error("Failed to insert cards", slog.Any("err_msg", err))
	}
	slog.Debug("Inserted rows", slog.Int64("count", count))
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
			slog.Error("Failed to parse card", slog.Any("err_msg", err))
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

func (r *PgRepo) GetImageURLs() ([]string, error) {
	query := `
	SELECT image_small FROM cards
	`

	rows, err := r.Conn.Query(context.Background(), query)
	if err != nil {
		return []string{}, nil
	}

	imageURLs := []string{}
	for rows.Next() {
		var url string
		err = rows.Scan(&url)
		if err != nil {
			slog.Error("Failed to get url", slog.Any("err_msg", err))
		}

		imageURLs = append(imageURLs, url)
	}

	return imageURLs, nil

}

func (r *PgRepo) InsertEmbeddings(ids []string, embeddings [][]float32) {
	query := `
	UPDATE cards SET image_embedding = $1 WHERE id = $2
	`

	for i := 0; i < len(ids); i++ {
		_, err := r.Conn.Exec(context.Background(), query, pgvector.NewVector(embeddings[i]), ids[i])
		if err != nil {
			slog.Error("Failed to insert embedding", slog.Any("err_msg", err))
			continue
		}
		slog.Debug("Inserted embedding", slog.String("cardId", ids[i]))
	}

}

func (r *PgRepo) FindSimilarCards(embedding []float32) []models.SimilarSearchResponse {
	query := `
	SELECT
		 id
		,image_small
		,1 - (image_embedding <=> $1) AS similarity
	FROM
		cards
	WHERE
		1=1
		--AND set_id = 'sv8'
	ORDER BY
		image_embedding <=> $1
	LIMIT 5
	`
	similarCards := []models.SimilarSearchResponse{}
	rows, err := r.Conn.Query(context.Background(), query, pgvector.NewVector(embedding))

	if err != nil {
		slog.Error("Failed to query similar cards", slog.Any("err_msg", err))
		return similarCards
	}

	for rows.Next() {
		var response models.SimilarSearchResponse
		err = rows.Scan(
			&response.Id,
			&response.Image,
			&response.Similarity,
		)

		if err != nil {
			slog.Error("Failed to parse similar card", slog.Any("err_msg", err))
		}

		similarCards = append(similarCards, response)

	}

	return similarCards[:5]
}

func (r *PgRepo) ComputeDistancetoCard(embedding []float32) float32 {
	query := `
	SELECT
		1 - (image_embedding <=> $1)
	FROM
		cards
	WHERE
		1=1
		AND id = 'sv8-123'
	`
	row := r.Conn.QueryRow(context.Background(), query, pgvector.NewVector(embedding))
	var similarity float32

	err := row.Scan(&similarity)
	if err != nil {
		slog.Error("Failed to parse similarity")
	}

	return similarity
}
