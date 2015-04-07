package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "AWS Polling Sample"
	app.Usage = "./go-poll"
	app.Commands = getCommands()
	app.Run(os.Args)
}
