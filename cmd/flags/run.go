package flags

import (
	"github.com/urfave/cli"
)

// ValidateFlagOptions runs the flag validators / options tasks that they
// are supposed to do.
func ValidateFlagOptions(c *cli.Context) error {
	if c.String("config") != "" {
		validateConfig(c)
	}

	return nil
}
