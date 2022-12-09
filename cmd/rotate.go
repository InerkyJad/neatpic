// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/spf13/cobra"
	"image/color"
	"os"
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

func imagePathExists(args []string) bool {
	if len(args) > 1 {
		println("You can only rotate one image at a time or pass * to rotate all images in the directory")
		os.Exit(1)
	} else if len(args) == 0 {
		println("You must provide an image")
		os.Exit(1)
	}
	return true
}

func rotateImage(image string, angle int) {
	img, err := imaging.Open(image)
	if err != nil {
		fmt.Println("Error while opening the image")
		os.Exit(1)
	}

	img = imaging.Rotate(img, float64(angle), color.Black)

	err = imaging.Save(img, image)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

var rotateCmd = &cobra.Command{
	Use:        "rotate image [flags]",
	Short:      "Rotate an image, pass a deg e.g 90, 180 or don't and it will get rotated 90 degrees",
	Long:       ``,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"image"},
	Run: func(cmd *cobra.Command, args []string) {
		imagePathExists(args)
		var path string = args[0]

		/*
		*	Rotate All Images in a directory or subdirectories
		* */
		if path == "*" {
			path, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
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
			if _, err := os.Stat(path); os.IsNotExist(err) {
				fmt.Println("Image doesn't exist")
				os.Exit(1)
			}

			if path[len(path)-3:] != "jpg" && path[len(path)-3:] != "png" && path[len(path)-4:] != "jpeg" {
				fmt.Println("Image is not a jpg, png or jpeg")
				os.Exit(1)
			}

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
