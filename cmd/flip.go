package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

func flip(image string, horizontal bool, vertical bool) {
	if vertical {
		img, err := imaging.Open(image)
		if err != nil {
			fmt.Println("Error while opening the image")
			os.Exit(1)
		}
		img = imaging.FlipV(img)
		err = imaging.Save(img, image)
		if err != nil {
			fmt.Println("Error while saving the image")
			os.Exit(1)
		}
	}

	if horizontal {
		img, err := imaging.Open(image)
		if err != nil {
			fmt.Println("Error while opening the image")
			os.Exit(1)
		}
		img = imaging.FlipH(img)
		err = imaging.Save(img, image)
		if err != nil {
			fmt.Println("Error while saving the image")
			os.Exit(1)
		}
	}
}

// flipCmd represents the flip command
var flipCmd = &cobra.Command{
	Use:        "flip image [flags]",
	Short:      "Flip an image horizontally or vertically, if no flag is provided, the image is flipped horizontally",
	Long:       `Flip an image`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"image"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			println("You can only flip one image at a time or pass * to flip all images in the directory")
			os.Exit(1)
		} else if len(args) == 0 {
			println("You must provide an image")
			os.Exit(1)
		}

		var path string = args[0]

		if path == "*" {
			var images []string = getImages(path, false)

			for _, image := range images {
				if cmd.Flag("horizontal").Value.String() == "true" {
					flip(image, true, false)
				} else if cmd.Flag("vertical").Value.String() == "true" {
					flip(image, false, true)
				} else {
					flip(image, true, false)
				}
			}
		} else {
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("No Image found at path: " + path)
				os.Exit(1)
			}

			horizontal, err := cmd.Flags().GetBool("horizontal")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			vertical, err := cmd.Flags().GetBool("vertical")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if horizontal {
				flip(path, true, false)
			}

			if vertical {
				flip(path, false, true)
			}

			if !horizontal && !vertical {
				flip(path, true, false)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(flipCmd)
	flipCmd.Flags().BoolP("horizontal", "H", false, "Flip horizontally")
	flipCmd.Flags().BoolP("vertical", "V", false, "Flip vertically")
}
