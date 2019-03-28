package flags

import (
	"os"

	"github.com/urfave/cli"
)

func validateConfig(c *cli.Context) error {
	configPath := c.String("config")

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return cli.NewExitError("config file not found", 10)
	}

	return nil
}
