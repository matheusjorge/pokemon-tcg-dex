package utils

import (
	"io"
	"log/slog"
	"net/http"
	"os"
)

func FetchResource(url string, filepath string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Error("Could not create request", slog.Any("err_msg", err))
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		slog.Error("Request not made", slog.Any("err_msg", err))
	}

	defer resp.Body.Close()

	file, err := os.Create(filepath)
	if err != nil {
		slog.Error("Could not create file", slog.Any("err_msg", err))
	}

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		slog.Error("Error reading body", slog.Any("err_msg", err))
	}
}
