package cmd

import (
	"github.com/spf13/cobra"
)

// invertCmd represents the invert command
var invertCmd = &cobra.Command{
	Use:   "invert",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(invertCmd)
}
