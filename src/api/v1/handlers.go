package v1

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
)

func FindSimilarsWrapper(cfg *internal.Config, pgRepo *repositories.PgRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("image")
		if err != nil {
			http.Error(w, "Failed get image file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		nSimilar, err := strconv.Atoi(r.FormValue("n_similar"))
		if err != nil {
			slog.Warn("Failed to retrieve n_similar value. Using default", slog.Any("err_msg", err))
		}

		slog.Debug(
			"File info",
			slog.String("filename", handler.Filename),
			slog.Int64("fileSize", handler.Size),
		)

		tempFilePath := fmt.Sprintf("tmp/%s", handler.Filename)

		tempFile, err := os.Create(tempFilePath)
		if err != nil {
			http.Error(w, "Failed to create tmp file", http.StatusInternalServerError)
		}
		defer tempFile.Close()

		_, err = io.Copy(tempFile, file)
		if err != nil {
			http.Error(w, "Failed to sabe file", http.StatusInternalServerError)
		}
		targetEmbedding, err := internal.GetEmbedding([]string{tempFilePath}, cfg)
		if err != nil {
			slog.Error("Failed to get target embedding", slog.Any("err_msg", err))
		}
		similarCards := pgRepo.FindSimilarCards(targetEmbedding[0], nSimilar)
		slog.Debug(
			"Similar cards retrieved",
			slog.Any("cards", similarCards),
		)

		os.Remove(tempFilePath)
		if err != nil {
			slog.Error("Could not remove temp file", slog.Any("err_msg", err))
		}

		res, err := json.Marshal(FindSimilarsResponse{Cards: similarCards})
		if err != nil {
			http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		}
		_, err = w.Write(res)
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
		}
	}
}
