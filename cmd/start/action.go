package start

import (
	"fmt"
	"kim/config"
	"kim/database/schema"
	"kim/util"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func runAction(c *cli.Context, config config.Configstructure, msName string) error {
	dir, err := os.Getwd()
	dockerRepo := ""

	if err != nil {
		util.Throw(err)
	}

	msPath := dir + "/" + msName
	isPresent := util.Exists(msPath)

	if isPresent {
		util.FatalPrint("A folder with following name already exists")
		return nil
	}

	util.PrintTabbed("Generating microservice infra folder", msPath, color.FgCyan)
	os.Mkdir(msName, os.ModePerm)

	if c.Bool("include-files") {
		createEmptyFiles(config, msPath)
	}

	if c.Bool("enable-docker") {
		dockerRepo = createDockerRepository(config, msName)
	}

	msDetail := schema.MicroserviceDetail{
		MicroserviceName:             msName,
		MicroserviceDockerRepository: dockerRepo,
	}

	fmt.Println(&msDetail)
	//
	// if c.Bool("enable-git") {
	// 	enable
	// }

	// database.SaveMicroserviceDetail(config, &msDetail)

	return nil
}

// tears down all the tasks if there is a panic error
func teardownAction() {}
