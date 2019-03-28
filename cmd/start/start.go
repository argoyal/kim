package start

import (
	"fmt"
	"kim/config"
	"kim/util"

	"github.com/urfave/cli"
)

// Generatecommand generates the cli command for init
func Generatecommand() cli.Command {
	return cli.Command{
		Name:                   "init",
		Usage:                  "Perform initialization of a raw infrastructure package",
		Action:                 runInit,
		UseShortOptionHandling: true,
		BashComplete: func(c *cli.Context) {
			fmt.Fprintf(c.App.Writer, "init")
		},
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "include-files, e",
				Usage: "Create empty files in the app created on initialization",
			},
			cli.BoolFlag{
				Name:  "enable-docker, d",
				Usage: "Enables docker repository as per configurations",
			},
			cli.BoolFlag{
				Name:  "enable-git, g",
				Usage: "Enables git repository for the microservice",
			},
		},
	}
}

func runInit(c *cli.Context) error {
	msName := c.Args().Get(0)

	if msName == "" {
		return cli.NewExitError("Microservice name is a must", 10)
	}

	config := config.GetConfig(c)

	util.Block{
		Try: func() {
			runAction(c, config, msName)
		},
		Catch: func(e util.Exception) {
			util.FatalPrint(fmt.Sprintf("%v", e))
		},
		Finally: func() {},
	}.Do()

	return nil
}
