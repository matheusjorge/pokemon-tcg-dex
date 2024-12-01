package models

type EmbeddingResponse struct {
	Embeddings [][]float32 `json:"embeddings"`
}

type EmbeddingRequest struct {
	Filenames string `json:"filenames"`
	Model     string `json:"model"`
}
