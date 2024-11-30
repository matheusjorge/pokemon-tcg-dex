package cmd

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/utils"
	"github.com/spf13/cobra"
)

func FetchImages(imgsURLs []string, cfg *internal.Config) {
	var wg sync.WaitGroup
	queue := make(chan string, len(imgsURLs))
	slog.Debug("Start Downloading Images")

	for _, url := range imgsURLs {
		queue <- url
	}
	close(queue)

	slog.Debug("Created url queue")

	wg.Add(cfg.ImageDownloaderWorkers)
	for i := 0; i < cfg.ImageDownloaderWorkers; i++ {
		slog.Debug("Starting worker", slog.Int("id", i))
		go func(q <-chan string, wg *sync.WaitGroup) {
			defer wg.Done()

			for url := range q {
				slog.Debug("Downloading Image", slog.String("url", url))

				filepath := fmt.Sprintf("%s/images/%s", cfg.DataPath, utils.ImageURLToFilename(url))
				utils.FetchResource(url, filepath)
			}

		}(queue, &wg)
	}

	wg.Wait()
	slog.Debug("Finished Downloading Images")
}

func DownloadImages(cfg *internal.Config, pgRepo *repositories.PgRepo) cobra.Command {
	return cobra.Command{
		Use:   "download-images",
		Short: "Download card images to local directory",
		Run: func(cmd *cobra.Command, args []string) {

			imageURLs, err := pgRepo.GetImageURLs()
			if err != nil {
				slog.Error("Failed to get urls", slog.Any("err_msg", err))
			}
			slog.Debug("Got Image URLs", slog.Any("len", len(imageURLs)))

			FetchImages(imageURLs, cfg)
		},
	}
}
