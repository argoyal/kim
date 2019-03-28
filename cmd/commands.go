/*
Package cmd for the cli app instance

SetupCommands sets up the commands needed for the app
*/
package cmd

import (
	"kim/cmd/start"

	"github.com/urfave/cli"
)

// import (
// 	"github.com/urfave/cli"
// )
//
// type command struct {
// 	Name, Usage string
// 	Action      func(c *cli.Context) error
// }
//
// type registrar interface {
// 	registerCommand()
// 	getCommandUsage() string
// }
//
// func registerCommand(allCommands *[]cli.Command, comm command) {
// 	newCommand := cli.Command{
// 		Name:   comm.Name,
// 		Usage:  comm.Usage,
// 		Action: comm.Action,
// 	}
//
// 	*allCommands = append(*allCommands, newCommand)
// }
//
// func getCommandUsage(comm string) string {
// 	return Commandusage(comm)
// }

// SetupCommands registers all the commands needed to the cli app
// instance
func SetupCommands(app *cli.App) {
	allCommands := []cli.Command{}

	command := start.Generatecommand()

	app.Commands = append(allCommands, command)
}
