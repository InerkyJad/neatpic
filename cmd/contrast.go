package cmd

import (
	"github.com/spf13/cobra"
)

// contrastCmd represents the contrast command
var contrastCmd = &cobra.Command{
	Use:   "contrast",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(contrastCmd)
}
