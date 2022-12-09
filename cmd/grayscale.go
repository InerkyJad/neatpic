package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

func grayscale(path string) {
	img := getImage(path)
	img = imaging.Grayscale(img)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// greyscaleCmd represents the greyscale command
var grayscaleCmd = &cobra.Command{
	Use:   "grayscale",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		grayscale(path)
	},
}

func init() {
	rootCmd.AddCommand(grayscaleCmd)
}
