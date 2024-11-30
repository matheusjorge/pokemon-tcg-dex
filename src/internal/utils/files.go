package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

func LoadJson[T any](filepath string, placeholderVar *T) error {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		slog.Error("Failed to open json file", slog.Any("err_msg", err))
		return err
	}

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		slog.Error("Failed to open json file", slog.Any("err_msg", err))
		return err
	}

	err = json.Unmarshal(byteValue, placeholderVar)
	if err != nil {
		slog.Error("Failed to open json file", slog.Any("err_msg", err))
		return err
	}
	return nil
}

func ImageURLToFilename(url string) string {
	parts := strings.Split(url, "/")
	filename := fmt.Sprintf("%s-%s", parts[len(parts)-2], parts[len(parts)-1])

	return filename
}
