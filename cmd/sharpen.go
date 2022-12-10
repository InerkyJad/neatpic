package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// change the sharpness of image in path
func sharpen(path string, value float64) {
	img := getImage(path)
	img = imaging.Sharpen(img, value)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// sharpenCmd represents the sharpen command
var sharpenCmd = &cobra.Command{
	Use:   "sharpen image [flags]",
	Short: "Sharpen the image",
	Long:  `Sharpen the image with the given value (-v) from 0 to 10 and 0 returns the image as it is`,
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

		sharpen(path, value)
	},
}

func init() {
	rootCmd.AddCommand(sharpenCmd)
	sharpenCmd.Flags().Float64P("value", "v", 1, "sharpen value")
}
