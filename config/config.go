package config

import (
	"encoding/json"
	"io/ioutil"
	"kim/util"

	"github.com/fatih/color"
	hjson "github.com/hjson/hjson-go"
	"github.com/urfave/cli"
)

// GetConfig gets the config from the file
func GetConfig(c *cli.Context) Configstructure {
	configPath := c.GlobalString("config")

	util.PrintTabbed("Using configuration file from", configPath, color.FgBlue)

	var dat map[string]interface{}

	f, err := ioutil.ReadFile(configPath)

	if err != nil {
		util.Throw(err)
	}

	if err := hjson.Unmarshal(f, &dat); err != nil {
		util.Throw(err)
	}

	b, _ := json.Marshal(dat)

	var AppConfig Configstructure

	json.Unmarshal(b, &AppConfig)

	return AppConfig
}
