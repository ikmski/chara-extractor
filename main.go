package main

import (
	"fmt"

	"github.com/urfave/cli"
)

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
