package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JsonMarshal(data interface{}) string {
	byteJson, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(byteJson)
}

func JsonUnmarshal[T any](data string, placeholderVar *T) {
	byteData := []byte(data)
	_ = json.Unmarshal(byteData, placeholderVar)
}

func ImageURLToFilename(url string) string {
	parts := strings.Split(url, "/")
	filename := fmt.Sprintf("%s-%s", parts[len(parts)-2], parts[len(parts)-1])

	return filename
}
func ImageURLToCardIDd(url string) string {
	parts := strings.Split(url, "/")
	cardID := fmt.Sprintf("%s-%s", parts[len(parts)-2], parts[len(parts)-1])
	cardID = strings.Replace(cardID, ".png", "", 1)

	return cardID
}
