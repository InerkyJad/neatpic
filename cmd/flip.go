// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// flipCmd represents the flip command
var flipCmd = &cobra.Command{
	Use:        "flip image [flags]",
	Short:      "Flip an image",
	Long:       `Flip an image`,
	Args:       cobra.MinimumNArgs(1),
	ArgAliases: []string{"image"},
	Run: func(cmd *cobra.Command, args []string) {
		//var path string = cmd.Flag("image").Value.String()
		fmt.Println(args)

		// print the last flag value

		/*var path string = cmd.Flag("image").Value.String()

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Println("The image path does not exist")
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

		fmt.Println(horizontal, vertical)*/
	},
}

func init() {
	rootCmd.AddCommand(flipCmd)
	flipCmd.Flags().BoolP("horizontal", "H", false, "Flip horizontally")
	flipCmd.Flags().BoolP("vertical", "V", false, "Flip vertically")
}
