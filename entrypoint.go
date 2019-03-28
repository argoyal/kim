package main

import (
	// "fmt"

	"log"
	"os"

	"kim/cmd"
	"kim/cmd/flags"

	"github.com/urfave/cli"
)

func main() {
	app := *cli.NewApp()

	app.Name = os.Getenv("APPLICATION_NAME")
	app.Usage = "A tool for infrastructure development and deployments"
	app.Version = "1.0.0"
	app.EnableBashCompletion = true

	cmd.SetupCommands(&app)
	baseFlags := flags.AddBaseFlags()

	app.Flags = baseFlags
	app.Before = flags.ValidateFlagOptions

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
