package main

import (
	"fmt"

	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func outputChar(runes []rune, oneline bool) {

	for i := 0; i < len(runes)-1; i++ {
		if oneline {
			fmt.Printf("%v", string(runes[i]))
		} else {
			fmt.Printf("%v\n", string(runes[i]))
		}
	}
	fmt.Printf("%v\n", string(runes[len(runes)-1]))
}

func outputDecimal(runes []rune, oneline bool) {

	for i := 0; i < len(runes)-1; i++ {
		if oneline {
			fmt.Printf("%d,", runes[i])
		} else {
			fmt.Printf("%d\n", runes[i])
		}
	}
	fmt.Printf("%d\n", runes[len(runes)-1])
}

func outputHexa(runes []rune, oneline bool) {

	for i := 0; i < len(runes)-1; i++ {
		if oneline {
			fmt.Printf("%X,", runes[i])
		} else {
			fmt.Printf("%X\n", runes[i])
		}
	}
	fmt.Printf("%X\n", runes[len(runes)-1])
}

func mainAction(c *cli.Context) error {

	if c.NArg() == 0 {
		return fmt.Errorf("input is required")
	}
	text := c.Args()

	if c.Bool("d") {
		outputDecimal(extract(text), c.Bool("l"))
	} else if c.Bool("x") {
		outputHexa(extract(text), c.Bool("l"))
	} else {
		outputChar(extract(text), c.Bool("l"))
	}

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "chara-extractor"
	app.Usage = "Character Extractor"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "decimal, d",
			Usage: "output decimal number?",
		},
		cli.BoolFlag{
			Name:  "hexa, x",
			Usage: "output hexadecimal number?",
		},
		cli.BoolFlag{
			Name:  "oneline, l",
			Usage: "output oneline?",
		},
	}

	app.Action = mainAction

	app.RunAndExitOnError()
}
