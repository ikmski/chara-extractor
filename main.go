package main

import (
	"os"

	"github.com/urfave/cli"
)

func mainAction(c *cli.Context) error {

	return nil
}

func main() {

	app := cli.NewApp()
	app.Name = "chara-extractor"
	app.Usage = "Character Extractor"

	app.Action = mainAction

	app.Run(os.Args)
}
