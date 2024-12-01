package cmd

import (
	"fmt"
	"log/slog"
	"net/http"

	v1 "github.com/matheusjorge/pokemon-tcg-dex/src/api/v1"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
	"github.com/spf13/cobra"
)

func InitAPI(cfg *internal.Config, pgRepo *repositories.PgRepo) cobra.Command {
	return cobra.Command{
		Use:   "api",
		Short: "Starts API server",
		Run: func(cmd *cobra.Command, args []string) {
			router := v1.InitRoutes(cfg, pgRepo)
			server := &http.Server{
				Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
				Handler: router,
			}

			slog.Debug("Starting Server", slog.Int64("port", int64(cfg.ServerPort)))
			err := server.ListenAndServe()
			if err != nil {
				slog.Error("Api failed", slog.Any("err_msg", err))
			}
		},
	}

}
