package main

import (
	"context"
	"fmt"
	//"reflect"
	//"github.com/bobkoffandrei/go-project-244/code"
	"github.com/urfave/cli/v3"
	"os"
	"sort"
	"github.com/bobkoffandrei/go-project-244/cmd/parsing"
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

			fileMap1, err := parsing.ParseFile(c.Args().Get(0))
			if err != nil {
				return err
			}
			
			fileMap2, err := parsing.ParseFile(c.Args().Get(1))

			if err != nil {
				return err
			}

			fmt.Println(genDiff(fileMap1, fileMap2))

			return nil
		},
	}


		if err := cmd.Run(context.Background(), os.Args); err != nil {
		    fmt.Fprintf(os.Stderr, "ошибка выполнения программы: %v\n", err)
    		os.Exit(1)
	}


}

func genDiff(map1, map2 map[string]any) string {
	var result string

		keys1 := make([]string, 0, len(map1))
		keys2 := make([]string, 0, len(map2))

			for k := range map1 {
		keys1 = append(keys1, k)
	}

			for k := range map2 {
		keys2 = append(keys2, k)
	}

	sort.Strings(keys1)
	sort.Strings(keys2)
	
	for _, key := range keys1 {
		if map1[key] == map2[key] {
		result += fmt.Sprintf("  %s: %v\n", key, map1[key])
		}
		if map1[key] != map2[key] {
		result += fmt.Sprintf("- %s: %v\n", key, map1[key])
		if map2[key] != nil {
		result += fmt.Sprintf("+ %s: %v\n", key, map2[key])
		}
		
		}
	}

	for _, key := range keys2 {
		if map1[key] == nil  {
		
		result += fmt.Sprintf("+ %s: %v\n", key, map2[key])
		}
		}

	return result
}