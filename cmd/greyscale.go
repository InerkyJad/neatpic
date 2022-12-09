package cmd

import (
	"github.com/spf13/cobra"
)

// greyscaleCmd represents the greyscale command
var greyscaleCmd = &cobra.Command{
	Use:   "greyscale",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(greyscaleCmd)
}
