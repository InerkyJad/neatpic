package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// blur the image in path
func blur(path string, degree float64) {
	img := getImage(path)
	img = imaging.Blur(img, degree)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// blurCmd represents the blur command
var blurCmd = &cobra.Command{
	Use:   "blur image [flags]",
	Short: "Blur the image",
	Long:  `Blur the image with the given value of blur (-v)`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		degree, err := cmd.Flags().GetFloat64("value")
		if err != nil {
			fmt.Println("Error getting degree flag")
			os.Exit(1)
		}

		if degree > 10 {
			degree = 10
		}

		blur(path, degree)
	},
}

func init() {
	rootCmd.AddCommand(blurCmd)
	blurCmd.Flags().Float64P("value", "v", 1, "blur value")
}
