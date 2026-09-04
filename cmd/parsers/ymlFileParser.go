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
var ErrEmptyFile = errors.New("файл пуст")

func ParseFile(path string) (map[string]any, error) {
	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrFileNotFound, path)
	}

	// Проверяем, что файл не пустой
	if len(fileData) == 0 {
		return nil, fmt.Errorf("%w: %s", ErrEmptyFile, path)
	}

	// Создаем map для результата
	result := make(map[string]any)
	
	// Парсим YAML напрямую в map[string]any
	// Для этого нужно использовать промежуточный шаг
	var rawData map[interface{}]interface{}
	if err := yaml.Unmarshal(fileData, &rawData); err != nil {
		return nil, fmt.Errorf("%w: %s - %v", ErrParsingFile, path, err)
	}

	// Если данные пустые, возвращаем пустую map
	if len(rawData) == 0 {
		return result, nil
	}

	// Конвертируем в map[string]any
	for key, value := range rawData {
		if keyStr, ok := key.(string); ok {
			result[keyStr] = convertValue(value)
		} else {
			// Если ключ не строка, конвертируем его в строку
			result[fmt.Sprintf("%v", key)] = convertValue(value)
		}
	}

	return result, nil
}


func convertValue(value interface{}) interface{} {
	switch v := value.(type) {
	case map[interface{}]interface{}:
		converted := make(map[string]any)
		for key, val := range v {
			if keyStr, ok := key.(string); ok {
				converted[keyStr] = convertValue(val)
			} else {
				converted[fmt.Sprintf("%v", key)] = convertValue(val)
			}
		}
		return converted
	case []interface{}:
		for i, val := range v {
			v[i] = convertValue(val)
		}
		return v
	default:
		return v
	}
}