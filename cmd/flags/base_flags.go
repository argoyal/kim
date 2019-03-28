package flags

import (
	"os"

	"github.com/urfave/cli"
)

// AddBaseFlags adds all the flags for the cli tool
func AddBaseFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Specify the config file needed for the project",
			Value: os.Getenv("HOME") + "/.kimrc.hjson",
		},
	}
}
