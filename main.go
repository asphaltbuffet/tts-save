package main

import (
	"os"
	"github.com/urfave/cli/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := &cli.App{
		Name: "Tabletop Simulator Saves Tool"
		Usage: "Import/export tabletop simulator save files"
		Action: func(c *cli.Context) error{
			fmt.Println("Wrong. Try again.")
			return nil
		},
	}

	err:= app.Run(os.Args)
	if err != nil{
		logrus.Fatal(err)
	}
}