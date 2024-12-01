package cmd

import (
	"log/slog"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/spf13/cobra"
)

func FindSimilars(cfg *internal.Config, pgRepo *repositories.PgRepo) cobra.Command {
	var filename string
	Cmd := cobra.Command{
		Use:   "find-similars",
		Short: "Search cards whose image are similar to the image provided",
	}

	Cmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the file to search similar cards")
	Cmd.Run = func(cmd *cobra.Command, args []string) {

		targetEmbedding, err := internal.GetEmbedding([]string{filename}, cfg)
		if err != nil {
			slog.Error("Failed to get target embedding", slog.Any("err_msg", err))
		}
		similarCards := pgRepo.FindSimilarCards(targetEmbedding[0], 5)
		slog.Debug(
			"Similar cards retrieved",
			slog.Any("cards", similarCards),
		)
	}

	return Cmd
}
