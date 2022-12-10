// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"image"
	"image/color"
	"os"
	"path/filepath"
)

// get an image, it will check if the image exists ...
func getImage(path string) image.Image {
	// Check if the path is a directory
	fi, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	if fi.IsDir() {
		fmt.Println("Error:", path, "is a directory")
		os.Exit(1)
	}

	// Check if the path is an image file
	ext := filepath.Ext(path)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		fmt.Println("Error:", path, "is not an image file")
		os.Exit(1)
	}

	// Open the image file
	img, err := imaging.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return img
}

// rotate the image and save it
func rotateImage(image string, angle int) {
	img := getImage(image)
	img = imaging.Rotate(img, float64(angle), color.Black)
	err := imaging.Save(img, image)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// check if the path was passed and is valid
func imagePathExists(args []string) string {
	if len(args) == 0 {
		fmt.Println("Error: No image path was passed")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("Error: The image path is not valid")
		os.Exit(1)
	}

	path := args[0]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Error: The image path is not valid")
		os.Exit(1)
	}

	return args[0]
}

var rotateCmd = &cobra.Command{
	Use:        "rotate image [flags]",
	Short:      "Rotate the image",
	Long:       `Rotate the image by 90 or 180 degrees (default 90)`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"image"},
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		saveImage(path)

		if cmd.Flag("angle").Value.String() == "90" {
			rotateImage(path, 90)
		} else if cmd.Flag("angle").Value.String() == "180" {
			rotateImage(path, 180)
		} else {
			rotateImage(path, 90)
		}

	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)
	rotateCmd.Flags().IntP("angle", "a", 90, "Angle of rotation")
}
