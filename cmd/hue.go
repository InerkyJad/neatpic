package cmd

import (
	"github.com/spf13/cobra"
)

// hueCmd represents the hue command
var hueCmd = &cobra.Command{
	Use:   "hue",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
	},
}

func init() {
	rootCmd.AddCommand(hueCmd)

}
