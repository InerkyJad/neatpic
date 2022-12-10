package cmd

import (
	"fmt"
	"github.com/disintegration/imaging"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// save image for restoring
func saveImage(path string) {
	tempFolder := os.TempDir()
	tempPath := filepath.Join(tempFolder, "neatpic")
	if _, err := os.Stat(tempPath); os.IsNotExist(err) {
		err := os.Mkdir(tempPath, 0755)
		if err != nil {
			fmt.Println("Error while creating the temp folder")
			os.Exit(1)
		}
	}

	tempImage := filepath.Join(tempPath, "temp.png")
	err := imaging.Save(getImage(path), tempImage)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// restore image from temp folder
func restoreImage(path string) {
	tempFolder := os.TempDir()
	tempPath := filepath.Join(tempFolder, "neatpic")
	tempImage := filepath.Join(tempPath, "temp.png")
	err := imaging.Save(getImage(tempImage), path)
	if err != nil {
		fmt.Println("Error while saving the image")
		os.Exit(1)
	}
}

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore image",
	Short: "Restore the image",
	Long:  `Restore the image with the given path, it will restore the image to the state before the last command`,
	Run: func(cmd *cobra.Command, args []string) {
		var path string = imagePathExists(args)
		restoreImage(path)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
