package v1

import "github.com/matheusjorge/pokemon-tcg-dex/src/internal/models"

type FindSimilarsRequest struct {
}

type FindSimilarsResponse struct {
	Cards []models.SimilarSearchResponse `json:"cards"`
}
