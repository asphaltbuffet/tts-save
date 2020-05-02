package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	log "github.com/sirupsen/logrus"
)

func main() {
	app := &cli.App{
		Name:  "TTSSave",
		Usage: "Import/export tabletop simulator save files",
		Action: func(c *cli.Context) error {
			fmt.Println("Wrong. Try again.")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
