package cmd

import (
	"github.com/spf13/cobra"
)

// blurCmd represents the blur command
var blurCmd = &cobra.Command{
	Use:   "blur",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(blurCmd)
}
