package util

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

// PrintMap pretty print based on the alignment of the keys and their length
func PrintMap(v map[string]string) (err error) {
	if err == nil {
		for key, value := range v {
			fmt.Printf("\n%v%100v\n\n", key, value)
		}
	}
	return
}

// InfoPrint prints the info text in cyan
func InfoPrint(text string, col color.Attribute) {
	d := color.New(col, color.Bold)

	log.Printf(d.Sprintf(text))
}

// FatalPrint prints the info text in cyan
func FatalPrint(text string) {
	d := color.New(color.FgRed, color.Bold)

	log.Fatal(d.Sprintf(text))
}

// PrintTabbed prints a text along with the value in a nice neat structure
func PrintTabbed(text, value string, col color.Attribute) {
	output := fmt.Sprintf("%-50s...%s", text, value)

	InfoPrint(output, col)
}
