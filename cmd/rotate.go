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

// Remove from Slice by value
func remove(slice []string, s string) []string {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

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

// Get the list of images in a directory
func getImages(path string, recursive bool) []string {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error while opening the directory")
		os.Exit(1)
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println("Error while reading the directory")
		os.Exit(1)
	}
	var images []string
	for _, file := range files {
		if file.IsDir() {
			if recursive {
				images = append(images, getImages(path+file.Name(), recursive)...)
			}
		} else {
			images = append(images, path+file.Name())
		}
	}

	for _, image := range images {
		if image[len(image)-3:] != "jpg" && image[len(image)-3:] != "png" && image[len(image)-4:] != "jpeg" {
			images = remove(images, image)
		}
	}

	return images
}

func rotateImage(image string, angle int) {
	img := getImage(image)
	img = imaging.Rotate(img, float64(angle), color.Black)
	err := imaging.Save(img, image)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// get the working directory
func getDir() string {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return path
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

	err, _ := os.Stat(args[0])
	if err != nil {
		fmt.Println("Error: The image path is not valid")
		os.Exit(1)
	}

	return args[0]
}

var rotateCmd = &cobra.Command{
	Use:        "rotate image [flags]",
	Short:      "Rotate an image, pass a deg e.g 90, 180 or don't and it will get rotated 90 degrees",
	Long:       ``,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"image"},
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)

		/*
		*	Rotate All Images in a directory or subdirectories
		* */
		if path == "*" {
			path = getDir()
			var images []string = getImages(path, false)

			/*
			* Rotate all the images
			 */
			for _, image := range images {
				if cmd.Flag("angle").Value.String() == "90" {
					rotateImage(image, 90)
				} else if cmd.Flag("angle").Value.String() == "180" {
					rotateImage(image, 180)
				} else {
					rotateImage(image, 90)
				}
			}

		} else
		/*
		*	Rotate a single image
		* */
		{
			if cmd.Flag("angle").Value.String() == "90" {
				rotateImage(path, 90)
			} else if cmd.Flag("angle").Value.String() == "180" {
				rotateImage(path, 180)
			} else {
				rotateImage(path, 90)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)
	rotateCmd.Flags().IntP("angle", "a", 90, "Angle of the rotation, 90, 180 (optional)")
}
