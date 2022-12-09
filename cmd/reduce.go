package cmd

import (
	"github.com/spf13/cobra"
)

// reduceCmd represents the reduce command
var reduceCmd = &cobra.Command{
	Use:   "reduce",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		imagePathExists(args)
		var path string = args[0]
		var percentage string = cmd.Flag("percentage").Value.String()
		var recursive bool = cmd.Flag("recursive").Value.String() == "true"

		if path == "*" {
			var images []string = getImages(path, recursive)
			for _, image := range images {
				reduce(image, percentage)
			}
		} else {
			reduce(path, percentage)
		}

	},
}

func init() {
	rootCmd.AddCommand(reduceCmd)
	reduceCmd.Flags().StringP("percentage", "p", "25", "Percentage of the image to reduce")
}
