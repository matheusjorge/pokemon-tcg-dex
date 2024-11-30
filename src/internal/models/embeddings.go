package models

type EmbeddingResponse struct {
	Embeddings [][]float64 `json:"embedding"`
}

type EmbeddingRequest struct {
	Filenames string `json:"filenames"`
	Model     string `json:"model"`
}
