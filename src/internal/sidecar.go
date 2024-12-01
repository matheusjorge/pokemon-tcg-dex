package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/models"
)

func GetEmbedding(filenames []string, cfg *Config) ([][]float32, error) {
	payload := models.EmbeddingRequest{
		Filenames: strings.Join(filenames, ","),
		Model:     "mobilenetv3_large_100",
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		slog.Error("Failed to parse request", slog.Any("err_msg", err))
		return [][]float32{}, nil
	}
	payloadBytes := []byte(payloadJson)
	sidecarURL := fmt.Sprintf("http://0.0.0.0:%d/v1/images/embeddings", cfg.SidecarPort)

	req, err := http.NewRequest("POST", sidecarURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		slog.Error("Failed to create request", slog.Any("err_msg", err))
		return [][]float32{}, err
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		slog.Error("Failed to make request", slog.Any("err_msg", err))
		return [][]float32{}, err
	}

	defer res.Body.Close()

	var embeddings models.EmbeddingResponse
	err = json.NewDecoder(res.Body).Decode(&embeddings)
	if err != nil {
		slog.Error("Failed decode response", slog.Any("err_msg", err))
		return [][]float32{}, err
	}

	return embeddings.Embeddings, nil

}
