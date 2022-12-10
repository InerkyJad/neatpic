package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// corp the image in path
func corpImageWidthHeight(path string, height string, width string) {
	img := getImage(path)
	heightInt, err := strconv.Atoi(height)
	if err != nil {
		heightInt = img.Bounds().Dy()
	}

	widthInt, err := strconv.Atoi(width)
	if err != nil {
		widthInt = img.Bounds().Dx()
	}

	img = imaging.CropCenter(img, widthInt, heightInt)
	err = imaging.Save(img, path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// corpCmd represents the corp command
var corpCmd = &cobra.Command{
	Use:   "corp image [flags]",
	Short: "Corp the image",
	Long:  `Corp the image with the given value of width (-W) and height (-H)`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)
		var width string = cmd.Flag("width").Value.String()
		var height string = cmd.Flag("height").Value.String()

		if width != "" || height != "" {
			if len(width) < 1 && len(height) < 1 {
				fmt.Println("Error: width or height cannot be empty")
				os.Exit(1)
			}
			corpImageWidthHeight(path, height, width)
		} else {
			fmt.Println("Error: Please provide any or --width and --height")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(corpCmd)
	corpCmd.Flags().StringP("width", "W", "", "width value")
	corpCmd.Flags().StringP("height", "H", "", "height value")
}
