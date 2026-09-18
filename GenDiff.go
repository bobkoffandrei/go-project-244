package code

import (
	"code/internal/formatters"
	"code/internal/parsers"
	"code/internal/parsing"
	"code/models"
	"fmt"
	"path/filepath"
	"reflect"
	"sort"
)

func GenDiff(file1, file2, format string) (string, error) {

	ext1 := filepath.Ext(file1)
	ext2 := filepath.Ext(file2)

	var fileMap1, fileMap2 map[string]any

	var err error

	if ext1 == ".json" && ext2 == ".json" {
		fileMap1, err = parsing.ParseFile(file1)
		if err != nil {
			return "", fmt.Errorf("parsing %s as json: %v", file1, err)
		}
		fileMap2, err = parsing.ParseFile(file2)
		if err != nil {
			return "", fmt.Errorf("parsing %s as json: %v", file2, err)
		}
	} else if (ext1 == ".yaml" && ext2 == ".yaml") || (ext1 == ".yml" && ext2 == ".yml") {
		fileMap1, err = parsers.ParseFile(file1)
		if err != nil {
			return "", fmt.Errorf("parsing %s as yaml: %v", file1, err)
		}
		fileMap2, err = parsers.ParseFile(file2)
		if err != nil {
			return "", fmt.Errorf("parsing %s as yaml: %v", file2, err)
		}
	} else if ext1 != ext2 {
		return "", fmt.Errorf("various file extensions: %s и %s", ext1, ext2)
	} else {
		return "", fmt.Errorf("unsupported format: %s", ext1)
	}

	var result string

	diffTree := genDiff(fileMap1, fileMap2)

	formatter := formatters.GetFormatter(format)

	if format == "plain" {
		result = formatter(diffTree)

	}

	if format == "stylish" || format != "plain" && format != "json" && format != "stylish" {
		result = "{\n" + formatter(diffTree) + "}"

	}

	if format == "json" {

		result = formatter(diffTree)

	}

	return result, nil

}

func genDiff(map1, map2 map[string]any) []models.Node {
	var result []models.Node

	allKeys := make(map[string]bool)

	for k := range map1 {
		allKeys[k] = true
	}
	for k := range map2 {
		allKeys[k] = true
	}

	keys := make([]string, 0, len(allKeys))
	for k := range allKeys {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		val1 := map1[key]
		val2 := map2[key]

		_, ok1 := val1.(map[string]any)
		_, ok2 := val2.(map[string]any)

		if ok1 && ok2 {
			node := models.Node{
				Type:     models.NESTED,
				Key:      key,
				Children: genDiff(val1.(map[string]any), val2.(map[string]any)),
			}
			result = append(result, node)
			continue
		}

		if _, exists := map2[key]; !exists {
			node := models.Node{
				Type:     models.REMOVED,
				Key:      key,
				OldValue: val1,
			}
			result = append(result, node)
			continue
		}

		if _, exists := map1[key]; !exists {
			node := models.Node{
				Type:  models.ADDED,
				Key:   key,
				Value: val2,
			}
			result = append(result, node)
			continue
		}

		if reflect.DeepEqual(val1, val2) {
			node := models.Node{
				Type:  models.UNCHANGED,
				Key:   key,
				Value: val1,
			}
			result = append(result, node)
			continue
		}

		node := models.Node{
			Type:     models.CHANGED,
			Key:      key,
			OldValue: val1,
			Value:    val2,
		}
		result = append(result, node)
	}

	return result
}
