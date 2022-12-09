package cmd

import (
	"github.com/spf13/cobra"
)

// sharpenCmd represents the sharpen command
var sharpenCmd = &cobra.Command{
	Use:   "sharpen",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(sharpenCmd)
}