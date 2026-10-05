package parsing

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

var ErrFileNotFound = errors.New("file not found")
var ErrParsingFile = errors.New("various file extensions: .json и .jso")

func ParseFile(path string) (map[string]any, error) {
	var data map[string]any

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrFileNotFound)
	}

	if err := json.Unmarshal(fileData, &data); err != nil {

		return nil, fmt.Errorf("%w", err)
	}

	return data, nil

}
