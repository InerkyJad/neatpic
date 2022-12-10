package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

// resize image wile keeping aspect ratio (if only one parameter is passed)
func resize(path string, width float64, height float64) {
	img := getImage(path)
	img = imaging.Resize(img, int(width), int(height), imaging.Lanczos)
	err := imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
	}
}

// resizeCmd represents the resize command
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		height, err := cmd.Flags().GetFloat64("height")
		if err != nil {
			fmt.Println("Error: unable to read height flag")
			os.Exit(1)
		}
		width, err := cmd.Flags().GetFloat64("width")
		if err != nil {
			fmt.Println("Error: unable to read width flag")
		}

		if height == 0 && width == 0 {
			fmt.Println("Error: you need to pass at least of of --height or --width")
			os.Exit(1)
		}

		resize(path, width, height)
	},
}

func init() {
	rootCmd.AddCommand(resizeCmd)
	resizeCmd.Flags().Float64P("height", "H", 0, "desired height")
	resizeCmd.Flags().Float64P("width", "W", 0, "desired width")
}
