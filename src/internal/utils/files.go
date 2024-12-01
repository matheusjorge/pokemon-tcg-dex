package utils

import (
	"encoding/json"
	"io"
	"log/slog"
	"os"
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
