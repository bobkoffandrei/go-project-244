package parsing

import (
	"fmt"
	"os"
	"errors"
	"encoding/json"
)

var ErrFileNotFound = errors.New("не найден файл")
var ErrParsingFile = errors.New("ошибка парсинга файла")

func ParseFile(path string) (map[string]any, error) {
	var data map[string]any

	fileData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrFileNotFound)
	}

		if err := json.Unmarshal(fileData, &data); err != nil {
		
		return nil, fmt.Errorf("%w", ErrParsingFile)
	}

	// printElements(data, 0)

	return data, nil

	//return nil, nil

}
/*
func printElements(data any, level int) {
    indent := ""
    for i := 0; i < level; i++ {
        indent += "  " 
    }

    switch v := data.(type) {
    case map[string]any:

        for key, value := range v {
            fmt.Printf("%s%s:\n", indent, key) 
            printElements(value, level+1)      
        }
        
    case []any:

        for i, value := range v {
            fmt.Printf("%s[%d]:\n", indent, i)  
            printElements(value, level+1)     
        }
        
    default:
        fmt.Printf("%s%v\n", indent, v)
    }
}
*/