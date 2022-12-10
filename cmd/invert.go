package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// invert the image in path
func invert(path string) {
	img := getImage(path)
	img = imaging.Invert(img)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// invertCmd represents the invert command
var invertCmd = &cobra.Command{
	Use:   "invert",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		invert(path)
	},
}

func init() {
	rootCmd.AddCommand(invertCmd)
}
