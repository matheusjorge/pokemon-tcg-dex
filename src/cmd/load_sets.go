package cmd

import (
	"github.com/spf13/cobra"

	"fmt"
	"log/slog"
	"sync"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
)

func FetchCardsData(cfg *internal.Config) {

	baseUrl := fmt.Sprintf("%s/cards/en", cfg.RemoteDataURL)
	var dataConfig map[string][]string
	err := utils.LoadJson(cfg.DataConfigFile, &dataConfig)
	if err != nil {
		slog.Error(
			"Failed to load data config file. Aborting.",
			slog.Any("err_msg", err),
		)
	}

	var wg sync.WaitGroup

	for _, set := range dataConfig["cards_sets"] {
		url := fmt.Sprintf("%s/%s.json", baseUrl, set)
		filepath := fmt.Sprintf("%s/cards/%s.json", cfg.DataPath, set)
		wg.Add(1)
		go func() {
			defer wg.Done()
			utils.FetchResource(url, filepath)
		}()
	}
	wg.Wait()
}

func DownloadSets(cfg *internal.Config) cobra.Command {
	return cobra.Command{
		Use:   "download-sets",
		Short: "Fetch sets data from github repository",
		Run: func(cmd *cobra.Command, args []string) {
			FetchCardsData(cfg)
		},
	}
}
