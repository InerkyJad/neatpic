package cmd

import (
	"github.com/spf13/cobra"
)

// gammaCmd represents the gamma command
var gammaCmd = &cobra.Command{
	Use:   "gamma",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(gammaCmd)
}
