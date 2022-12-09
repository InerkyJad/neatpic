package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
)

func compress(path string, percentage string) error {
	var img image.Image = getImage(path)

	p, err := strconv.Atoi(percentage)
	if err != nil {
		return err
	}

	ext := filepath.Ext(path)
	if ext == ".jpg" || ext == ".jpeg" {
		return jpeg.Encode(os.Stdout, img, &jpeg.Options{Quality: p})
	} else if ext == ".png" {
		return imaging.Save(img, path, imaging.JPEGQuality(p))
	}

	return nil
}

// reduceCmd represents the reduce command
var reduceCmd = &cobra.Command{
	Use:   "reduce",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		var path string = imagePathExists(args)
		var percentage string = cmd.Flag("percentage").Value.String()
		recursive, err := cmd.Flags().GetBool("recursive")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if percentage == "" {
			percentage = "75"
		}

		if path == "*" {
			var images []string = getImages(path, recursive)
			for _, image := range images {
				err := compress(image, percentage)
				if err != nil {
					fmt.Println("Error while compressing the image")
					os.Exit(1)
				}
			}
		} else {
			err := compress(path, percentage)
			if err != nil {
				fmt.Println("Error while compressing the image")
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(reduceCmd)
	reduceCmd.Flags().StringP("percentage", "p", "25", "Percentage of the image to reduce")
}
