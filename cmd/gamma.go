package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// Change the gamma
func gamma(path string, value float64) {
	img := getImage(path)
	img = imaging.AdjustGamma(img, value)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// gammaCmd represents the gamma command
var gammaCmd = &cobra.Command{
	Use:   "gamma image [flags]",
	Short: "Change the gamma of the image",
	Long:  `Change the gamma of the image with the given value (-v) from 100 to -100 and 0 returns the image as it is`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		value, err := cmd.Flags().GetFloat64("value")
		if err != nil {
			fmt.Println("Error: cannot get value")
		}

		if value < 0 {
			fmt.Println("Error: value must be greater than 0")
			os.Exit(1)
		}

		gamma(path, value)
	},
}

func init() {
	rootCmd.AddCommand(gammaCmd)
	gammaCmd.Flags().Float64P("value", "v", 1, "gamma value")
}
