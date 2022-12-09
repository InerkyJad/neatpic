package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

func brighten(path string, value float64) {
	img := getImage(path)
	img = imaging.AdjustBrightness(img, value)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// brightnessCmd represents the brightness command
var brightnessCmd = &cobra.Command{
	Use:   "brightness",
	Short: "A brief description of your command",
	Long:  `from 100 to -100 and 0 returns the image as it is`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		value, err := cmd.Flags().GetFloat64("value")
		if err != nil {
			fmt.Println("Error: cannot get value")
		}

		if value > 100 {
			value = 100
		} else if value < -100 {
			value = -100
		}

		brighten(path, value)
	},
}

func init() {
	rootCmd.AddCommand(brightnessCmd)
	brightnessCmd.Flags().Float64P("value", "v", 0, "from 100 to -100 and 0 returns the image as it is")
}
