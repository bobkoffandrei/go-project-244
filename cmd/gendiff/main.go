package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
	"sort"
    "github.com/bobkoffandrei/go-project-244/code/models"
    "github.com/bobkoffandrei/go-project-244/code"
	//"strings"
	
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

            format := c.String("format")

            result, err := code.GenDiff(file1, file2, format)
            
			if err != nil {
				return err
			}

			fmt.Println(result)    
			return nil
		},
	}


		if err := cmd.Run(context.Background(), os.Args); err != nil {
		    fmt.Fprintf(os.Stderr, "ошибка выполнения программы: %v\n", err)
    		os.Exit(1)
	}


}







/*
func genDiff(map1, map2 map[string]any) string {
	return "{\n" + genDiff(map1, map2) + "}"
}
*/





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
        
        if val1 == val2 {
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






       
