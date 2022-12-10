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

// compress/reduce the image size (Not Working)
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
	Use:   "reduce image [flags]",
	Short: "Reduce the image size",
	Long:  `Reduce the image size`,
	Run: func(cmd *cobra.Command, args []string) {

		var path string = imagePathExists(args)
		saveImage(path)
		var percentage string = cmd.Flag("value").Value.String()

		if percentage == "" {
			percentage = "75"
		}

		err := compress(path, percentage)
		if err != nil {
			fmt.Println("Error while compressing the image")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(reduceCmd)
	reduceCmd.Flags().StringP("value", "v", "25", "Percentage of compression")
}
