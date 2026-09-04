package main

import (
	"context"
	"fmt"
	"path/filepath"
	"github.com/urfave/cli/v3"
	"os"
	"sort"
	"github.com/bobkoffandrei/go-project-244/cmd/parsing"
	"github.com/bobkoffandrei/go-project-244/cmd/parsers"
	"strings"
	
)


func main() {

	cmd := &cli.Command{

		Name: "gendiff",

		Usage: "Compares two configuration files and shows a difference",

		Flags: []cli.Flag{
						&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Value:   "stylish",
				Usage:   "output format",
			},

		},

		Action: func(ctx context.Context, c *cli.Command) error {

			if c.Args().Get(0) == "" || c.Args().Get(1) == "" {
			err := cli.ShowAppHelp(c)
			if err != nil {
				return err
			}
			return fmt.Errorf("отсутствуют агрументы")

			}


            file1 := c.Args().Get(0)
			file2 := c.Args().Get(1)

			ext1 := filepath.Ext(c.Args().Get(0))
			ext2 := filepath.Ext(c.Args().Get(1))

            var fileMap1, fileMap2 map[string]any

            var err error

			if ext1 == ".json" && ext2 == ".json" {
				fileMap1, err = parsing.ParseFile(file1)
				if err != nil {
					return err
				}
				fileMap2, err = parsing.ParseFile(file2)
				if err != nil {
					return err
				}
			} else if (ext1 == ".yaml" && ext2 == ".yaml") || (ext1 == ".yml" && ext2 == ".yml") {
				fileMap1, err = parsers.ParseFile(file1)
				if err != nil {
					return err
				}
				fileMap2, err = parsers.ParseFile(file2)
				if err != nil {
					return err
				}
			} else if ext1 != ext2 {
				return fmt.Errorf("разные расширения файлов: %s и %s", ext1, ext2)
			} else {
				return fmt.Errorf("неподдерживаемый формат: %s", ext1)
			}

                diffTree := genDiff(fileMap1, fileMap2)

				formatter := getFormatter(c.String("format"))
                result := "{\n" + formatter(diffTree) + "}"

             			fmt.Println(result)




    
			return nil
		},
	}


		if err := cmd.Run(context.Background(), os.Args); err != nil {
		    fmt.Fprintf(os.Stderr, "ошибка выполнения программы: %v\n", err)
    		os.Exit(1)
	}


}


func formatStylish(nodes []Node) string {
	return formatStylishWithDepth(nodes, 0)
}

func getFormatter(format string) func([]Node) string {
    switch format {
    case "stylish":
        return formatStylish

    default:
        return formatStylish
    }
}
/*
func genDiff(map1, map2 map[string]any) string {
	return "{\n" + genDiff(map1, map2) + "}"
}
*/
const (
    UNCHANGED = "unchanged"
    ADDED     = "added"
    REMOVED   = "removed"
    CHANGED   = "changed"
    NESTED    = "nested"
)

type Node struct {
    Type     string          
    Key      string           
    Value    interface{}     
    OldValue interface{}      
    Children []Node           
}



func genDiff(map1, map2 map[string]any) []Node {
    var result []Node

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
            node := Node{
                Type:     NESTED,
                Key:      key,
                Children: genDiff(val1.(map[string]any), val2.(map[string]any)),
            }
            result = append(result, node)
            continue
        }
        
        if _, exists := map2[key]; !exists {
            node := Node{
                Type:     REMOVED,
                Key:      key,
                OldValue: val1,
            }
            result = append(result, node)
            continue
        }
        
        if _, exists := map1[key]; !exists {
            node := Node{
                Type:  ADDED,
                Key:   key,
                Value: val2,
            }
            result = append(result, node)
            continue
        }
        
        if val1 == val2 {
            node := Node{
                Type:  UNCHANGED,
                Key:   key,
                Value: val1,
            }
            result = append(result, node)
            continue
        }
        
        node := Node{
            Type:     CHANGED,
            Key:      key,
            OldValue: val1,
            Value:    val2,
        }
        result = append(result, node)
    }
    
    return result
}

func isMap(v any) bool {
    _, ok := v.(map[string]any)
    return ok
}

func formatMap(m map[string]any, indent string) string {
    var result string
    
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    
    for _, key := range keys {
        value := m[key]
        if isMap(value) {
            result += fmt.Sprintf("%s%s: {\n", indent, key)
            result += formatMap(value.(map[string]any), indent+"  ")
            result += fmt.Sprintf("%s}\n", indent)
        } else {
            result += fmt.Sprintf("%s%s: %v\n", indent, key, value)
        }
    }
    
    return result
}


       
func formatStylishWithDepth(nodes []Node, depth int) string {
    var result string
    indent := strings.Repeat("    ", depth)
    
    for _, node := range nodes {
        switch node.Type {
        case NESTED:

            result += fmt.Sprintf("%s  %s: {\n", indent, node.Key)
            result += formatStylishWithDepth(node.Children, depth+1)
            result += fmt.Sprintf("%s  }\n", indent)
            
        case UNCHANGED:
            result += fmt.Sprintf("%s  %s: %v\n", indent, node.Key, node.Value)
            
        case ADDED:

            if isMap(node.Value) {
                result += fmt.Sprintf("%s+ %s: {\n", indent, node.Key)
                result += formatMap(node.Value.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s+ %s: %v\n", indent, node.Key, node.Value)
            }
            
        case REMOVED:
            if isMap(node.OldValue) {
                result += fmt.Sprintf("%s- %s: {\n", indent, node.Key)
                result += formatMap(node.OldValue.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s- %s: %v\n", indent, node.Key, node.OldValue)
            }
            
        case CHANGED:

            if isMap(node.OldValue) {
                result += fmt.Sprintf("%s- %s: {\n", indent, node.Key)
                result += formatMap(node.OldValue.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s- %s: %v\n", indent, node.Key, node.OldValue)
            }
            
            if isMap(node.Value) {
                result += fmt.Sprintf("%s+ %s: {\n", indent, node.Key)
                result += formatMap(node.Value.(map[string]any), indent+"    ")
                result += fmt.Sprintf("%s  }\n", indent)
            } else {
                result += fmt.Sprintf("%s+ %s: %v\n", indent, node.Key, node.Value)
            }
        }
    }
    
    return result
}