package parsers

import (
	"fmt"
	"os"
	"errors"
	//"encoding/json"
	"github.com/go-yaml/yaml"
)

var ErrFileNotFound = errors.New("не найден файл")
var ErrParsingFile = errors.New("ошибка парсинга файла")

func ParseFile(path string) (map[string]any, error) {
	var data map[string]any

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrFileNotFound)
	}

		if err := yaml.Unmarshal(fileData, &data); err != nil {
		
		return nil, fmt.Errorf("%w", ErrParsingFile)
	}

	return data, nil

}