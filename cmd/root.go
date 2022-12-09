// Package cmd /*
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "neatpic",
	Short: "NeatPic is a CLI tool to manipulate images",
	Long: `
	/$$   /$$                       /$$     /$$$$$$$  /$$          
	| $$$ | $$                      | $$    | $$__  $$|__/          
	| $$$$| $$  /$$$$$$   /$$$$$$  /$$$$$$  | $$  \ $$ /$$  /$$$$$$$
	| $$ $$ $$ /$$__  $$ |____  $$|_  $$_/  | $$$$$$$/| $$ /$$_____/
	| $$  $$$$| $$$$$$$$  /$$$$$$$  | $$    | $$____/ | $$| $$      
	| $$\  $$$| $$_____/ /$$__  $$  | $$ /$$| $$      | $$| $$      
	| $$ \  $$|  $$$$$$$|  $$$$$$$  |  $$$$/| $$      | $$|  $$$$$$$
	|__/  \__/ \_______/ \_______/   \___/  |__/      |__/ \_______/

	NeatPic is a CLI tool to manipulate images
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
