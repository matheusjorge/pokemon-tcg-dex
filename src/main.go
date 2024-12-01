package main

import (
	"log/slog"
	"os"

	"github.com/matheusjorge/pokemon-tcg-dex/src/cmd"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/spf13/cobra"
)

func main() {

	internal.SetupLogger()
	config, err := internal.LoadConfigs()
	if err != nil {
		slog.Error("Could no load configs", slog.Any("err_msg", err))

	}
	pgRepo, err := repositories.PgConnect(config.PostgresURL)
	if err != nil {
		slog.Error("Could not connect to postgres", slog.Any("err_msg", err))
	}

	commands := []cobra.Command{
		cmd.DownloadSets(config),
		cmd.InsertCardsData(config, pgRepo),
		cmd.DownloadImages(config, pgRepo),
		cmd.InsertImageEmbeddings(config, pgRepo),
		cmd.FindSimilars(config, pgRepo),
	}
	cmd.AddCommands(commands)

	if err := cmd.Execute(); err != nil {
		slog.Error("Failed to run command", slog.Any("err_msg", err))
		os.Exit(1)
	}

	// var imagePaths []string
	// for i, url := range imageURLs {
	// 	imagePaths = append(imagePaths, fmt.Sprintf("%s/images/%s", config.DataPath, utils.ImageURLToFilename(url)))
	//
	// 	if i > 4 {
	// 		break
	// 	}
	// }
	//
	// slog.Debug("Get image paths", slog.Any("paths", imagePaths))
	//
	// embeddings, err := internal.GetEmbedding(imagePaths, *config)
	// if err != nil {
	// 	slog.Error("Failed to get embeddings", slog.Any("err_msg", err))
	// }
	//
	// println(embeddings[0])

}
