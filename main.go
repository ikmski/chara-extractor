package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func mainAction(c *cli.Context) error {

	if c.NArg() == 0 {
		return fmt.Errorf("input is required")
	}
	txt := c.Args()

	unique := make(map[rune]rune)

	for _, line := range txt {
		chars := []rune(line)

		for _, char := range chars {

			_, ok := unique[char]
			if !ok {
				unique[char] = char
			}
		}
	}

	runes := []rune{}
	for key, _ := range unique {
		runes = append(runes, key)
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for _, r := range runes {
		fmt.Printf("%v", string(r))
	}

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "chara-extractor"
	app.Usage = "Character Extractor"

	app.Action = mainAction

	app.Run(os.Args)
}
