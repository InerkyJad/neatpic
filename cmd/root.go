// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "neatpic [command] [flags] [image]",
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
