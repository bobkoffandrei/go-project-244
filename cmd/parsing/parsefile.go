package parsing

import (
	"fmt"
	"os"
	//"errors"
	"encoding/json"
)



func ParseFile(path string) (map[string]any, error) {
	var data map[string]any

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла: %w", err)
	}

		if err := json.Unmarshal(fileData, &data); err != nil {
		
		return nil, fmt.Errorf("Ошибка парсинга файла: %w", err)
	}

	return data, nil

}