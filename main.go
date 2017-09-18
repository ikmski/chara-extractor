package main

import (
	"fmt"
	"sort"

	"github.com/urfave/cli"
)

func extract(text []string) []rune {

	m := make(map[rune]rune)

	for _, line := range text {

		chars := []rune(line)
		for _, char := range chars {

			_, ok := m[char]
			if !ok {
				m[char] = char
			}
		}
	}

	runes := []rune{}
	for key, _ := range m {
		runes = append(runes, key)
	}

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return runes
}

func outputChar(runes []rune) {

	for _, r := range runes {
		fmt.Printf("%v", string(r))
	}
	fmt.Print("\n")
}

func outputDecimal(runes []rune) {

	for i := 0; i < len(runes)-1; i++ {
		fmt.Printf("%d,", runes[i])
	}
	fmt.Printf("%d\n", runes[len(runes)-1])
}

func mainAction(c *cli.Context) error {

	if c.NArg() == 0 {
		return fmt.Errorf("input is required")
	}
	text := c.Args()

	if c.Bool("d") {
		outputDecimal(extract(text))
	} else {
		outputChar(extract(text))
	}

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "chara-extractor"
	app.Usage = "Character Extractor"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "decimal, d",
			Usage: "output decimal number?",
		},
	}

	app.Action = mainAction

	app.RunAndExitOnError()
}
