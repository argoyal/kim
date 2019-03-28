package start

import (
	"fmt"
	"kim/config"
	"kim/docker"
	"kim/util"

	"github.com/fatih/color"
)

// Creates docker repository
func createDockerRepository(config config.Configstructure, msName string) string {
	suggestedRepo := config.InfrastructureEnvironment + "/" + msName

	util.PrintTabbed("Suggested Repository Name", suggestedRepo, color.FgCyan)

	var userInput string
	color.Green("Do you want to use the Suggested Repository URI? [y/n]")

	_, err := fmt.Scanln(&userInput)

	if err != nil {
		util.Throw(err)
	}

	if userInput == "y" || userInput == "" {
		util.PrintTabbed("Creating docker repository on", config.ImageRepositoryProvider, color.FgCyan)
		return docker.CreateDockerRepository(config, suggestedRepo)
	}

	color.Green("Enter the Repository name for your image [Enter only the name]")

	_, err = fmt.Scanln(&userInput)

	if err != nil {
		util.Throw(err)
	}

	util.PrintTabbed("Creating docker repository on", config.ImageRepositoryProvider, color.FgCyan)
	return docker.CreateDockerRepository(config, userInput)
}
