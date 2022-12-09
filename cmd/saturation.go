package cmd

import (
	"github.com/spf13/cobra"
)

// saturationCmd represents the saturation command
var saturationCmd = &cobra.Command{
	Use:   "saturation",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(saturationCmd)
}
