package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// change saturation
func saturation(path string, percentage float64) {
	img := getImage(path)
	img = imaging.AdjustSaturation(img, percentage)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// saturationCmd represents the saturation command
var saturationCmd = &cobra.Command{
	Use:   "saturation",
	Short: "A brief description of your command",
	Long:  ``,
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

		saturation(path, value)
	},
}

func init() {
	rootCmd.AddCommand(saturationCmd)
	saturationCmd.Flags().Float64P("value", "v", 0, "Saturation value from -100 to 100")
}
