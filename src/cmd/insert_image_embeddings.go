package cmd

import (
	"fmt"
	"log/slog"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/spf13/cobra"
)

func InsertImageEmbeddings(cfg *internal.Config, pgRepo *repositories.PgRepo) cobra.Command {
	return cobra.Command{
		Use:   "insert-embeddings",
		Short: "Get embedding for all images and insert them into the database",
		Run: func(cmd *cobra.Command, args []string) {

			imageURLs, err := pgRepo.GetImageURLs()
			if err != nil {
				slog.Error("Failed to get urls", slog.Any("err_msg", err))
			}
			var imagePaths []string
			var cardIDs []string
			for _, url := range imageURLs {
				imagePaths = append(imagePaths, fmt.Sprintf("%s/images/%s", cfg.DataPath, utils.ImageURLToFilename(url)))
				cardIDs = append(cardIDs, utils.ImageURLToCardIDd(url))
			}

			slog.Debug("Parsed image paths", slog.Any("paths", imagePaths), slog.Any("ids", cardIDs))

			chunkSize := 100
			start_idx := 0
			for {
				end_idx := min(len(cardIDs)-1, start_idx+chunkSize)
				embeddings, err := internal.GetEmbedding(imagePaths[start_idx:end_idx], cfg)
				if err != nil {
					slog.Error("Failed to get embeddings", slog.Any("err_msg", err))
				}
				pgRepo.InsertEmbeddings(cardIDs[start_idx:end_idx], embeddings)
				start_idx = end_idx
				if end_idx == len(cardIDs)-1 {
					break
				}
			}

		},
	}
}
