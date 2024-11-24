package utils

import (
	"encoding/json"
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
