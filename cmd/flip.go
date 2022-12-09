package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"os"
)

func flip(image string, horizontal bool, vertical bool) {
	if vertical {
		img := imaging.FlipV(getImage(image))
		err := imaging.Save(img, image)
		if err != nil {
			fmt.Println("Error while saving the image")
			os.Exit(1)
		}
	}

	if horizontal {
		img := imaging.FlipH(getImage(image))
		err := imaging.Save(img, image)
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
		var path string = imagePathExists(args)
		saveImage(path)

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
	},
}

func init() {
	rootCmd.AddCommand(flipCmd)
	flipCmd.Flags().BoolP("horizontal", "H", false, "Flip horizontally")
	flipCmd.Flags().BoolP("vertical", "V", false, "Flip vertically")
}
