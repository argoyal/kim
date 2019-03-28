package start

import (
	"kim/config"
	"kim/util"
	"os"

	"github.com/fatih/color"
)

// createEmptyFiles generates empty files in the msName
func createEmptyFiles(config config.Configstructure, msPath string) error {
	emptyTemplates := config.EmptyTemplates

	if config.IncludeTemplateVariableFile {
		emptyTemplates = append(emptyTemplates, config.TemplateVariableFileName)
	}

	util.PrintTabbed("Generating templates in folder", msPath, color.FgCyan)

	for _, template := range emptyTemplates {
		templatePath := msPath + "/" + template
		_, err := os.Create(templatePath)

		if err != nil {
			util.Throw(err)
		}
	}

	return nil
}
