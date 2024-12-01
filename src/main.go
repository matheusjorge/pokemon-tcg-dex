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
		cmd.InitAPI(config, pgRepo),
	}
	cmd.AddCommands(commands)

	if err := cmd.Execute(); err != nil {
		slog.Error("Failed to run command", slog.Any("err_msg", err))
		os.Exit(1)
	}

}
