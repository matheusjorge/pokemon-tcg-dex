package models

type SimilarSearchResponse struct {
	Id         string  `json:"id"`
	Image      string  `json:"image"`
	Similarity float32 `json:"similarity"`
}
