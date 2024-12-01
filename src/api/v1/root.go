package v1

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/matheusjorge/pokemon-tcg-dex/src/internal"
	"github.com/matheusjorge/pokemon-tcg-dex/src/internal/repositories"
)

func InitRoutes(cfg *internal.Config, pgRepo *repositories.PgRepo) *http.ServeMux {
	mux := http.ServeMux{}
	err := os.MkdirAll("./temp", os.ModePerm)
	if err != nil {
		slog.Error("Failed to create temp folder", slog.Any("err_msg", err))
		return &mux
	}

	mux.HandleFunc("POST /v1/cards/similar", FindSimilarsWrapper(cfg, pgRepo))

	return &mux
}
