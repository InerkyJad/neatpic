package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// change the image contrast
func contrast(path string, value float64) {
	img := getImage(path)
	img = imaging.AdjustContrast(img, value)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// contrastCmd represents the contrast command
var contrastCmd = &cobra.Command{
	Use:   "contrast",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		value, err := cmd.Flags().GetFloat64("value")
		if err != nil {
			fmt.Println("Error: cannot read value flag")
			os.Exit(1)
		}

		if value > 100 {
			value = 100
		} else if value < -100 {
			value = -100
		}

		contrast(path, value)
	},
}

func init() {
	rootCmd.AddCommand(contrastCmd)
	contrastCmd.Flags().Float64P("value", "v", 0, "Contrast percentage (from -100 to 100)")
}
