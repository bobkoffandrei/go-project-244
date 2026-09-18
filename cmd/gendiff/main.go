package main

import (
	"code"
	"context"
	"fmt"
	"github.com/urfave/cli/v3"
	"os"
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
				return fmt.Errorf("missing arguments: Pass two files as arguments for comparison")

			}

			file1 := c.Args().Get(0)
			file2 := c.Args().Get(1)

			format := c.String("format")

			result, err := code.GenDiff(file1, file2, format)

			if err != nil {
				return fmt.Errorf("parsing error: %v", err)
			}

			fmt.Println(result)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "program execution error: %v\n", err)
		os.Exit(1)
	}

}
